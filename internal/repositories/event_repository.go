package repositories

import (
	"database/sql"
	"strings"
	"time"

	"middleware/internal/models"
)

type EventRepository struct {
	DB *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{DB: db}
}

// InitSchema crée la table events si elle n'existe pas.
func (r *EventRepository) InitSchema() error {
	_, err := r.DB.Exec(`
CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	agenda_id TEXT NOT NULL,
	uid TEXT NOT NULL,
	title TEXT NOT NULL,
	description TEXT,
	location TEXT,
	start TEXT NOT NULL,
	end TEXT NOT NULL,
	last_modified TEXT,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL,
	UNIQUE(agenda_id, uid)
);
`)
	return err
}

// CreateEvent insère un event en base.
func (r *EventRepository) CreateEvent(e *models.Event) (*models.Event, error) {
	if e == nil {
		return nil, &models.ErrorUnprocessableEntity{Message: "event is required"}
	}
	if strings.TrimSpace(e.AgendaID) == "" {
		return nil, &models.ErrorUnprocessableEntity{Message: "agenda_id is required"}
	}
	if strings.TrimSpace(e.UID) == "" {
		return nil, &models.ErrorUnprocessableEntity{Message: "uid is required"}
	}
	if strings.TrimSpace(e.Title) == "" {
		return nil, &models.ErrorUnprocessableEntity{Message: "title is required"}
	}
	if !e.Start.IsZero() && !e.End.IsZero() && !e.Start.Before(e.End) {
		return nil, &models.ErrorUnprocessableEntity{Message: "start must be before end"}
	}

	now := time.Now().UTC()
	if e.CreatedAt.IsZero() {
		e.CreatedAt = now
	}
	e.UpdatedAt = now

	res, err := r.DB.Exec(`
INSERT INTO events (agenda_id, uid, title, description, location, start, end, last_modified, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`,
		e.AgendaID,
		e.UID,
		e.Title,
		nullIfEmpty(e.Description),
		nullIfEmpty(e.Location),
		e.Start.UTC().Format(time.RFC3339),
		e.End.UTC().Format(time.RFC3339),
		nullTimeRFC3339(e.LastModified),
		e.CreatedAt.UTC().Format(time.RFC3339),
		e.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		// UNIQUE(agenda_id, uid)
		if isSQLiteUniqueViolation(err) {
			return nil, &models.ErrorConflict{Message: "event already exists for this agenda_id and uid"}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err == nil {
		e.ID = id
	}
	return e, nil
}

// GetAllEvents liste les events, avec filtre optionnel agenda_id.
func (r *EventRepository) GetAllEvents(agendaID string) ([]models.Event, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if strings.TrimSpace(agendaID) == "" {
		rows, err = r.DB.Query(`
SELECT id, agenda_id, uid, title, description, location, start, end, last_modified, created_at, updated_at
FROM events
ORDER BY start ASC
`)
	} else {
		rows, err = r.DB.Query(`
SELECT id, agenda_id, uid, title, description, location, start, end, last_modified, created_at, updated_at
FROM events
WHERE agenda_id = ?
ORDER BY start ASC
`, agendaID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Event
	for rows.Next() {
		ev, scanErr := scanEvent(rows)
		if scanErr != nil {
			return nil, scanErr
		}
		out = append(out, *ev)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// GetEventByID récupère un event par id.
func (r *EventRepository) GetEventByID(id int64) (*models.Event, error) {
	row := r.DB.QueryRow(`
SELECT id, agenda_id, uid, title, description, location, start, end, last_modified, created_at, updated_at
FROM events
WHERE id = ?
`, id)

	ev, err := scanEvent(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.ErrorNotFound{Message: "event not found"}
		}
		return nil, err
	}
	return ev, nil
}

// UpdateEvent met à jour title/description/location/start/end/last_modified.
func (r *EventRepository) UpdateEvent(id int64, patch map[string]interface{}) (*models.Event, error) {
	// Charge l'existant
	existing, err := r.GetEventByID(id)
	if err != nil {
		return nil, err
	}

	// Applique patch (logique simple)
	if v, ok := patch["title"].(string); ok {
		if strings.TrimSpace(v) == "" {
			return nil, &models.ErrorUnprocessableEntity{Message: "title cannot be empty"}
		}
		existing.Title = v
	}
	if v, ok := patch["description"].(string); ok {
		existing.Description = v
	}
	if v, ok := patch["location"].(string); ok {
		existing.Location = v
	}
	if v, ok := patch["start"].(time.Time); ok {
		existing.Start = v
	}
	if v, ok := patch["end"].(time.Time); ok {
		existing.End = v
	}
	if v, ok := patch["last_modified"].(time.Time); ok {
		existing.LastModified = v
	}

	if !existing.Start.IsZero() && !existing.End.IsZero() && !existing.Start.Before(existing.End) {
		return nil, &models.ErrorUnprocessableEntity{Message: "start must be before end"}
	}

	existing.UpdatedAt = time.Now().UTC()

	_, err = r.DB.Exec(`
UPDATE events
SET title = ?, description = ?, location = ?, start = ?, end = ?, last_modified = ?, updated_at = ?
WHERE id = ?
`,
		existing.Title,
		nullIfEmpty(existing.Description),
		nullIfEmpty(existing.Location),
		existing.Start.UTC().Format(time.RFC3339),
		existing.End.UTC().Format(time.RFC3339),
		nullTimeRFC3339(existing.LastModified),
		existing.UpdatedAt.UTC().Format(time.RFC3339),
		id,
	)
	if err != nil {
		return nil, err
	}

	return existing, nil
}

// DeleteEvent supprime un event par id.
func (r *EventRepository) DeleteEvent(id int64) error {
	res, err := r.DB.Exec(`DELETE FROM events WHERE id = ?`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err == nil && n == 0 {
		return &models.ErrorNotFound{Message: "event not found"}
	}
	return err
}

// ---- Helpers internes ----

// scanEvent lit un event depuis *sql.Rows ou *sql.Row (interface commune).
type rowScanner interface {
	Scan(dest ...any) error
}

func scanEvent(rs rowScanner) (*models.Event, error) {
	var (
		id           int64
		agendaID     string
		uid          string
		title        string
		description  sql.NullString
		location     sql.NullString
		startStr     string
		endStr       string
		lastModStr   sql.NullString
		createdAtStr string
		updatedAtStr string
	)

	if err := rs.Scan(
		&id, &agendaID, &uid, &title,
		&description, &location,
		&startStr, &endStr,
		&lastModStr,
		&createdAtStr, &updatedAtStr,
	); err != nil {
		return nil, err
	}

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(time.RFC3339, updatedAtStr)
	if err != nil {
		return nil, err
	}

	var lastMod time.Time
	if lastModStr.Valid && strings.TrimSpace(lastModStr.String) != "" {
		lastMod, err = time.Parse(time.RFC3339, lastModStr.String)
		if err != nil {
			return nil, err
		}
	}

	ev := &models.Event{
		ID:           id,
		AgendaID:     agendaID,
		UID:          uid,
		Title:        title,
		Description:  description.String,
		Location:     location.String,
		Start:        start,
		End:          end,
		LastModified: lastMod,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
	return ev, nil
}

func nullIfEmpty(s string) interface{} {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return s
}

func nullTimeRFC3339(t time.Time) interface{} {
	if t.IsZero() {
		return nil
	}
	return t.UTC().Format(time.RFC3339)
}

func isSQLiteUniqueViolation(err error) bool {
	// sqlite3 renvoie souvent un message contenant "UNIQUE constraint failed"
	return err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed")
}
