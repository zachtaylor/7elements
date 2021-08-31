package api

import (
	"errors"
	"strings"

	"taylz.io/keygen/charset"
)

const charsetEmail = charset.AlphaCapitalNumeric + `-+@.`

var ErrIllegalEmail = errors.New("illegal email address")

func CheckEmail(email string) (err error) {
	if illegal := strings.Trim(email, charsetEmail); len(illegal) > 0 {
		err = ErrIllegalEmail
	}
	if strings.ContainsAny(email, "'()\"") {
		err = ErrIllegalEmail
	}
	return
}
