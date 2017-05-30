package persistence

import (
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/options"
	_ "github.com/mattes/migrate/driver/sqlite3"
	"github.com/mattes/migrate/migrate"
)

func Patch() {
	errors, ok := migrate.UpSync("sqlite3://"+options.String("db-path"), options.String("patch-path"))

	if !ok {
		for _, err := range errors {
			log.Add("Path", options.String("db-path")).Error(err.Error())
		}
	}
}

func CheckPatch() uint64 {
	patch, _ := migrate.Version("sqlite3://"+options.String("db-path"), options.String("patch-path"))

	return patch
}
