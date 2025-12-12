package events

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/events" // Assurez-vous d'importer votre service
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteEvent
// @Tags         events
// @Summary      Delete an event
// @Description  Delete an event by its UUID
// @Param        id             path      string  true  "Event UUID formatted ID"
// @Success      204            "No Content"
// @Failure      404            "Event not found"
// @Failure      500            "Something went wrong"
// @Router       /events/{id} [delete]
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, _ := ctx.Value("eventId").(uuid.UUID)

	err := events.DeleteEvent(eventId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
