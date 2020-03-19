package vii

import (
	"time"

	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
)

type AccountDeck struct {
	ID       int
	Name     string
	Username string
	Register time.Time
	Cards    map[int]int
	Wins     int
	CoverID  int
}

func NewAccountDeck() *AccountDeck {
	return &AccountDeck{
		Cards: make(map[int]int),
	}
}

func NewAccountDeckWith(proto *deck.Prototype, username string) *AccountDeck {
	ad := NewAccountDeck()
	ad.ID = -proto.ID
	ad.Name = proto.Name
	ad.Username = username
	ad.Register = time.Now()
	ad.Cards = make(map[int]int)
	for k, v := range proto.Cards {
		ad.Cards[k] = v
	}
	return ad
}

func (deck *AccountDeck) Count() int {
	total := 0
	for _, count := range deck.Cards {
		total += count
	}
	return total
}

func (deck *AccountDeck) JSON() cast.JSON {
	cardsJSON := cast.JSON{}
	for cardid, count := range deck.Cards {
		cardsJSON[cast.StringI(cardid)] = count
	}
	return cast.JSON{
		"id":       deck.ID,
		"name":     deck.Name,
		"username": deck.Username,
		"cards":    cardsJSON,
		"wins":     deck.Wins,
		"cover":    "/img/card/" + cast.StringI(deck.CoverID) + ".jpg",
	}
}

type AccountDecks []*AccountDeck

func (decks AccountDecks) JSON() cast.JSON {
	data := cast.JSON{}
	for _, deck := range decks {
		data[cast.StringI(deck.ID)] = deck.JSON()
	}
	return data
}

type AccountDeckService interface {
	Find(username string) (AccountDecks, error)
	Forget(username string)
	Update(deck *AccountDeck) error
	UpdateName(username string, id int, name string) error
}
