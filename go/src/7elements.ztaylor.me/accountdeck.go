package SE

import (
	"time"
)

type AccountDeck struct {
	Name     string
	Id       int
	Register time.Time
	Cards    map[int]int
	Wins     int
}

func NewAccountDeck() *AccountDeck {
	return &AccountDeck{
		Cards: make(map[int]int),
	}
}

// persistence headers
var AccountsDecks = struct {
	Cache  map[string]map[int]*AccountDeck
	Get    func(string) (map[int]*AccountDeck, error)
	Insert func(string, int) error
	Delete func(string, int) error
}{make(map[string]map[int]*AccountDeck), nil, nil, nil}
