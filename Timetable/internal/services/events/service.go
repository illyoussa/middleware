//Brouillon - il faudra le corriger plus tard

package events

import (
	"database/sql"
	"fmt"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/events"

	"github.com/sirupsen/logrus"
)

func GetEventsByAgenda(agendaID string) ([]models.Event, error) {
	events, err := repository.GetEventsByAgenda(agendaID)
	if err != nil {
		logrus.Errorf("error retrieving events for agenda %s : %s", agendaID, err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving events for agenda %s", agendaID),
		}
	}
	return events, nil
}

func GetEventByUID(uid string) (*models.Event, error) {
	event, err := repository.GetEventByUID(uid)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "event not found",
			}
		}
		logrus.Errorf("error retrieving event %s : %s", uid, err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving event %s", uid),
		}
	}
	return event, nil
}

func UpsertEvent(e models.Event) error {
	err := repository.UpsertEvent(e)
	if err != nil {
		logrus.Errorf("error upserting event %s : %s", e.UID, err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while upserting event %s", e.UID),
		}
	}
	return nil
}

func DeleteEventsByAgenda(agendaID string) error {
	err := repository.DeleteEventsByAgenda(agendaID)
	if err != nil {
		logrus.Errorf("error deleting events for agenda %s : %s", agendaID, err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while deleting events for agenda %s", agendaID),
		}
	}
	return nil
}
