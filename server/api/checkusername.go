package api

import (
	"errors"
	"strings"
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
	ErrUsernameSymbols  = errors.New("username must contain only letters and numbers")
	ErrUsernameBanned   = errors.New("username contains banned word")
)

const UsernameCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
	} else if symbols := strings.Trim(username, UsernameCharset); len(symbols) > 0 {
		return ErrUsernameSymbols
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
