package api

import "ztaylor.me/charset"

// CheckUsername determines if a username is valid
func CheckUsername(username string) (_ bool) {
	if len(username) < 4 {
		return
	}
	return charset.In(username, charset.AlphaNumericCapital)
}
