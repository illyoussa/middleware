package events

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events"
	"net/http"
)

// GetEvent
func GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uid, ok := ctx.Value("uid").(string)
	if !ok || uid == "" {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "Invalid event UID in context"})
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	agendaId := r.URL.Query().Get("agendaId")
	if agendaId == "" {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "Missing agendaId query parameter"})
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	event, err := services.GetEventByUID(agendaId, uid)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(event)
}
