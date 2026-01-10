package events

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	services "middleware/example/internal/services/events"
	"net/http"
)

// GetEvents
// @Tags         events
// @Summary      Get all events
// @Description  Retrieve the list of all existing events
// @Produce      json
// @Success      200      {array}  models.Event
// @Failure      500      "Something went wrong"
// @Router       /events [get]
func GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := services.GetAllEvents()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(events)
	_, _ = w.Write(body)
	return
}
