package account

import (
	"time"

	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
)

type Deck struct {
	ID       int
	Name     string
	Username string
	Register time.Time
	Cards    map[int]int
	Wins     int
	CoverID  int
}

func NewDeck() *Deck {
	return &Deck{
		Cards: make(map[int]int),
	}
}

func NewDeckWith(proto *deck.Prototype, username string) *Deck {
	ad := NewDeck()
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

func (deck *Deck) Count() int {
	total := 0
	for _, count := range deck.Cards {
		total += count
	}
	return total
}

func (deck *Deck) JSON() cast.JSON {
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

type Decks []*Deck

func (decks Decks) JSON() cast.JSON {
	data := cast.JSON{}
	for _, deck := range decks {
		data[cast.StringI(deck.ID)] = deck.JSON()
	}
	return data
}

type DeckService interface {
	Find(username string) (Decks, error)
	Forget(username string)
	Update(deck *Deck) error
	UpdateName(username string, id int, name string) error
}
