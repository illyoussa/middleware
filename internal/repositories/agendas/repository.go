package agendas

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllAgendas() ([]models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM agendas")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	agendas := []models.Agenda{}
	for rows.Next() {
		var data models.Agenda
		err = rows.Scan(&data.Id, &data.Name, &data.IcalURL)
		if err != nil {
			return nil, err
		}
		agendas = append(agendas, data)
	}
	_ = rows.Close()

	return agendas, err
}

func GetAgendaById(id uuid.UUID) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM agendas WHERE id=?", id.String())
	helpers.CloseDB(db)

	var agenda models.Agenda
	err = row.Scan(&agenda.Id, &agenda.Name, &agenda.IcalURL)
	if err != nil {
		return nil, err
	}
	return &agenda, err
}

func CreateAgenda(a *models.Agenda) (*uuid.UUID, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO agendas(id, name, ical_url) VALUES (?, ?, ?)",
		a.Id, a.Name, a.IcalURL)

	if err != nil {
		return nil, err
	}

	return a.Id, nil
}

func UpdateAgenda(a *models.Agenda) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE agendas SET name = ?, ical_url = ? WHERE id = ?",
		a.Name, a.IcalURL, a.Id)

	return err
}

func DeleteAgenda(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM agendas WHERE id = ?", id)
	return err
}
