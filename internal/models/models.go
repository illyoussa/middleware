package models

import (
	"github.com/gofrs/uuid"
)

// Agenda représente un agenda côté UCA (ID utile pour récupérer l'ical).
type Agenda struct {
	Id      *uuid.UUID `json:"id"`
	Name    string     `json:"name"`
	IcalURL string     `json:"ical_url"`
}

// Alert représente une règle de notification liée à un agenda.
type Alert struct {
	Id        *uuid.UUID `json:"id"`
	Recipient string     `json:"recipient"`
	AgendaID  string     `json:"agendaId"`
	Condition string     `json:"condition"`
	Method    string     `json:"method"`
}

// Event représente un VEVENT extrait d'un fichier iCal.
type Event struct {
	UID          string `json:"uid"`
	DTStamp      string `json:"dtstamp"`
	DTStart      string `json:"dtstart"`
	DTEnd        string `json:"dtend"`
	Summary      string `json:"summary"`
	Location     string `json:"location"`
	Description  string `json:"description"`
	Created      string `json:"created"`
	LastModified string `json:"lastModified"`
	Sequence     int64  `json:"sequence"`
	AgendaID     string `json:"agendaId"`
}
