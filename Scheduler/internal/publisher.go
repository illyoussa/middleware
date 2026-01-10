package events

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func PublishEvents(events []Event, subject string) error {
	// Connexion au serveur NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	defer nc.Close()

	// Publication des événements
	for _, event := range events {
		data, err := json.Marshal(event)
		if err != nil {
			log.Printf("Erreur lors de la conversion de l'événement en JSON : %s", err)
			continue
		}

		if err := nc.Publish(subject, data); err != nil {
			log.Printf("Erreur lors de la publication de l'événement : %s", err)
		}
	}

	return nil
}
