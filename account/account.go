package account

import (
	"time"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
)

// T is an Account
type T struct {
	Username  string
	Email     string
	Password  string
	Coins     int
	Skill     int
	Register  time.Time
	LastLogin time.Time
	Verify    int
	// runtime
	GameID    string
	PlayerID  string
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

func Make(username, email, password, sessionid string) *T {
	time := time.Now()
	return &T{
		Username:  username,
		Email:     email,
		Password:  password,
		Register:  time,
		LastLogin: time,
		SessionID: sessionid,
		Cards:     make(card.Count),
		Decks:     make(deck.Prototypes),
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

func (a *T) Data() map[string]any {
	if a == nil {
		return nil
	}
	return map[string]any{
		"username": a.Username,
		"email":    a.Email,
		"session":  a.SessionID,
		"coins":    a.Coins,
		"cards":    a.Cards.Data(),
		"decks":    a.Decks.Data(),
		"game":     a.GameID,
	}
}
