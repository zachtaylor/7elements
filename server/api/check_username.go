package api

import (
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
)

// CheckUsername determines if a username is valid
func CheckUsername(username string) (_ bool) {
	if len(username) < 4 {
		return
	}
	return cast.InCharset(username, charset.AlphaCapitalNumeric)
}
