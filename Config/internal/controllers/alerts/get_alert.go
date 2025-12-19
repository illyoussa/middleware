package alerts

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/repositories/alerts"
	"net/http"

	"github.com/gofrs/uuid"
)

// GetAlert
// @Tags         alerts
// @Summary      Get a alerts.
// @Description  Get a alerts.
// @Param        id           	path      string  true  "Alerts UUID formatted ID"
// @Success      200            {object}  models.Alert
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /alerts/{id} [get]
func GetAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, _ := ctx.Value("alertId").(uuid.UUID) // getting key set in context.go

	alert, err := alerts.GetAlertById(alertId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alert)
	_, _ = w.Write(body)
	return
}
