package db

import (
	"ztaylor.me/db"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

var conn *db.DB

// ConnPatch connects using db.OpenEnv and returns the patch id
func ConnPatch() (int, error) {
	if conn == nil {
		var err error
		conn, err = db.OpenEnv()
		if err != nil {
			log.Add(db.DB_TABLE, env.Get(db.DB_TABLE)).Add(db.DB_HOST, env.Get(db.DB_HOST)).Add("Error", err).Error("db: openenv failed")
		} else {
			log.Add(db.DB_TABLE, env.Get(db.DB_TABLE)).Add(db.DB_HOST, env.Get(db.DB_HOST)).Debug("db: openenv")
		}
	}
	return db.Patch(conn)
}
