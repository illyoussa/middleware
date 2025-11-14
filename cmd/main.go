package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/helpers"
    _ "middleware/example/internal/models"
)

func main() {
    r := chi.NewRouter()

    // health
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ok"))
    })

    // API placeholder : tu vas ajouter tes handlers sous /api/...
    // Exemple minimal : r.Mount("/api", apiRouter)

    // Serve static frontend (prébuild dans ./build)
    fs := http.FileServer(http.Dir("./build"))
    // Serve root index and other app assets
    r.Handle("/*", fs)
    r.Handle("/", fs)

    logrus.Info("[INFO] Web server started. Now listening on *:8080")
    logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
    db, err := helpers.OpenDB()
    if err != nil {
        logrus.Fatalf("error while opening database : %s", err.Error())
    }

    schemes := []string{
        `-- Agendas : identifiants UCA pour récupération iCal
        CREATE TABLE IF NOT EXISTS agendas (
            id VARCHAR(255) PRIMARY KEY NOT NULL,
            name TEXT,
            ical_url TEXT
        );`,
        `-- Alerts : règles de notification liées à un agenda
        CREATE TABLE IF NOT EXISTS alerts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            recipient TEXT NOT NULL,
            agenda_id VARCHAR(255) NOT NULL,
            condition TEXT DEFAULT 'always',
            method TEXT DEFAULT 'email'
        );`,
        `-- Events : stockage simplifié des VEVENTs extraits d'un fichier iCal
        CREATE TABLE IF NOT EXISTS events (
            uid TEXT PRIMARY KEY NOT NULL,
            dtstamp TEXT,
            dtstart TEXT,
            dtend TEXT,
            summary TEXT,
            location TEXT,
            description TEXT,
            created TEXT,
            last_modified TEXT,
            sequence INTEGER,
            agenda_id VARCHAR(255)
        );`,
    }

    for _, scheme := range schemes {
        if _, err := db.Exec(scheme); err != nil {
            logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
        }
    }
    helpers.CloseDB(db)
}