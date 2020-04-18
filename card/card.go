package card

import "ztaylor.me/cast"

// T is a Card
type T struct {
	ID       string
	Proto    *Prototype
	Username string
}

func New(c *Prototype) *T {
	return &T{
		Proto: c,
	}
}

func (c *T) String() string {
	if c == nil {
		return `<nil>`
	}
	return cast.StringN(
		`card.T{`,
		c.ID,
		`card:`, c.Proto.String(),
		`user:`, c.Username,
		`}`,
	)
}

// JSON returns a representation of a game card
func (c *T) JSON() cast.JSON {
	if c == nil {
		return nil
	}
	return cast.JSON{
		"id":       c.ID,
		"cardid":   c.Proto.ID,
		"name":     c.Proto.Name,
		"costs":    c.Proto.Costs.JSON(),
		"text":     c.Proto.Text,
		"username": c.Username,
		"image":    c.Proto.Image,
		"type":     c.Proto.Type.String(),
		"powers":   c.Proto.Powers.JSON(),
	}
}
