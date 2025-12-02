package helpers

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/sirupsen/logrus"
)

func OpenDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "file:middleware.db?cache=shared&_foreign_keys=1")
    if err != nil {
        return nil, err
    }
    db.SetMaxOpenConns(1)
    return db, nil
}

func CloseDB(db *sql.DB) {
    if db == nil {
        return
    }
    err := db.Close()
    if err != nil {
        logrus.Errorf("error closing db : %s", err.Error())
    }
}