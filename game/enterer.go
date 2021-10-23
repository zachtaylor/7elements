package game

import (
	"github.com/zachtaylor/7elements/deck"
	"taylz.io/http/user"
)

type Enterer interface {
	Deck() *deck.Prototype
	Writer() user.Writer
}

type Entry struct {
	deck   *deck.Prototype
	writer user.Writer
}

func (*T) NewEntry(deck *deck.Prototype, user user.Writer) *Entry {
	return NewEntry(deck, user)
}

func NewEntry(deck *deck.Prototype, user user.Writer) *Entry {
	return &Entry{
		deck:   deck,
		writer: user,
	}
}

func (e Entry) Deck() *deck.Prototype { return e.deck }

func (e Entry) Writer() user.Writer { return e.writer }
