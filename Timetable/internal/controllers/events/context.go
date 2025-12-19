package events

import (
	"context"
	"fmt"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "id")
		if uid == "" {
			body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
				Message: fmt.Sprintf("missing id (uid) in path"),
			})
			w.WriteHeader(status)
			if body != nil {
				_, _ = w.Write(body)
			}
			return
		}

		ctx := context.WithValue(r.Context(), "uid", uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
