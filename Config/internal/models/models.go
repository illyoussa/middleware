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