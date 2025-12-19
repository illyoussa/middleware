package events

import (
	"middleware/example/internal/helpers"
	services "middleware/example/internal/services/events"
	"net/http"
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

	agendaId, _ := ctx.Value("agendaId").(string)
	uid, _ := ctx.Value("uid").(string)

	err := services.DeleteEvent(agendaId, uid)
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
