package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

// WriteJSON écrit une réponse JSON standard
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// WriteError écrit une erreur JSON simple
func WriteError(w http.ResponseWriter, status int, message string, details interface{}) {
	WriteJSON(w, status, map[string]interface{}{
		"error":   message,
		"details": details,
	})
}

// DecodeJSON décode le body JSON dans dst
func DecodeJSON(r *http.Request, dst interface{}) error {
	if r.Body == nil {
		return errors.New("empty body")
	}
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}

// MethodNotAllowed renvoie une 405
func MethodNotAllowed(w http.ResponseWriter) {
	WriteError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
}
