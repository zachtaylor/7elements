package security

import (
	"7elements.ztaylor.me/options"
	"crypto/md5"
)

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password + options.String("password-salt")))
	password = string(hash[:])
	return password
}
