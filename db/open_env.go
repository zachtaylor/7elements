package db

import (
	"taylz.io/db"
	"taylz.io/db/mysql"
	"taylz.io/env"
)

// OpenEnv connects using mysql.Open db.OpenEnv and returns the conn and patch id
func OpenEnv(env env.Service) (*db.DB, error) {
	return mysql.Open(db.DSN(
		env["DB_USER"],
		env["DB_PASSWORD"],
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_NAME"],
	))
	// conn, err := db.OpenEnv(env)
	// if conn == nil {
	// 	log.StdOutService(log.LevelInfo).New().Warn("failed to open env")
	// 	return conn, -1, err
	// }
	// patch, err := Patch(conn)
	// return conn, patch, err
}
