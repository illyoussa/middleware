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

// UpdateAgenda
// @Tags         agendas
// @Summary      Update an agenda
// @Description  Update an agenda by its UUID
// @Accept       json
// @Produce      json
// @Param        id             path      string         true  "Agenda UUID formatted ID"
// @Param        agenda         body      models.Agenda  true  "Agenda data to update (name, ical_url)"
// @Success      200            {object}  models.Agenda
// @Failure      400            "Invalid JSON"
// @Failure      404            "Agenda not found"
// @Failure      500            "Something went wrong"
// @Router       /agendas/{id} [put]
// (dans votre package controllers/agendas)

// UpdateAgenda
func UpdateAgenda(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agendaId, ok := ctx.Value("agendaId").(uuid.UUID)
	if !ok {
		errResp := &models.ErrorGeneric{Message: "Invalid agenda ID in context"}
		body, status := helpers.RespondError(errResp)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

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

	var agendaToUpdate models.Agenda
	err = json.Unmarshal(body, &agendaToUpdate)
	if err != nil {
		errResp := &models.ErrorGeneric{Message: "Invalid JSON format"}
		bodyResp, statusResp := helpers.RespondError(errResp)
		w.WriteHeader(statusResp)
		if bodyResp != nil {
			_, _ = w.Write(bodyResp)
		}
		return
	}

	agendaToUpdate.Id = &agendaId

	err = services.UpdateAgenda(&agendaToUpdate)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// renvoyer une réponse 200 OK avec l'objet mis à jour
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agendaToUpdate)
}
