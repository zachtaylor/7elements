package card

// T is a Card
type T struct {
	Proto *Prototype
	ID    string
	User  string
}

func New(proto *Prototype, user string) *T {
	return &T{
		Proto: proto,
		User:  user,
	}
}

func (c *T) String() string {
	if c == nil {
		return `<nil>`
	}
	return `card.T{` + c.ID + ` card:` + c.Proto.String() + ` user:` + c.User + `}`
}

// Data returns a representation of a game card
func (c *T) Data() map[string]interface{} {
	if c == nil {
		return nil
	}
	return map[string]interface{}{
		"id":     c.ID,
		"user":   c.User,
		"cardid": c.Proto.ID,
	}
}
