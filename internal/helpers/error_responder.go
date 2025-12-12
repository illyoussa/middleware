package helpers

import (
	"encoding/json"
	"net/http"

	"middleware/internal/models"

	"github.com/sirupsen/logrus"
)

// RespondError transforme une erreur métier en réponse HTTP (status + body JSON)
func RespondError(err error) (body []byte, status int) {
	// Par défaut : erreur interne
	status = http.StatusInternalServerError

	switch err.(type) {
	case *models.ErrorNotFound:
		status = http.StatusNotFound
	case *models.ErrorUnprocessableEntity:
		status = http.StatusUnprocessableEntity
	}

	// On ne renvoie le body que si ce n'est PAS une 500
	// (bonne pratique : ne pas leak les erreurs internes)
	if status != http.StatusInternalServerError {
		body, _ = json.Marshal(err)
	}

	// Log serveur (visible seulement côté backend)
	logrus.WithError(err).Warnf("HTTP error %d", status)

	return
}
