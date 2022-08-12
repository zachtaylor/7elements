package game

import (
	"strconv"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/power"
)

// TokenContext is an in-play game object
type TokenContext struct {
	Card   string
	Name   string
	Text   string
	Image  string
	Awake  bool
	Body   *card.Body
	Powers power.Set
}

func DefaultTokenImage(cardno int) string {
	return strconv.FormatInt(int64(cardno), 10)
}

func NewTokenContext(card *Card) TokenContext {
	return TokenContext{
		Card:   card.ID(),
		Name:   card.T.Name,
		Text:   card.T.Text,
		Image:  DefaultTokenImage(card.T.ID),
		Body:   card.T.Body.Copy(),
		Powers: card.T.Powers.Copy(),
	}
}

func (t *TokenContext) String() string {
	if t == nil {
		return "nil"
	}
	return "Token{" + t.Name + " " + t.Image + " " + t.Body.String() + "}"
}

// Data returns a representation of a game token as map[string]any
func (t *TokenContext) Data() map[string]any {
	if t == nil {
		return nil
	}
	return map[string]any{
		"card":   t.Card,
		"name":   t.Name,
		"image":  t.Image,
		"awake":  t.Awake,
		"body":   t.Body.JSON(),
		"powers": t.Powers.JSON(),
	}
}
