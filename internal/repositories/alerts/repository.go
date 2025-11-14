package alerts

import (
    "database/sql"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
)

func GetAllAlerts() ([]models.Alert, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    defer helpers.CloseDB(db)

    rows, err := db.Query("SELECT id, recipient, agenda_id, condition, method FROM alerts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    out := []models.Alert{}
    for rows.Next() {
        var a models.Alert
        if err := rows.Scan(&a.ID, &a.Recipient, &a.AgendaID, &a.Condition, &a.Method); err != nil {
            return nil, err
        }
        out = append(out, a)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return out, nil
}

func GetAlertByID(id int64) (*models.Alert, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    defer helpers.CloseDB(db)

    row := db.QueryRow("SELECT id, recipient, agenda_id, condition, method FROM alerts WHERE id = ?", id)

    var a models.Alert
    if err := row.Scan(&a.ID, &a.Recipient, &a.AgendaID, &a.Condition, &a.Method); err != nil {
        if err == sql.ErrNoRows {
            return nil, err
        }
        return nil, err
    }
    return &a, nil
}

func CreateAlert(a *models.Alert) (int64, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return 0, err
    }
    defer helpers.CloseDB(db)

    res, err := db.Exec("INSERT INTO alerts(recipient, agenda_id, condition, method) VALUES (?, ?, ?, ?)",
        a.Recipient, a.AgendaID, a.Condition, a.Method)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}

func UpdateAlert(a models.Alert) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("UPDATE alerts SET recipient = ?, agenda_id = ?, condition = ?, method = ? WHERE id = ?",
        a.Recipient, a.AgendaID, a.Condition, a.Method, a.ID)
    return err
}

func DeleteAlert(id int64) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("DELETE FROM alerts WHERE id = ?", id)
    return err
}