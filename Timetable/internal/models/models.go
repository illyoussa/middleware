package models

// Event structure
type Event struct {
	ResourceIds []int `json:"resourceIds"`

	Id string `json:"id"`

	AgendaID     string `json:"agendaId"`
	DTStamp      string `json:"dtstamp"`
	DTStart      string `json:"dtstart"`
	DTEnd        string `json:"dtend"`
	Summary      string `json:"summary"`
	Location     string `json:"location"`
	Description  string `json:"description"`
	Created      string `json:"created"`
	LastModified string `json:"lastModified"`
	Sequence     int64  `json:"sequence"`
}
