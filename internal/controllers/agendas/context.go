package agendas

import (
    "encoding/json"
    "net/http"

    svc "middleware/example/internal/services/agendas"

    "github.com/go-chi/chi/v5"
    "middleware/example/internal/models"
)

func Router() chi.Router {
    r := chi.NewRouter()
    r.Get("/", GetAll)
    r.Post("/", Create)
    r.Route("/{id}", func(r chi.Router) {
        r.Get("/", GetByID)
        r.Delete("/", Delete)
    })
    return r
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    out, err := svc.GetAllAgendas()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(out)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    a, err := svc.GetAgendaByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(a)
}

func Create(w http.ResponseWriter, r *http.Request) {
    var a models.Agenda
    if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
        http.Error(w, "invalid payload", http.StatusBadRequest)
        return
    }
    if a.ID == "" {
        http.Error(w, "id required", http.StatusBadRequest)
        return
    }
    if err := svc.CreateAgenda(a); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    if err := svc.DeleteAgenda(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}