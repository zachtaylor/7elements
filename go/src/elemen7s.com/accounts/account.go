package accounts

import "time"

type Account struct {
	Username  string
	Email     string
	Password  string
	Language  string
	Coins     int
	Skill     int
	Packs     int
	Register  time.Time
	LastLogin time.Time
	SessionId uint
}
