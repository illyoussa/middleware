//Brouillon - il faudra le corriger plus tard

package alerts

import (
	"database/sql"
	"fmt"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/alerts"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllAlerts() ([]models.Alert, error) {
	var err error
	// calling repository
	alerts, err := repository.GetAllAlerts()
	if err != nil {
		logrus.Errorf("error retrieving alerts : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving alerts",
		}
	}
	return alerts, nil
}

func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	alert, err := repository.GetAlertById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "alert not found",
			}
		}
		logrus.Errorf("error retrieving alert %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving alert %s", id.String()),
		}
	}
	return alert, err
}

func CreateAlert(a *models.Alert) (*uuid.UUID, error) {
	id, err := repository.CreateAlert(a)

	if err != nil {
		logrus.Errorf("error creating alert: %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while creating alert",
		}
	}
	return id, nil
}

func UpdateAlert(a *models.Alert) error {
	err := repository.UpdateAlert(a)
	if err != nil {
		logrus.Errorf("error updating alert %s : %s", a.Id.String(), err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while updating alert %s", a.Id.String()),
		}
	}
	return nil
}
