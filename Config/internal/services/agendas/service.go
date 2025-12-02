package agendas

import (
	"database/sql"
	"fmt"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/agendas"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllAgendas() ([]models.Agenda, error) {
	var err error
	// calling repository
	agendas, err := repository.GetAllAgendas()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving agendas : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving agendas",
		}
	}

	return agendas, nil
}

func GetAgendaById(id uuid.UUID) (*models.Agenda, error) {
	agenda, err := repository.GetAgendaById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "agenda not found",
			}
		}
		logrus.Errorf("error retrieving agenda %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving agenda %s", id.String()),
		}
	}

	return agenda, err
}

func CreateAgenda(a *models.Agenda) (*uuid.UUID, error) {
	id, err := repository.CreateAgenda(a)

	if err != nil {
		logrus.Errorf("error creating agenda: %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while creating agenda",
		}
	}
	return id, nil
}

func UpdateAgenda(a *models.Agenda) error {
	err := repository.UpdateAgenda(a)
	if err != nil {
		logrus.Errorf("error updating agenda %s : %s", a.Id.String(), err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while updating agenda %s", a.Id.String()),
		}
	}
	return nil
}

func DeleteAgenda(id uuid.UUID) error {
	err := repository.DeleteAgenda(id)
	if err != nil {
		logrus.Errorf("error deleting agenda %s : %s", id.String(), err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while deleting agenda %s", id.String()),
		}
	}
	return nil
}
