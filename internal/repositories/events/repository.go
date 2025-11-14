package events

import (
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
)

func GetEventsByAgenda(agendaID string) ([]models.Event, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    defer helpers.CloseDB(db)

    rows, err := db.Query("SELECT uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id FROM events WHERE agenda_id = ?", agendaID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    out := []models.Event{}
    for rows.Next() {
        var e models.Event
        if err := rows.Scan(&e.UID, &e.DTStamp, &e.DTStart, &e.DTEnd, &e.Summary, &e.Location, &e.Description, &e.Created, &e.LastModified, &e.Sequence, &e.AgendaID); err != nil {
            return nil, err
        }
        out = append(out, e)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return out, nil
}

func GetEventByUID(uid string) (*models.Event, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    defer helpers.CloseDB(db)

    row := db.QueryRow("SELECT uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id FROM events WHERE uid = ?", uid)

    var e models.Event
    if err := row.Scan(&e.UID, &e.DTStamp, &e.DTStart, &e.DTEnd, &e.Summary, &e.Location, &e.Description, &e.Created, &e.LastModified, &e.Sequence, &e.AgendaID); err != nil {
        return nil, err
    }
    return &e, nil
}

// UpsertEvent creates or replaces an event by UID.
func UpsertEvent(e models.Event) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec(`INSERT OR REPLACE INTO events(
        uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        e.UID, e.DTStamp, e.DTStart, e.DTEnd, e.Summary, e.Location, e.Description, e.Created, e.LastModified, e.Sequence, e.AgendaID)
    return err
}

func DeleteEventsByAgenda(agendaID string) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("DELETE FROM events WHERE agenda_id = ?", agendaID)
    return err