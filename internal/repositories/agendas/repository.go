package agendas

import (
    "database/sql"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
)

func GetAllAgendas() ([]models.Agenda, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    defer helpers.CloseDB(db)

    rows, err := db.Query("SELECT id, name, ical_url FROM agendas")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    out := []models.Agenda{}
    for rows.Next() {
        var a models.Agenda
        if err := rows.Scan(&a.ID, &a.Name, &a.IcalURL); err != nil {
            return nil, err
        }
        out = append(out, a)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return out, nil
}

func GetAgendaByID(id string) (*models.Agenda, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    defer helpers.CloseDB(db)

    row := db.QueryRow("SELECT id, name, ical_url FROM agendas WHERE id = ?", id)

    var a models.Agenda
    if err := row.Scan(&a.ID, &a.Name, &a.IcalURL); err != nil {
        if err == sql.ErrNoRows {
            return nil, err
        }
        return nil, err
    }
    return &a, nil
}

func CreateAgenda(a models.Agenda) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("INSERT INTO agendas(id, name, ical_url) VALUES (?, ?, ?)", a.ID, a.Name, a.IcalURL)
    return err
}

func DeleteAgenda(id string) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("DELETE FROM agendas WHERE id = ?", id)
    return err
}