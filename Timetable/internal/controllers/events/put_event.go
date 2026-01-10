package events

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5" // Important pour récupérer l'ID dans l'URL

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events"
)

// UpdateEvent
// @Tags         events
// @Summary      Update an event
// @Description  Update an event by its ID (string)
// @Accept       json
// @Produce      json
// @Param        id       path     string       true  "Event ID (string)"
// @Param        agendaId query    string       true  "Agenda ID (string)"
// @Param        event    body     models.Event true  "Event data to update"
// @Success      200      {object} models.Event
// @Failure      400      "Invalid JSON"
// @Failure      404      "Event not found"
// @Failure      422      "Unprocessable Entity"
// @Failure      500      "Something went wrong"
// @Router       /events/{id} [put]
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
			Message: "Missing Event ID in URL",
		})
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	agendaID := r.URL.Query().Get("agendaId")
	if agendaID == "" {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
			Message: "Missing agendaId query parameter",
		})
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	var eventToUpdate models.Event
	if err := json.NewDecoder(r.Body).Decode(&eventToUpdate); err != nil {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
			Message: "Invalid JSON format",
		})
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}
	defer r.Body.Close()

	eventToUpdate.Id = id
	eventToUpdate.AgendaID = agendaID

	if err := services.UpdateEvent(&eventToUpdate); err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(eventToUpdate)
}
