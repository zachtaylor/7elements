package api

import (
	"errors"
	"strings"

	"taylz.io/keygen/charset"
)

var userAllow = []string{
	"zach",
}

var userBan = []string{
	"fuck",
}

var (
	ErrUsernameTooShort = errors.New("username too short")
	ErrUsernameTooLong  = errors.New("username too long")
	ErrUsernameBanned   = errors.New("username contains banned word")
)

func CheckUsername(username string) error {
	for _, allow := range userAllow {
		if username == allow {
			return nil
		}
	}

	if len(username) < 7 {
		return ErrUsernameTooShort
	} else if len(username) > 21 {
		return ErrUsernameTooLong
	} else if symbols := strings.Trim(username, charset.AlphaCapitalNumeric); len(symbols) > 0 {
		return errors.New("username symbols not allowed: " + symbols)
	} else if ban := checkUsernameContainsBan(username); len(ban) > 1 {
		return ErrUsernameBanned
	}

	return nil
}

func checkUsernameContainsBan(username string) string {
	for _, ban := range userBan {
		if strings.Contains(username, ban) {
			return ban
		}
	}
	return ""
}
