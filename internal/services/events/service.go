package events

import (
	"database/sql"
	"fmt"

	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/events"

	"github.com/sirupsen/logrus"
)

// GetAllEvents retourne tous les events (optionnellement filtrés par agendaId)
func GetAllEvents(agendaId string) ([]models.Event, error) {
	events, err := repository.GetAllEvents(agendaId)
	if err != nil {
		logrus.Errorf("error retrieving events : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving events",
		}
	}

	return events, nil
}

// GetEventByUID retourne un event via (agendaId, uid)
func GetEventByUID(agendaId string, uid string) (*models.Event, error) {
	event, err := repository.GetEventByUID(agendaId, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.ErrorNotFound{
				Message: "event not found",
			}
		}
		logrus.Errorf(
			"error retrieving event uid=%s agendaId=%s : %s",
			uid,
			agendaId,
			err.Error(),
		)
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving event",
		}
	}

	return event, nil
}

// CreateEvent crée un nouvel event
func CreateEvent(e *models.Event) error {
	// TODO uuid, _ = uuid.NewUUID()
	err := repository.CreateEvent(e)
	if err != nil {
		logrus.Errorf(
			"error creating event uid=%s agendaId=%s : %s",
			e.UID,
			e.AgendaID,
			err.Error(),
		)
		return &models.ErrorGeneric{
			Message: "Something went wrong while creating event",
		}
	}
	return nil
}

// UpdateEvent met à jour un event existant
func UpdateEvent(e *models.Event) error {
	err := repository.UpdateEvent(e)
	if err != nil {
		logrus.Errorf(
			"error updating event uid=%s agendaId=%s : %s",
			e.UID,
			e.AgendaID,
			err.Error(),
		)
		return &models.ErrorGeneric{
			Message: "Something went wrong while updating event",
		}
	}
	return nil
}

// DeleteEvent supprime un event via (agendaId, uid)
func DeleteEvent(agendaId string, uid string) error {
	err := repository.DeleteEvent(agendaId, uid)
	if err != nil {
		logrus.Errorf(
			"error deleting event uid=%s agendaId=%s : %s",
			uid,
			agendaId,
			err.Error(),
		)
		return &models.ErrorGeneric{
			Message: fmt.Sprintf(
				"Something went wrong while deleting event uid=%s agendaId=%s",
				uid,
				agendaId,
			),
		}
	}
	return nil
}
