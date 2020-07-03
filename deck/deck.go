package deck

import (
	"github.com/cznic/mathutil"
	"github.com/zachtaylor/7elements/card"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

// T is a Deck
type T struct {
	Proto *Prototype
	User  string
	Cards []*card.T
}

// New returns a new Deck
func New(log log.Service, cards card.PrototypeService, proto *Prototype, user string) *T {
	buf := make([]*card.T, len(proto.Cards))
	i := 0
	for k, v := range proto.Cards {
		if proto, err := cards.Get(k); proto == nil {
			log.New().Add("CardID", k).Error(err, "invalid cardid")
		} else {
			for j := 0; j < v; j++ {
				buf[i] = card.New(proto)
				i++
			}
		}
	}
	return &T{
		Proto: proto,
		User:  user,
		Cards: buf,
	}
}

func (deck *T) Count() int {
	return len(deck.Cards)
}

// Draw removes the top card of the Deck and returns it
func (deck *T) Draw() *card.T {
	if len(deck.Cards) < 1 {
		return nil
	}
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

// Prepend places a Card on top of the Deck
func (deck *T) Prepend(c *card.T) {
	deck.Cards = append([]*card.T{c}, deck.Cards...)
}

// Append places a Card on bottom of the Deck
func (deck *T) Append(c *card.T) {
	deck.Cards = append(deck.Cards, c)
}

// Shuffle randomizes the order of Cards in the Deck
func (deck *T) Shuffle() {
	shuffleRandom, _ := mathutil.NewFC32(0, len(deck.Cards)-1, true)
	cp := make([]*card.T, len(deck.Cards))
	for i := 0; i < len(deck.Cards); i++ {
		cp[shuffleRandom.Next()] = deck.Cards[i]
	}
	deck.Cards = cp
}

func (deck *T) JSON() cast.JSON {
	cardsJSON := cast.JSON{}
	for cardid, count := range deck.Cards {
		cardsJSON[cast.StringI(cardid)] = count
	}
	return cast.JSON{
		"id":    deck.Proto.ID,
		"name":  deck.Proto.Name,
		"cards": cardsJSON,
		"wins":  deck.Proto.Wins,
		"cover": "/img/card/" + cast.StringI(deck.Proto.Cover) + ".jpg",
	}
}
