package main

import (
	"middleware/example/internal/controllers/events"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/events", func(r chi.Router) {
		r.Get("/", events.GetEvents)
		r.Post("/", events.CreateEvent)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(events.Context)
			r.Get("/", events.GetEvent)
			r.Put("/", events.UpdateEvent)
			r.Delete("/", events.DeleteEvent)
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

	// Drop + recreate
	if _, err := db.Exec("DROP TABLE IF EXISTS events;"); err != nil {
		logrus.Fatalln("Could not drop table ! Error was : " + err.Error())
	}

	schemes := []string{
		`-- Events : stockage simplifié des VEVENTs extraits d'un fichier iCal
		CREATE TABLE IF NOT EXISTS events (
			uid TEXT NOT NULL,
			dtstamp TEXT NOT NULL DEFAULT '1111',
			dtstart TEXT NOT NULL DEFAULT '1111',
			dtend TEXT NOT NULL DEFAULT '1111',
			summary TEXT NOT NULL DEFAULT '1111',
			location TEXT NOT NULL DEFAULT '1111',
			description TEXT NOT NULL DEFAULT '1111',
			created TEXT NOT NULL DEFAULT '1111',
			last_modified TEXT NOT NULL DEFAULT '1111',
			sequence INTEGER NOT NULL DEFAULT 0,
			agenda_id VARCHAR(255) NOT NULL,
			PRIMARY KEY (agenda_id, uid)
		);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}

	logrus.Info("[INFO] Seeding database with test data...")

	agendaTestId1 := "11111111-1111-1111-1111-111111110001"
	agendaTestId2 := "11111111-1111-1111-1111-111111110002"

	// Nettoyage
	if _, err := db.Exec("DELETE FROM events"); err != nil {
		logrus.Fatalln("Could not clear events table ! Error was : " + err.Error())
	}

	// On remplit TOUTES les colonnes pour éviter les NULL et rester cohérent avec ton model Go (string/int64)
	seedQueries := []string{
		`INSERT INTO events (
			uid, dtstamp, dtstart, dtend, summary, location, description, created, last_modified, sequence, agenda_id
		) VALUES
			('evt1', '1111', '2025-11-20T08:00:00Z', '2025-11-20T10:00:00Z', 'Cours de Go', 'Amphi 1', '1111', '1111', '1111', 0, '` + agendaTestId1 + `'),
			('evt2', '1111', '2025-11-21T14:00:00Z', '2025-11-21T15:00:00Z', 'Réunion Projet', 'Salle B204', '1111', '1111', '1111', 0, '` + agendaTestId1 + `'),
			('evt3', '1111', '2025-11-22T12:30:00Z', '2025-11-22T13:30:00Z', 'Rdv Dentiste', '12 rue du Test', '1111', '1111', '1111', 0, '` + agendaTestId2 + `');`,
	}

	for _, query := range seedQueries {
		if _, err := db.Exec(query); err != nil {
			logrus.Fatalln("Could not seed data ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)
}
