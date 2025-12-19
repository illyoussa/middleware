package events

import (
	"encoding/json"
	"io"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events"
	"net/http"

	"github.com/sirupsen/logrus"
)

// UpdateEvent
// @Tags         events
// @Summary      Update an event
// @Description  Update an event by its UID (string)
// @Accept       json
// @Produce      json
// @Param        id       path     string       true  "Event UID (string)"
// @Param        agendaId query    string       true  "Agenda ID (string)"
// @Param        event    body     models.Event true  "Event data to update"
// @Success      200      {object} models.Event
// @Failure      400      "Invalid JSON"
// @Failure      404      "Event not found"
// @Failure      500      "Something went wrong"
// @Router       /events/{id} [put]
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// uid vient du Context middleware (context.go)
	uid, ok := ctx.Value("uid").(string)
	if !ok || uid == "" {
		errResp := &models.ErrorUnprocessableEntity{Message: "Invalid event UID in context"}
		body, status := helpers.RespondError(errResp)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// agendaId obligatoire (car repo/service utilisent (agendaId, uid) comme clé logique)
	agendaId := r.URL.Query().Get("agendaId")
	if agendaId == "" {
		errResp := &models.ErrorUnprocessableEntity{Message: "Missing agendaId query parameter"}
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
	if err := json.Unmarshal(body, &eventToUpdate); err != nil {
		errResp := &models.ErrorUnprocessableEntity{Message: "Invalid JSON format"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	// On force la clé depuis l'URL (comme Config force l'Id)
	eventToUpdate.UID = uid
	eventToUpdate.AgendaID = agendaId

	if err := services.UpdateEvent(&eventToUpdate); err != nil {
		resp, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if resp != nil {
			_, _ = w.Write(resp)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(eventToUpdate)
}
