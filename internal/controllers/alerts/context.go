package alerts

import (
    "encoding/json"
    "net/http"
    "strconv"

    svc "middleware/example/internal/services/alerts"

    "github.com/go-chi/chi/v5"
    "middleware/example/internal/models"
)

func Router() chi.Router {
    r := chi.NewRouter()
    r.Get("/", GetAll)
    r.Post("/", Create)
    r.Route("/{id}", func(r chi.Router) {
        r.Get("/", GetByID)
        r.Put("/", Update)
        r.Delete("/", Delete)
    })
    return r
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    out, err := svc.GetAllAlerts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(out)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
    idp := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idp, 10, 64)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }
    a, err := svc.GetAlertByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(a)
}

func Create(w http.ResponseWriter, r *http.Request) {
    var a models.Alert
    if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
        http.Error(w, "invalid payload", http.StatusBadRequest)
        return
    }
    if a.Recipient == "" || a.AgendaID == "" {
        http.Error(w, "recipient and agendaId required", http.StatusBadRequest)
        return
    }
    id, err := svc.CreateAlert(&a)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func Update(w http.ResponseWriter, r *http.Request) {
    idp := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idp, 10, 64)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }
    var a models.Alert
    if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
        http.Error(w, "invalid payload", http.StatusBadRequest)
        return
    }
    a.ID = id
    if err := svc.UpdateAlert(a); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    idp := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idp, 10, 64)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }
    if err := svc.DeleteAlert(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}