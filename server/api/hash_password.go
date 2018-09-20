package api

import "crypto/md5"

// HashPassword performs a md5 hash using a given salt
func HashPassword(password string, salt string) string {
	hash := md5.Sum([]byte(password + salt))
	password = string(hash[:])
	return password
}
