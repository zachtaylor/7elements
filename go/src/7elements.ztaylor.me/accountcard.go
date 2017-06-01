package SE

import (
	"time"
)

type AccountCard struct {
	Username string
	Card     uint
	Register time.Time
	Notes    string
}

// persistence headers
var AccountsCards = struct {
	Cache  map[string]CardCollection
	Get    func(string) (CardCollection, error)
	Insert func(string) error
	Delete func(string) error
}{make(map[string]CardCollection), nil, nil, nil}
