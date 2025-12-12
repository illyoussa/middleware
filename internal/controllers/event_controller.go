package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"middleware/internal/helpers"
	"middleware/internal/models"
	"middleware/internal/services"
)

type EventController struct {
	service *services.EventService
}

func NewEventController(svc *services.EventService) *EventController {
	return &EventController{service: svc}
}

// HandleEvents gère /events (GET, POST)
func (ec *EventController) HandleEvents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ec.listEvents(w, r)
	case http.MethodPost:
		ec.createEvent(w, r)
	default:
		helpers.MethodNotAllowed(w)
	}
}

// HandleEventByID gère /events/{id} (GET, PATCH, DELETE)
func (ec *EventController) HandleEventByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/events/")
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	switch r.Method {
	case http.MethodGet:
		ec.getEvent(w, r, id)
	case http.MethodPatch:
		ec.patchEvent(w, r, id)
	case http.MethodDelete:
		ec.deleteEvent(w, r, id)
	default:
		helpers.MethodNotAllowed(w)
	}
}

// --- sous-fonctions ---

func (ec *EventController) listEvents(w http.ResponseWriter, r *http.Request) {
	agendaID := strings.TrimSpace(r.URL.Query().Get("agenda_id"))

	events, err := ec.service.ListEvents(agendaID)
	if err != nil {
		respondWithBusinessError(w, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, events)
}

func (ec *EventController) createEvent(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AgendaID     string    `json:"agenda_id"`
		UID          string    `json:"uid"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		Location     string    `json:"location"`
		Start        time.Time `json:"start"`
		End          time.Time `json:"end"`
		LastModified time.Time `json:"last_modified"`
	}

	if err := helpers.DecodeJSON(r, &input); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "invalid json", err.Error())
		return
	}

	ev := &models.Event{
		AgendaID:     strings.TrimSpace(input.AgendaID),
		UID:          strings.TrimSpace(input.UID),
		Title:        strings.TrimSpace(input.Title),
		Description:  input.Description,
		Location:     input.Location,
		Start:        input.Start,
		End:          input.End,
		LastModified: input.LastModified,
	}

	created, err := ec.service.CreateEvent(ev)
	if err != nil {
		respondWithBusinessError(w, err)
		return
	}

	w.Header().Set("Location", "/events/"+strconv.FormatInt(created.ID, 10))
	helpers.WriteJSON(w, http.StatusCreated, created)
}

func (ec *EventController) getEvent(w http.ResponseWriter, r *http.Request, id int64) {
	ev, err := ec.service.GetEvent(id)
	if err != nil {
		respondWithBusinessError(w, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, ev)
}

func (ec *EventController) patchEvent(w http.ResponseWriter, r *http.Request, id int64) {
	// Patch: on accepte uniquement certains champs
	var raw map[string]any
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "invalid json", err.Error())
		return
	}

	patch := make(map[string]interface{})

	// title
	if v, ok := raw["title"]; ok {
		if s, ok := v.(string); ok {
			patch["title"] = strings.TrimSpace(s)
		} else {
			helpers.WriteError(w, http.StatusUnprocessableEntity, "title must be a string", nil)
			return
		}
	}

	// description
	if v, ok := raw["description"]; ok {
		if s, ok := v.(string); ok {
			patch["description"] = s
		} else {
			helpers.WriteError(w, http.StatusUnprocessableEntity, "description must be a string", nil)
			return
		}
	}

	// location
	if v, ok := raw["location"]; ok {
		if s, ok := v.(string); ok {
			patch["location"] = s
		} else {
			helpers.WriteError(w, http.StatusUnprocessableEntity, "location must be a string", nil)
			return
		}
	}

	// start/end/last_modified (attendus en RFC3339 string)
	if v, ok := raw["start"]; ok {
		t, err := parseRFC3339Field(v, "start")
		if err != nil {
			helpers.WriteError(w, http.StatusUnprocessableEntity, err.Error(), nil)
			return
		}
		patch["start"] = t
	}
	if v, ok := raw["end"]; ok {
		t, err := parseRFC3339Field(v, "end")
		if err != nil {
			helpers.WriteError(w, http.StatusUnprocessableEntity, err.Error(), nil)
			return
		}
		patch["end"] = t
	}
	if v, ok := raw["last_modified"]; ok {
		t, err := parseRFC3339Field(v, "last_modified")
		if err != nil {
			helpers.WriteError(w, http.StatusUnprocessableEntity, err.Error(), nil)
			return
		}
		patch["last_modified"] = t
	}

	updated, err := ec.service.UpdateEvent(id, patch)
	if err != nil {
		respondWithBusinessError(w, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, updated)
}

func (ec *EventController) deleteEvent(w http.ResponseWriter, r *http.Request, id int64) {
	if err := ec.service.DeleteEvent(id); err != nil {
		respondWithBusinessError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- helpers locaux controller ---

func respondWithBusinessError(w http.ResponseWriter, err error) {
	body, status := helpers.RespondError(err)
	if status == http.StatusInternalServerError {
		helpers.WriteError(w, status, "internal server error", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func parseIDFromPath(path, prefix string) (int64, error) {
	if !strings.HasPrefix(path, prefix) {
		return 0, strconv.ErrSyntax
	}
	idStr := strings.TrimPrefix(path, prefix)
	// Sécurité: pas d'autres segments
	if strings.Contains(idStr, "/") || idStr == "" {
		return 0, strconv.ErrSyntax
	}
	return strconv.ParseInt(idStr, 10, 64)
}

func parseRFC3339Field(v any, field string) (time.Time, error) {
	s, ok := v.(string)
	if !ok {
		return time.Time{}, &models.ErrorUnprocessableEntity{Message: field + " must be a RFC3339 string"}
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, &models.ErrorUnprocessableEntity{Message: field + " must be RFC3339 (e.g. 2025-12-12T10:00:00Z)"}
	}
	return t, nil
}
