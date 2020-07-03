package account

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
)

// T is an Account
type T struct {
	Username  string
	Email     string
	Password  string
	Coins     int
	Skill     int
	Register  cast.Time
	LastLogin cast.Time
	SessionID string
	Cards     card.Count
	Decks     deck.Prototypes
}

// New returns an empty Account
func New() *T {
	return &T{
		Cards: make(card.Count),
		Decks: make(deck.Prototypes),
	}
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

func (a *T) JSON() cast.JSON {
	if a == nil {
		return nil
	}
	return cast.JSON{
		"username": a.Username,
		"email":    a.Email,
		"session":  a.SessionID,
		"coins":    a.Coins,
		"cards":    a.Cards.JSON(),
		"decks":    a.Decks.JSON(),
	}
}
