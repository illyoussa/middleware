package models

// ErrorNotFound représente une erreur métier "ressource inexistante"
type ErrorNotFound struct {
	Message string `json:"message"`
}

func (e *ErrorNotFound) Error() string {
	return e.Message
}

// ErrorUnprocessableEntity représente une erreur métier "données invalides"
type ErrorUnprocessableEntity struct {
	Message string `json:"message"`
}

func (e *ErrorUnprocessableEntity) Error() string {
	return e.Message
}

// (OPTIONNEL) ErrorConflict pour les cas de doublons / conflits métier
type ErrorConflict struct {
	Message string `json:"message"`
}

func (e *ErrorConflict) Error() string {
	return e.Message
}
