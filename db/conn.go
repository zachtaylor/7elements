package db

import (
	"ztaylor.me/db"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

// DB_TABLE is name of env var
const DB_TABLE = "DB_TABLE"

var conn *db.DB

func Patch() (int, error) {
	if conn == nil {
		tableName := env.Get(DB_TABLE)
		var err error
		conn, err = db.Open(tableName)
		if err != nil {
			log.Add(DB_TABLE, tableName).Add(db.DB_HOST, env.Get(db.DB_HOST)).Error(err)
		} else {
			log.Add(DB_TABLE, tableName).Add(db.DB_HOST, env.Get(db.DB_HOST)).Debug("loaded db from env")
		}
	}
	return db.Patch(conn)
}
