package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// OpenDB ouvre la base SQLite de l'API Timetable
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:timetable.db")
	if err != nil {
		return nil, err
	}

	// SQLite ne supporte pas bien plusieurs connexions ouvertes
	db.SetMaxOpenConns(1)

	return db, nil
}

// CloseDB ferme proprement la connexion à la base
func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		logrus.Errorf("error closing db: %s", err.Error())
	}
}
