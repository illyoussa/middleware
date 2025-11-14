package events

import (
    "encoding/json"
    "net/http"

    svc "middleware/example/internal/services/events"

    "github.com/go-chi/chi/v5"
    "middleware/example/internal/models"
)

func Router() chi.Router {
    r := chi.NewRouter()
    // GET / -> ?agenda_id=ID
    r.Get("/", GetByAgenda)
    // GET /{uid}
    r.Get("/{uid}", GetByUID)
    return r
}

func GetByAgenda(w http.ResponseWriter, r *http.Request) {
    agendaID := r.URL.Query().Get("agenda_id")
    if agendaID == "" {
        http.Error(w, "agenda_id required", http.StatusBadRequest)
        return
    }
    out, err := svc.GetEventsByAgenda(agendaID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(out)
}

func GetByUID(w http.ResponseWriter, r *http.Request) {
    uid := chi.URLParam(r, "uid")
    e, err := svc.GetEventByUID(uid)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(e)
}

// Optional endpoint to upsert events could be added later
func Upsert(w http.ResponseWriter, r *http.Request) {
    var e models.Event
    if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
        http.Error(w, "invalid payload", http.StatusBadRequest)
        return
    }
    if e.UID == "" {
        http.Error(w, "uid required", http.StatusBadRequest)
        return
    }
    if err := svc.UpsertEvent(e); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}