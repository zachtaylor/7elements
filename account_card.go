package vii

import (
	"time"

	"ztaylor.me/cast"
)

type AccountCard struct {
	Username string
	CardId   int
	Register time.Time
	Notes    string
}

func (card *AccountCard) String() string {
	return "{" + card.Username + ":" + cast.StringI(card.CardId) + "}"
}

type AccountCards map[int][]*AccountCard

func (stack AccountCards) JSON() cast.JSON {
	j := cast.JSON{}
	for cardId, list := range stack {
		j[cast.StringI(cardId)] = len(list)
	}
	return j
}

type AccountCardService interface {
	Test(username string) AccountCards
	Forget(username string)
	Find(username string) (AccountCards, error)
	Get(username string) (AccountCards, error)
	Insert(username string) error
	InsertCard(card *AccountCard) error
	Delete(username string) error
	DeleteAndInsert(username string) error
}
