package alerts

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllAlerts() ([]models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM alerts")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	alerts := []models.Alert{}
	for rows.Next() {
		var a models.Alert
		err = rows.Scan(&a.Id, &a.Recipient, &a.AgendaID, &a.Condition, &a.Method)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, a)
	}

	// don't forget to close rows
	_ = rows.Close()

	return alerts, err
}

func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT id, recipient, agenda_id, condition, method FROM alerts WHERE id = ?", id)
	helpers.CloseDB(db)

	var alert models.Alert
	err = row.Scan(&alert.Id, &alert.Recipient, &alert.AgendaID, &alert.Condition, &alert.Method)
	if err != nil {
		return nil, err
	}
	return &alert, err
}

func CreateAlert(a *models.Alert) (*uuid.UUID, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Vérifier ou générer l'ID
	if a.Id == nil {
		newId, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		a.Id = &newId
	}

	// Insérer l'alerte en incluant l'ID
	_, err = db.Exec(
		"INSERT INTO alerts(id, recipient, agenda_id, condition, method) VALUES (?, ?, ?, ?, ?)",
		a.Id.String(), a.Recipient, a.AgendaID, a.Condition, a.Method,
	)
	if err != nil {
		return nil, err
	}

	return a.Id, nil
}

func UpdateAlert(a *models.Alert) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE alerts SET recipient = ?, agenda_id = ?, condition = ?, method = ? WHERE id = ?",
		a.Recipient, a.AgendaID, a.Condition, a.Method, a.Id)
	return err
}

func DeleteAlert(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM alerts WHERE id = ?", id)
	return err
}
