package agendas

import (
	"encoding/json"
	"io"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/agendas" // Assurez-vous d'importer votre service
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// CreateAgenda
// @Tags         agendas
// @Summary      Create a new agenda
// @Description  Create a new agenda with data
// @Accept       json
// @Produce      json
// @Param        agenda         body      models.Agenda  true  "Agenda data (name and ical_url)"
// @Success      201            {object}  models.Agenda
// @Failure      400            "Bad request - Invalid JSON"
// @Failure      500            "Something went wrong"
// @Router       /agendas [post]
func CreateAgenda(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error reading request body: %s", err.Error())
		errResp := &models.ErrorGeneric{Message: "Something went wrong"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}
	defer r.Body.Close()

	var agendaToCreate models.Agenda
	err = json.Unmarshal(body, &agendaToCreate)
	if err != nil {
		errResp := &models.ErrorGeneric{Message: "Invalid JSON format"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp) // 400
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	newId, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Error generating UUID: %s", err.Error())
		errResp := &models.ErrorGeneric{Message: "Something went wrong"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	agendaToCreate.Id = &newId

	returnedId, err := services.CreateAgenda(&agendaToCreate)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	agendaToCreate.Id = returnedId

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 pour une création réussie
	json.NewEncoder(w).Encode(agendaToCreate)
}
