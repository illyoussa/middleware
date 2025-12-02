package models

import (
	"github.com/gofrs/uuid"
)

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
