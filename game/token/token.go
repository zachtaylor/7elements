package token

import (
	"strconv"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/power"
)

// T is an in-play game object
type T struct {
	ID      string
	Card    *card.T
	User    string
	Image   string
	IsAwake bool
	Body    *card.Body
	Powers  power.Set
}

func New(c *card.T, user string) *T {
	return &T{
		Card:   c,
		User:   user,
		Image:  strconv.FormatInt(int64(c.Proto.ID), 10),
		Body:   c.Proto.Body.Copy(),
		Powers: c.Proto.Powers.Copy(),
	}
}

func (t *T) String() string {
	if t == nil {
		return "nil"
	}
	return "T#" + t.ID + ":" + t.Card.Proto.Name
}

// Data returns a representation of a game token
func (t *T) Data() map[string]interface{} {
	if t == nil {
		return nil
	}
	return map[string]interface{}{
		"id":     t.ID,
		"cardid": t.Card.Proto.ID,
		"image":  t.Image,
		"awake":  t.IsAwake,
		"body":   t.Body.Data(),
		"name":   t.Card.Proto.Name,
		"text":   t.Card.Proto.Text,
		"owner":  t.Card.User,
		"user":   t.User,
		"type":   t.Card.Proto.Type.String(),
		"powers": t.Powers.Data(),
	}
}
