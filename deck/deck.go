package deck

import (
	"github.com/cznic/mathutil"
)

// T is a Deck
// type T struct {
// 	Proto *Prototype
// 	User  string
// 	Cards []*card.T
// }

// var ErrInvalidCardID = errors.New("invalid card id")

// // BuildNew returns a new Deck
// func BuildNew(username string, proto *Prototype, cards card.Prototypes) (*T, error) {
// 	buf := make([]*card.T, proto.Count())
// 	i := 0
// 	for cardid, copy := range proto.Cards {
// 		if proto := cards[cardid]; proto == nil {
// 			return nil, ErrInvalidCardID
// 		} else {
// 			for ii := 0; ii < copy; ii++ {
// 				buf[i] = card.New(proto, username)
// 				i++
// 			}
// 		}
// 	}
// 	return &T{
// 		Proto: proto,
// 		User:  username,
// 		Cards: buf,
// 	}, nil
// }

// func (deck *T) Count() int {
// 	return len(deck.Cards)
// }

// // Draw removes the top card of the Deck and returns it
// func (deck *T) Draw() *card.T {
// 	if len(deck.Cards) < 1 {
// 		return nil
// 	}
// 	card := deck.Cards[0]
// 	deck.Cards = deck.Cards[1:]
// 	return card
// }

// // Prepend places a Card on top of the Deck
// func (deck *T) Prepend(c *card.T) {
// 	deck.Cards = append([]*card.T{c}, deck.Cards...)
// }

// // Append places a Card on bottom of the Deck
// func (deck *T) Append(c *card.T) {
// 	deck.Cards = append(deck.Cards, c)
// }

// // Shuffle randomizes the order of Cards in the Deck
// func (deck *T) Shuffle() {
// 	shuffleRandom, _ := mathutil.NewFC32(0, len(deck.Cards)-1, true)
// 	cp := make([]*card.T, len(deck.Cards))
// 	for i := 0; i < len(deck.Cards); i++ {
// 		cp[shuffleRandom.Next()] = deck.Cards[i]
// 	}
// 	deck.Cards = cp
// }

func Shuffle(deck []string) []string {
	len := len(deck) - 1
	order, _ := mathutil.NewFC32(0, len, true)
	copy := make([]string, len+1)
	for i := 0; i < len; i++ {
		copy[order.Next()] = deck[i]
	}
	return copy
}

// func (deck *T) JSON() map[string]any {
// 	cardsJSON := map[string]any{}
// 	for cardid, count := range deck.Cards {
// 		cardsJSON[strconv.FormatInt(int64(cardid), 10)] = count
// 	}
// 	return map[string]any{
// 		"id":    deck.Proto.ID,
// 		"name":  deck.Proto.Name,
// 		"cards": cardsJSON,
// 		"wins":  deck.Proto.Wins,
// 		"cover": deck.Proto.Cover,
// 	}
// }
