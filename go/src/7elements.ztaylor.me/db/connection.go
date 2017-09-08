package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"ztaylor.me/log"
)

var Connection *sql.DB

func Open(path string) {
	log := log.Add("Path", path)

	var err error
	Connection, err = sql.Open("sqlite3", path)

	if err != nil {
		log.Add("Error", err).Error("db: connection failed")
	} else {
		log.Debug("db opened")
	}
}
