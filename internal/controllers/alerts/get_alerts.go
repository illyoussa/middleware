package alerts

import (
    "middleware/example/internal/models"
    repo "middleware/example/internal/repositories/alerts"
)

func GetAllAlerts() ([]models.Alert, error) {
    return repo.GetAllAlerts()
}

func GetAlertByID(id int64) (*models.Alert, error) {
    return repo.GetAlertByID(id)
}

func CreateAlert(a *models.Alert) (int64, error) {
    return repo.CreateAlert(a)
}

func UpdateAlert(a models.Alert) error {
    return repo.UpdateAlert(a)
}

func DeleteAlert(id int64) error {
    return repo.DeleteAlert(id)
}