package events

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	services "middleware/example/internal/services/events"
	"net/http"
)

// GetEvents
func GetEvents(w http.ResponseWriter, r *http.Request) {
	agendaId, _ := r.Context().Value("agendaId").(string)

	events, err := services.GetAllEvents(agendaId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(events)
}
