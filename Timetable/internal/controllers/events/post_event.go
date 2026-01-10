package events

import (
	"encoding/json"
	"net/http"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events"
)

// CreateEvent
// @Tags         events
// @Summary      Create a new event
// @Description  Create a new event with the provided JSON data
// @Accept       json
// @Produce      json
// @Param        event    body     models.Event true  "Event data to create"
// @Success      201      {object} models.Event
// @Failure      400      "Invalid JSON"
// @Failure      422      "Unprocessable Entity"
// @Failure      500      "Something went wrong"
// @Router       /events [post]
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// 1. Décodage du JSON reçu dans le body
	var eventToCreate models.Event

	// On utilise NewDecoder pour lire le flux directement (plus performant)
	err := json.NewDecoder(r.Body).Decode(&eventToCreate)
	if err != nil {
		// Si le JSON est mal formé
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

	// 2. Appel au Service
	err = services.CreateEvent(&eventToCreate)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// 3. Réponse Succès (201 Created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(eventToCreate)
}
