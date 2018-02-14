package security

import (
	"crypto/md5"
	"elemen7s.com/options"
)

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password + options.String("password-salt")))
	password = string(hash[:])
	return password
}
