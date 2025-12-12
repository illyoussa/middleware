package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"middleware/internal/controllers"
	"middleware/internal/helpers"
	"middleware/internal/repositories"
	"middleware/internal/router"
	"middleware/internal/services"
)

func main() {
	// --- Logger ---
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// --- DB ---
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("failed to open db: %s", err.Error())
	}
	defer helpers.CloseDB(db)

	// --- Wiring (repo -> service -> controller) ---
	repo := repositories.NewEventRepository(db)
	svc := services.NewEventService(repo)
	ctrl := controllers.NewEventController(svc)

	// --- Init schema (create tables) ---
	if err := svc.InitSchema(); err != nil {
		logrus.Fatalf("failed to init schema: %s", err.Error())
	}

	// --- Router ---
	mux := http.NewServeMux()
	router.RegisterRoutes(mux, ctrl)

	// --- Server ---
	port := os.Getenv("TIMETABLE_PORT")
	if port == "" {
		port = "8081"
	}
	addr := ":" + port

	log.Printf("Timetable API listening on http://localhost%s", addr)

	// Important: on wrap avec mux (sinon rien ne marche)
	if err := http.ListenAndServe(addr, mux); err != nil {
		logrus.Fatalf("server error: %s", err.Error())
	}
}
