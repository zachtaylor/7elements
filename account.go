package vii

import "time"

type Account struct {
	Username  string
	Email     string
	Password  string
	Coins     int
	Skill     int
	Register  time.Time
	LastLogin time.Time
	SessionID string
}

func NewAccount() *Account {
	return &Account{}
}

var AccountService interface {
	Test(string) *Account
	Cache(*Account)
	Forget(string)
	Get(string) (*Account, error)
	Load(string) (*Account, error)
	Insert(*Account) error
	UpdateCoins(*Account) error
	UpdateLogin(*Account) error
	UpdatePassword(*Account) error
	Delete(string) error
}
