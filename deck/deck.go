package deck

import (
	"github.com/cznic/mathutil"
	"github.com/zachtaylor/7elements/card"
	"taylz.io/log"
)

// T is a Deck
type T struct {
	Proto *Prototype
	User  string
	Cards []*card.T
}

// New returns a new Deck
func New(log *log.T, cards card.Prototypes, proto *Prototype, user string) *T {
	buf := make([]*card.T, proto.Count())
	i := 0
	for cardid, copy := range proto.Cards {
		if proto, err := cards[cardid]; proto == nil {
			log.New().Add("CardID", cardid).Error(err, "invalid cardid")
		} else {
			for ii := 0; ii < copy; ii++ {
				buf[i] = card.New(proto, user)
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

// func (deck *T) JSON() websocket.MsgData {
// 	cardsJSON := websocket.MsgData{}
// 	for cardid, count := range deck.Cards {
// 		cardsJSON[strconv.FormatInt(int64(cardid), 10)] = count
// 	}
// 	return websocket.MsgData{
// 		"id":    deck.Proto.ID,
// 		"name":  deck.Proto.Name,
// 		"cards": cardsJSON,
// 		"wins":  deck.Proto.Wins,
// 		"cover": deck.Proto.Cover,
// 	}
// }
