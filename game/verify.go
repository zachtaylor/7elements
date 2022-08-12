package game

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
)

var ErrDeckTooSmall = errors.New("deck size too small")
var ErrDeckTooBig = errors.New("deck size too big")
var ErrDeckCopy = errors.New("deck contains too many copies")

func VerifyRulesDeck(rules Rules, cards card.Count) error {
	var actLen, actCopy int
	for _, count := range cards {
		actLen += count
		if count > actCopy {
			actCopy = count
		}
	}
	if actLen < rules.DeckMin {
		return ErrDeckTooSmall
	} else if rules.DeckMax > 0 && actLen > rules.DeckMax {
		return ErrDeckTooBig
	} else if actCopy > rules.DeckCopy {
		return ErrDeckCopy
	}
	return nil
}
