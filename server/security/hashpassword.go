package security

import (
	"crypto/md5"
	"ztaylor.me/env"
)

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password + env.Default("DB_PWSALT", "")))
	password = string(hash[:])
	return password
}
