package models

// Agenda représente un agenda côté UCA (ID utile pour récupérer l'ical).
type Agenda struct {
    ID      string `json:"id"`      // identifiant UCA (ex: "13295")
    Name    string `json:"name"`    // nom lisible optionnel
    IcalURL string `json:"icalUrl"` // url complète si utile
}

// Alert représente une règle de notification liée à un agenda.
type Alert struct {
    ID        int64  `json:"id"`
    Recipient string `json:"recipient"` // adresse email ou identifiant
    AgendaID  string `json:"agendaId"`  // lien logique vers Agenda.ID
    Condition string `json:"condition"` // ex: "always", "on_change", "on_room_change"
    Method    string `json:"method"`    // ex: "email", "webhook"
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