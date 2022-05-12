package db

import (
	"errors"
	"strconv"

	"taylz.io/db"
	"taylz.io/db/mysql"
	"taylz.io/db/patch"
)

// OpenEnv connects using mysql.Open, and initializes the connection by
// verifying the patch table value matches ex_patch
func OpenEnv(env map[string]string, ex_patch int) (*db.DB, error) {
	db, err := mysql.Open(db.DSN(
		env["DB_USER"],
		env["DB_PASSWORD"],
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_NAME"],
	))
	if err != nil {
		return nil, err
	} else if patch, err := patch.Get(db); err != nil {
		return nil, err
	} else if patch != ex_patch {
		return nil, errors.New("Patch mismatch: " + strconv.FormatInt(int64(patch), 10))
	}
	return db, nil
}
