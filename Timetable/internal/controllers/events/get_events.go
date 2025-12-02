package events

import (
    "middleware/example/internal/models"
    repo "middleware/example/internal/repositories/events"
)

func GetEventsByAgenda(agendaID string) ([]models.Event, error) {
    return repo.GetEventsByAgenda(agendaID)
}

func GetEventByUID(uid string) (*models.Event, error) {
    return repo.GetEventByUID(uid)
}

func UpsertEvent(e models.Event) error {
    return repo.UpsertEvent(e)
}

func DeleteEventsByAgenda(agendaID string) error {
    return repo.DeleteEventsByAgenda(agendaID)
}