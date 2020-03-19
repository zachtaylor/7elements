package account

import "ztaylor.me/cast"

// Card is a record of an Account owning a Card
type Card struct {
	Username string
	ProtoID  int
	Register cast.Time
	Notes    string
}

func (card *Card) String() string {
	return "{" + card.Username + ":" + cast.StringI(card.ProtoID) + "}"
}

// Cards is a map of Card Prototype ID to []*Card
type Cards map[int][]*Card

// JSON returns a representation of this set of Cards as type cast.JSON
func (stack Cards) JSON() cast.JSON {
	j := cast.JSON{}
	for cardId, list := range stack {
		j[cast.StringI(cardId)] = len(list)
	}
	return j
}

// CardService provides records of an Account owning Cards
type CardService interface {
	Test(username string) Cards
	Forget(username string)
	Find(username string) (Cards, error)
	Get(username string) (Cards, error)
	Insert(username string) error
	InsertCard(card *Card) error
	Delete(username string) error
	DeleteAndInsert(username string) error
}
