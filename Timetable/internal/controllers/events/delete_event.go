package events

import (
	"net/http"

	"github.com/go-chi/chi/v5" // Import nécessaire pour chi.URLParam

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events"
)

// DeleteEvent
// @Tags         events
// @Summary      Delete an event
// @Description  Delete an event by its ID
// @Param        id             path      string  true  "Event ID"
// @Success      204            "No Content"
// @Failure      400            "Invalid ID"
// @Failure      404            "Event not found"
// @Failure      500            "Something went wrong"
// @Router       /events/{id} [delete]
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	// 1. Récupération de l'ID depuis l'URL
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

	err := services.DeleteEvent(id)
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
