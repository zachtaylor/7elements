package SE

import (
	"time"
)

type Account struct {
	Username  string
	Email     string
	Password  string
	Language  string
	Register  time.Time
	LastLogin time.Time
	SessionId uint
}

// persistence headers
var Accounts = struct {
	Cache  map[string]*Account
	Get    func(string) (*Account, error)
	Insert func(string) error
	Delete func(string) error
}{make(map[string]*Account), nil, nil, nil}
