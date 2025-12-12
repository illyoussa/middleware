package events

import (
	"encoding/json"
	"io"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events" // Assurez-vous d'importer votre service
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateEvent
// @Tags         events
// @Summary      Update an event
// @Description  Update an event by its UUID
// @Accept       json
// @Produce      json
// @Param        id             path      string         true  "Event UUID formatted ID"
// @Param        event         body      models.Event  true  "Event data to update (name, ical_url)"
// @Success      200            {object}  models.Event
// @Failure      400            "Invalid JSON"
// @Failure      404            "Event not found"
// @Failure      500            "Something went wrong"
// @Router       /events/{id} [put]
// (dans votre package controllers/events)

// UpdateEvent
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, ok := ctx.Value("eventId").(uuid.UUID)
	if !ok {
		errResp := &models.ErrorGeneric{Message: "Invalid event ID in context"}
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

	var eventToUpdate models.Event
	err = json.Unmarshal(body, &eventToUpdate)
	if err != nil {
		errResp := &models.ErrorGeneric{Message: "Invalid JSON format"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	eventToUpdate.Id = &eventId

	err = services.UpdateEvent(&eventToUpdate)
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
	json.NewEncoder(w).Encode(eventToUpdate)
}
