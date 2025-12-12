package router

import (
	"net/http"

	"middleware/internal/controllers"
	"middleware/internal/helpers"
)

func RegisterRoutes(mux *http.ServeMux, ec *controllers.EventController) {
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			helpers.MethodNotAllowed(w)
			return
		}
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/events", ec.HandleEvents)
	mux.HandleFunc("/events/", ec.HandleEventByID)
}
