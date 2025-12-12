package main

import (
	"middleware/example/internal/controllers/agendas"
	"middleware/example/internal/controllers/alerts"
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

	r.Route("/alerts", func(r chi.Router) {
		r.Get("/", alerts.GetAlerts)
		r.Post("/", alerts.CreateAlert)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(alerts.Context)
			r.Get("/", alerts.GetAlert)
			r.Put("/", alerts.UpdateAlert)
			r.Delete("/", alerts.DeleteAlert)
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

	db.Exec("DROP TABLE IF EXISTS alerts;")
	db.Exec("DROP TABLE IF EXISTS agendas;")

	schemes := []string{
		`-- Agendas : identifiants UCA pour récupération iCal
        CREATE TABLE IF NOT EXISTS agendas (
            id VARCHAR(255) PRIMARY KEY NOT NULL,
            name TEXT,
            ical_url TEXT
        );`,

		`-- Alerts : règles de notification liées à un agenda
        CREATE TABLE IF NOT EXISTS alerts (
            id VARCHAR(255) PRIMARY KEY NOT NULL,
            recipient TEXT NOT NULL,
            agenda_id VARCHAR(255) NOT NULL,
            condition TEXT DEFAULT 'always',
            method TEXT DEFAULT 'email'
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
	agendaTestId3 := "11111111-1111-1111-1111-111111110003"

	alertTestId1 := "22222222-2222-2222-2222-222222220001"
	alertTestId2 := "22222222-2222-2222-2222-222222220002"
	alertTestId3 := "22222222-2222-2222-2222-222222220003"

	// Nettoyage
	if _, err := db.Exec("DELETE FROM alerts"); err != nil {
		logrus.Fatalln("Could not clear alerts table ! Error was : " + err.Error())
	}
	if _, err := db.Exec("DELETE FROM agendas"); err != nil {
		logrus.Fatalln("Could not clear agendas table ! Error was : " + err.Error())
	}

	seedQueries := []string{
		// Agendas
		`INSERT INTO agendas (id, name, ical_url) VALUES
            ('` + agendaTestId1 + `', 'Emploi du temps - L3 Info', 'https://example.com/ical/l3_info.ics'),
            ('` + agendaTestId2 + `', 'Agenda Personnel', 'https://example.com/ical/perso.ics'),
			('` + agendaTestId3 + `', 'Agenda Vacances', 'https://example.com/ical/vacances.ics');`,

		// Alerts
		`INSERT INTO alerts (id, recipient, agenda_id, condition, method) VALUES
            ('` + alertTestId1 + `', 'test@example.com', '` + agendaTestId1 + `', 'on_change', 'email'),
            ('` + alertTestId2 + `', 'admin@example.com', '` + agendaTestId1 + `', 'always', 'email'),
			('` + alertTestId3 + `', 'employe@example.com, '` + agendaTestId2 + `', 'on_change', 'sms');`,
	}

	for _, query := range seedQueries {
		if _, err := db.Exec(query); err != nil {
			logrus.Fatalln("Could not seed data ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)
}
