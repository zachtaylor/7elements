package account

import "time"

// T is an Account
type T struct {
	Username  string
	Email     string
	Password  string
	Coins     int
	Skill     int
	Register  time.Time
	LastLogin time.Time
	SessionID string
}

func (a *T) String() string {
	if a == nil {
		return ""
	}
	s := a.Username
	if a.Email != "" {
		s += "(" + a.Email + ")"
	}
	return s
}
