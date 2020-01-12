package vii

import (
	"time"
)

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

func (a *Account) String() string {
	if a == nil {
		return ""
	}
	s := a.Username
	if a.Email != "" {
		s += "(" + a.Email + ")"
	}
	return s
}

func NewAccount() *Account {
	return &Account{}
}

// AccountService provides Accounts
type AccountService interface {
	// Test returns an account from cache only
	Test(string) *Account
	// Cache stores an account
	Cache(*Account)
	// Forget uncaches an account
	Forget(string)
	// Find uses Test/Get/Cache best effort to provide account
	Find(string) (*Account, error)
	// Get loads an account from back end
	Get(string) (*Account, error)
	// GetCount returns a number of registered accounts from back end
	GetCount() (int, error)
	// Insert creates an account on back end
	Insert(*Account) error
	// UpdateCoins updates an accounts coin count on back end
	UpdateCoins(*Account) error
	// UpdateEmail updates an accounts email
	UpdateEmail(*Account) error
	// UpdateLogin updates an accounts login time on back end
	UpdateLogin(*Account) error
	// UpdatePassword updates an accounts password on back end
	UpdatePassword(*Account) error
	// Delete removes an account
	Delete(string) error
}
