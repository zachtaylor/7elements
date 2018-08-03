package vii

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
	UpdatePacks(*Account) error
	UpdateLogin(*Account) error
	UpdatePassword(*Account) error
	Delete(string) error
}
