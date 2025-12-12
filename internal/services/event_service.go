package services

import (
	"strings"

	"middleware/internal/models"
	"middleware/internal/repositories"
)

type EventService struct {
	repo *repositories.EventRepository
}

func NewEventService(repo *repositories.EventRepository) *EventService {
	return &EventService{repo: repo}
}

// InitSchema initialise le schéma DB (table events)
func (s *EventService) InitSchema() error {
	return s.repo.InitSchema()
}

// CreateEvent applique une validation métier simple puis crée l'event
func (s *EventService) CreateEvent(e *models.Event) (*models.Event, error) {
	if e == nil {
		return nil, &models.ErrorUnprocessableEntity{Message: "event is required"}
	}
	// validations "métier" simples (le repo revalide aussi, mais c'est ok d'être double)
	if strings.TrimSpace(e.AgendaID) == "" {
		return nil, &models.ErrorUnprocessableEntity{Message: "agenda_id is required"}
	}
	if strings.TrimSpace(e.UID) == "" {
		return nil, &models.ErrorUnprocessableEntity{Message: "uid is required"}
	}
	if strings.TrimSpace(e.Title) == "" {
		return nil, &models.ErrorUnprocessableEntity{Message: "title is required"}
	}

	return s.repo.CreateEvent(e)
}

// ListEvents renvoie tous les events, filtrable par agenda_id
func (s *EventService) ListEvents(agendaID string) ([]models.Event, error) {
	return s.repo.GetAllEvents(agendaID)
}

// GetEvent retourne un event par id
func (s *EventService) GetEvent(id int64) (*models.Event, error) {
	return s.repo.GetEventByID(id)
}

// UpdateEvent applique patch partiel (le controller prépare le patch proprement)
func (s *EventService) UpdateEvent(id int64, patch map[string]interface{}) (*models.Event, error) {
	// petite validation métier : si title présent et vide => 422
	if v, ok := patch["title"].(string); ok {
		if strings.TrimSpace(v) == "" {
			return nil, &models.ErrorUnprocessableEntity{Message: "title cannot be empty"}
		}
	}
	return s.repo.UpdateEvent(id, patch)
}

// DeleteEvent supprime un event
func (s *EventService) DeleteEvent(id int64) error {
	return s.repo.DeleteEvent(id)
}
