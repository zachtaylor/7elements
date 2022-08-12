package game

import "github.com/zachtaylor/7elements/card"

type Entry struct {
	Writer
	cardCount card.Count
}

func NewEntry(w Writer, cardCount card.Count) Entry {
	return Entry{
		Writer:    w,
		cardCount: cardCount,
	}
}
