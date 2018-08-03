package db

import (
	"ztaylor.me/db"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

// DB_TABLE is name of env var
const DB_TABLE = "DB_TABLE"

var conn *db.DB

func init() {
	tableName := env.Get("DB_TABLE")
	var err error
	conn, err = db.Open(tableName)
	if err != nil {
		log.Error(err)
	}
}

func Patch() (int, error) {
	return db.Patch(conn)
}
