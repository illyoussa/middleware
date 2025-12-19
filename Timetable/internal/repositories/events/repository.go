package events

import (
	"database/sql"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllEvents(agendaId string) ([]models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var rows *sql.Rows
	if agendaId == "" {
		rows, err = db.Query(`
			SELECT uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id
			FROM events
		`)
	} else {
		rows, err = db.Query(`
			SELECT uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id
			FROM events
			WHERE agenda_id = ?
		`, agendaId)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Event{}
	for rows.Next() {
		var e models.Event
		err = rows.Scan(
			&e.UID,
			&e.DTStamp,
			&e.DTStart,
			&e.DTEnd,
			&e.Summary,
			&e.Location,
			&e.Description,
			&e.Created,
			&e.LastModified,
			&e.Sequence,
			&e.AgendaID,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventByUID(agendaId string, uid string) (*models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	row := db.QueryRow(`
		SELECT uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id
		FROM events
		WHERE agenda_id = ? AND uid = ?
	`, agendaId, uid)

	var e models.Event
	err = row.Scan(
		&e.UID,
		&e.DTStamp,
		&e.DTStart,
		&e.DTEnd,
		&e.Summary,
		&e.Location,
		&e.Description,
		&e.Created,
		&e.LastModified,
		&e.Sequence,
		&e.AgendaID,
	)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func CreateEvent(e *models.Event) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec(`
		INSERT INTO events (
			uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		e.UID,
		e.DTStamp,
		e.DTStart,
		e.DTEnd,
		e.Summary,
		e.Location,
		e.Description,
		e.Created,
		e.LastModified,
		e.Sequence,
		e.AgendaID,
	)
	return err
}

func UpdateEvent(e *models.Event) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec(`
		UPDATE events
		SET dtstamp = ?, dtstart = ?, dtend = ?, summary = ?, location = ?, description = ?, created = ?, last_modified = ?, sequence = ?
		WHERE agenda_id = ? AND uid = ?
	`,
		e.DTStamp,
		e.DTStart,
		e.DTEnd,
		e.Summary,
		e.Location,
		e.Description,
		e.Created,
		e.LastModified,
		e.Sequence,
		e.AgendaID,
		e.UID,
	)
	return err
}

func DeleteEvent(agendaId string, uid string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec(`DELETE FROM events WHERE agenda_id = ? AND uid = ?`, agendaId, uid)
	return err
}
