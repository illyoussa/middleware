package agendas

import (
    "middleware/example/internal/models"
    repo "middleware/example/internal/repositories/agendas"
)

func GetAllAgendas() ([]models.Agenda, error) {
    return repo.GetAllAgendas()
}

func GetAgendaByID(id string) (*models.Agenda, error) {
    return repo.GetAgendaByID(id)
}

func CreateAgenda(a models.Agenda) error {
    return repo.CreateAgenda(a)
}

func DeleteAgenda(id string) error {
    return repo.DeleteAgenda(id)
}