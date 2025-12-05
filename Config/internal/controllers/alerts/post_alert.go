package alerts

import (
	"encoding/json"
	"io"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/alerts" // Assurez-vous d'importer votre service
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// CreateAlert
// @Tags         alerts
// @Summary      Create a new alert
// @Description  Create a new alert with data
// @Accept       json
// @Produce      json
// @Param        alert         body      models.Alert  true  "Alert data (name and ical_url)"
// @Success      201            {object}  models.Alert
// @Failure      400            "Bad request - Invalid JSON"
// @Failure      500            "Something went wrong"
// @Router       /alerts [post]
func CreateAlert(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error reading request body: %s", err.Error())
		errResp := &models.ErrorGeneric{Message: "Something went wrong"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}
	defer r.Body.Close()

	var alertToCreate models.Alert
	err = json.Unmarshal(body, &alertToCreate)
	if err != nil {
		errResp := &models.ErrorGeneric{Message: "Invalid JSON format"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp) // 400
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	newId, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Error generating UUID: %s", err.Error())
		errResp := &models.ErrorGeneric{Message: "Something went wrong"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	alertToCreate.Id = &newId

	returnedId, err := services.CreateAlert(&alertToCreate)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	alertToCreate.Id = returnedId

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 pour une création réussie
	json.NewEncoder(w).Encode(alertToCreate)
}