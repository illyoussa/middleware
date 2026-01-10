package events

import (
	"fmt"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/events"

	"github.com/sirupsen/logrus"
)

func GetAllEvents() ([]models.Event, error) {
	var err error
	events, err := repository.GetAllEvents()
	if err != nil {
		logrus.Errorf("error retrieving events : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving events",
		}
	}

	return events, nil
}

func GetEventById(id string) (*models.Event, error) {
	event, err := repository.GetEventById(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, &models.ErrorNotFound{
				Message: "event not found",
			}
		}
		logrus.Errorf("error retrieving event %s : %s", id, err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving event " + id,
		}
	}

	return event, nil
}

func CreateEvent(e *models.Event) error {
	_, err := repository.CreateEvent(e)
	if err != nil {
		logrus.Errorf("error creating event id=%s : %s", e.Id, err.Error())
		return &models.ErrorGeneric{
			Message: "Something went wrong while creating event",
		}
	}
	return nil
}

func UpdateEvent(e *models.Event) error {
	err := repository.UpdateEvent(e)
	if err != nil {
		logrus.Errorf(
			"error updating event id=%s agendaId=%s : %s",
			e.Id,
			e.AgendaID,
			err.Error(),
		)
		return &models.ErrorGeneric{
			Message: "Something went wrong while updating event",
		}
	}
	return nil
}

func DeleteEvent(id string) error {
	err := repository.DeleteEvent(id)
	if err != nil {
		logrus.Errorf(
			"error deleting event id=%s : %s",
			id,
			err.Error(),
		)
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while deleting event id=%s", id),
		}
	}
	return nil
}
