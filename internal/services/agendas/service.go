//Brouillon - il faudra le corriger plus tard

package agendas

import (
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/agendas"
)

func GetAllAgendas() ([]models.Agenda, error) {
	agendas, err := repository.GetAllAgendas()
	if err != nil {
		logrus.Errorf("error retrieving agendas : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving agendas",
		}
	}
	return agendas, nil
}

func GetAgendaByID(id uuid.UUID) (*models.Agenda, error) {
	agenda, err := repository.GetAgendaByID(id)
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
	return agenda, nil
}

func CreateAgenda(a *models.Agenda) (uuid.UUID, error) {
	id, err := repository.CreateAgenda(a)
	if err != nil {
		logrus.Errorf("error creating agenda : %s", err.Error())
		return uuid.Nil, &models.ErrorGeneric{
			Message: "Something went wrong while creating agenda",
		}
	}
	return id, nil
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