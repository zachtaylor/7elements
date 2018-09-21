package db

import (
	"ztaylor.me/db"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

var Conn *db.DB

// OpenEnv connects using db.OpenEnv and returns the patch id
func OpenEnv() (int, error) {
	if Conn == nil {
		var err error
		Conn, err = db.OpenEnv()
		if err != nil {
			log.Add(db.DB_TABLE, env.Get(db.DB_TABLE)).Add(db.DB_HOST, env.Get(db.DB_HOST)).Add("Error", err).Error("db: openenv failed")
		} else {
			log.Add(db.DB_TABLE, env.Get(db.DB_TABLE)).Add(db.DB_HOST, env.Get(db.DB_HOST)).Debug("db: openenv")
		}
	}
	return Patch(Conn)
}
