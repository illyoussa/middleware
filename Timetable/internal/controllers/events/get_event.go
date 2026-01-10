package events

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events"
)

// GetEvent
// @Tags         events
// @Summary      Get a specific event
// @Description  Retrieve a single event by its UID
// @Produce      json
// @Param        id   path      string  true  "Event UID"
// @Success      200  {object}  models.Event
// @Failure      404  "Event not found"
// @Failure      500  "Something went wrong"
// @Router       /events/{id} [get]
func GetEvent(w http.ResponseWriter, r *http.Request) {
	// On récupère "id" depuis l'URL
	id := chi.URLParam(r, "id")

	if id == "" {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "Missing ID"})
		w.WriteHeader(status)
		w.Write(body)
		return
	}

	// Appel au service (qui appellera le repository GetEventById)
	event, err := services.GetEventById(id)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		w.Write(body)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}
