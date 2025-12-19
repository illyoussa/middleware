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

// UpdateAlert
// @Tags         alerts
// @Summary      Update an alert
// @Description  Update an alert by its UUID
// @Accept       json
// @Produce      json
// @Param        id             path      string         true  "Alert UUID formatted ID"
// @Param        alert         body      models.Alert  true  "Alert data to update (name, ical_url)"
// @Success      200            {object}  models.Alert
// @Failure      400            "Invalid JSON"
// @Failure      404            "Alert not found"
// @Failure      500            "Something went wrong"
// @Router       /alerts/{id} [put]
// (dans votre package controllers/alerts)

// UpdateAlert
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, ok := ctx.Value("alertId").(uuid.UUID)
	if !ok {
		errResp := &models.ErrorGeneric{Message: "Invalid alert ID in context"}
		body, status := helpers.RespondError(errResp)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

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

	var alertToUpdate models.Alert
	err = json.Unmarshal(body, &alertToUpdate)
	if err != nil {
		errResp := &models.ErrorGeneric{Message: "Invalid JSON format"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	alertToUpdate.Id = &alertId

	err = services.UpdateAlert(&alertToUpdate)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// renvoyer une réponse 200 OK avec l'objet mis à jour
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alertToUpdate)
}
