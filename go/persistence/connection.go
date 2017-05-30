package persistence

import (
	"7elements.ztaylor.me/log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var connection *sql.DB

func SetConnection(path string) {
	log.Add("Path", path)

	var err error
	connection, err = sql.Open("sqlite3", path)

	if err != nil {
		log.Add("Error", err).Error("persistence: connection")
	} else {
		log.Debug("persistence: connection opened")
	}
}
