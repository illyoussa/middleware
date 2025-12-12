package models

import "time"

// Event représente un cours (évènement) dans l'emploi du temps.
type Event struct {
	// ID interne en base (auto-incrément).
	ID int64 `json:"id"`

	// AgendaID: identifiant d'agenda/groupe (vient de Config, sert à savoir à quel EDT appartient l'event).
	AgendaID string `json:"agenda_id"`

	// UID: identifiant unique stable venant de l'iCal (très utile pour détecter les changements).
	UID string `json:"uid"`

	// Title: nom du cours.
	Title string `json:"title"`

	// Description: détails éventuels.
	Description string `json:"description,omitempty"`

	// Location: salle / lieu.
	Location string `json:"location,omitempty"`

	// Start / End: horaires du cours (RFC3339 en JSON).
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`

	// LastModified: date de dernière modif vue côté iCal (utile pour détecter un changement).
	// Si tu ne l’as pas dans le flux iCal, tu peux laisser à zéro.
	LastModified time.Time `json:"last_modified,omitempty"`

	// UpdatedAt: date de dernière mise à jour en base.
	UpdatedAt time.Time `json:"updated_at"`

	// CreatedAt: date de création en base.
	CreatedAt time.Time `json:"created_at"`
}
