package vii

import (
	"fmt"
	"time"
	"ztaylor.me/js"
)

type AccountCard struct {
	Username string
	CardId   int
	Register time.Time
	Notes    string
}

type AccountsCards map[int][]*AccountCard

func (stack AccountsCards) Json() js.Object {
	j := js.Object{}
	for cardId, list := range stack {
		j[fmt.Sprintf("%d", cardId)] = len(list)
	}
	return j
}

var AccountCardService interface {
	Test(username string) AccountsCards
	Forget(username string)
	Get(username string) (AccountsCards, error)
	Load(username string) (AccountsCards, error)
	Insert(username string) error
	InsertCard(card *AccountCard) error
	Delete(username string) error
	DeleteAndInsert(username string) error
}
