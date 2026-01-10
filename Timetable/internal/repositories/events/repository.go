package events

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllEvents() ([]models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	query := `
        SELECT resource_ids, id, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id 
        FROM events`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Event{}

	for rows.Next() {
		var data models.Event
		var resourceIdsJSON []byte

		err = rows.Scan(
			&resourceIdsJSON,
			&data.Id,
			&data.DTStamp,
			&data.DTStart,
			&data.DTEnd,
			&data.Summary,
			&data.Location,
			&data.Description,
			&data.Created,
			&data.LastModified,
			&data.Sequence,
			&data.AgendaID,
		)
		if err != nil {
			return nil, err
		}

		if len(resourceIdsJSON) > 0 {
			_ = json.Unmarshal(resourceIdsJSON, &data.ResourceIds)
		}
		if data.ResourceIds == nil {
			data.ResourceIds = []int{}
		}

		events = append(events, data)
	}

	return events, nil
}

func GetEventById(id string) (*models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	query := `
        SELECT resource_ids, id, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id 
        FROM events 
        WHERE id = ?`

	row := db.QueryRow(query, id)

	var event models.Event
	var resourceIdsJSON []byte

	err = row.Scan(
		&resourceIdsJSON,
		&event.Id,
		&event.DTStamp,
		&event.DTStart,
		&event.DTEnd,
		&event.Summary,
		&event.Location,
		&event.Description,
		&event.Created,
		&event.LastModified,
		&event.Sequence,
		&event.AgendaID,
	)

	if err != nil {
		return nil, err
	}

	if len(resourceIdsJSON) > 0 {
		_ = json.Unmarshal(resourceIdsJSON, &event.ResourceIds)
	}
	if event.ResourceIds == nil {
		event.ResourceIds = []int{}
	}

	return &event, nil
}

func CreateEvent(e *models.Event) (string, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return "", err
	}
	defer helpers.CloseDB(db)

	if e.Id == "" {
		newID, err := uuid.NewV4()
		if err != nil {
			return "", err
		}
		e.Id = newID.String()
	}

	if e.ResourceIds == nil {
		e.ResourceIds = []int{}
	}
	resourceIdsJSON, err := json.Marshal(e.ResourceIds)
	if err != nil {
		return "", err
	}

	query := `
        INSERT INTO events (
            id, resource_ids, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(query,
		e.Id,
		string(resourceIdsJSON),
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

	if err != nil {
		return "", err
	}

	return e.Id, nil
}

func UpdateEvent(e *models.Event) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	if e.ResourceIds == nil {
		e.ResourceIds = []int{}
	}
	resourceIdsJSON, err := json.Marshal(e.ResourceIds)
	if err != nil {
		return err
	}

	query := `
        UPDATE events
        SET resource_ids = ?, dtstamp = ?, dtstart = ?, dtend = ?, summary = ?, location = ?, description = ?, created = ?, last_modified = ?, sequence = ?
        WHERE id = ?`

	_, err = db.Exec(query,
		string(resourceIdsJSON),
		e.DTStamp,
		e.DTStart,
		e.DTEnd,
		e.Summary,
		e.Location,
		e.Description,
		e.Created,
		e.LastModified,
		e.Sequence,
		e.Id,
	)

	return err
}

// DeleteEvent
func DeleteEvent(id string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	query := "DELETE FROM events WHERE id = ?"

	_, err = db.Exec(query, id)
	return err
}
