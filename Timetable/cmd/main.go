package main

import (
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/agendas", func(r chi.Router) {
		r.Get("/", agendas.GetAgendas)
		r.Post("/", agendas.CreateAgenda)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(agendas.Context)
			r.Get("/", agendas.GetAgenda)
			r.Put("/", agendas.UpdateAgenda)
			r.Delete("/", agendas.DeleteAgenda)
		})
	})


	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	db.Exec("DROP TABLE IF EXISTS events;")

	schemes := []string{
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

	logrus.Info("[INFO] Seeding database with test data...")

	// Nettoyage
	if _, err := db.Exec("DELETE FROM events"); err != nil {
		logrus.Fatalln("Could not clear events table ! Error was : " + err.Error())
	}

	seedQueries := []string{

		// Events
		`INSERT INTO events (uid, dtstart, dtend, summary, location, agenda_id) VALUES
            ('evt1', '2025-11-20T08:00:00Z', '2025-11-20T10:00:00Z', 'Cours de Go', 'Amphi 1', '` + agendaTestId1 + `'),
            ('evt2', '2025-11-21T14:00:00Z', '2025-11-21T15:00:00Z', 'Réunion Projet', 'Salle B204', '` + agendaTestId1 + `'),
            ('evt3', '2025-11-22T12:30:00Z', '2025-11-22T13:30:00Z', 'Rdv Dentiste', '12 rue du Test', '` + agendaTestId2 + `');`,
	}

	for _, query := range seedQueries {
		if _, err := db.Exec(query); err != nil {
			logrus.Fatalln("Could not seed data ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)
}
