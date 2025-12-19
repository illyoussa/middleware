package events

import (
	"encoding/json"
	"fmt"
	"io"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/events"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateEvent
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error reading request body: %s", err.Error())
		resp, status := helpers.RespondError(&models.ErrorGeneric{Message: "Something went wrong"})
		w.WriteHeader(status)
		if resp != nil {
			_, _ = w.Write(resp)
		}
		return
	}
	defer r.Body.Close()

	var eventToCreate models.Event
	if err := json.Unmarshal(body, &eventToCreate); err != nil {
		resp, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "Invalid JSON format"})
		w.WriteHeader(status)
		if resp != nil {
			_, _ = w.Write(resp)
		}
		return
	}

	fmt.Println(eventToCreate.AgendaID)

	//agendaId, _ := r.Context().Value("agendaId").(string)
	//eventToCreate.AgendaID = agendaId

	if err := services.CreateEvent(&eventToCreate); err != nil {
		resp, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if resp != nil {
			_, _ = w.Write(resp)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(eventToCreate)
}
