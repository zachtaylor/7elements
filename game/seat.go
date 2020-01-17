package game

import (
	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/cast"
)

type Seat struct {
	Username string
	Life     int
	Deck     *Deck
	Karma    vii.Karma
	Hand     Cards
	Present  Tokens
	Past     Cards
	Color    string
	Receiver vii.JSONWriter
}

func (game *T) NewSeat() *Seat {
	return &Seat{
		Karma:   vii.Karma{},
		Present: Tokens{},
		Hand:    Cards{},
		Past:    Cards{},
	}
}

func (seat *Seat) DrawCard(count int) {
	for i := 0; i < count && len(seat.Deck.Cards) > 0; i++ {
		card := seat.Deck.Draw()
		seat.Hand[card.ID] = card
	}
}

func (seat *Seat) DiscardHand() {
	for _, card := range seat.Hand {
		seat.Past[card.ID] = card
	}
	seat.Hand = Cards{}
}

// WriteJSON sends data to player agent if available
func (seat *Seat) WriteJSON(json cast.JSON) {
	if r := seat.Receiver; r != nil {
		r.WriteJSON(json)
	}
}

// func (seat *Seat) HasAwakePresentCards() bool {
// 	for _, card := range seat.Present {
// 		if card.IsAwake {
// 			return true
// 		}
// 	}
// 	return false
// }

func (seat *Seat) HasPresent(tid string) bool {
	for _, token := range seat.Present {
		if token.ID == tid {
			return true
		}
	}
	return false
}

func (seat *Seat) HasPastCard(cid string) bool {
	for _, card := range seat.Past {
		if card.ID == cid {
			return true
		}
	}
	return false
}

// func (seat *Seat) HasActiveElements() bool {
// 	return len(seat.Elements.GetActive()) > 0
// }

func (seat *Seat) HasCardInHand(cid string) bool {
	for _, card := range seat.Hand {
		if card.ID == cid {
			return true
		}
	}
	return false
}

func (seat *Seat) HasCardsInHand() bool {
	return len(seat.Hand) > 0
}

func (seat *Seat) String() string {
	return cast.StringN(
		`{`,
		seat.Username,
		` ♥:`, seat.Life,
		` ☼:`, seat.Karma.String(),
		` ♣:`, len(seat.Hand),
		` ◘:`, len(seat.Deck.Cards),
		`}`,
	)
}

// JSON returns JSON representation of a game seat
func (seat *Seat) JSON() cast.JSON {
	past := make([]string, 0, len(seat.Past))
	for _, c := range seat.Past {
		past = append(past, c.ID)
	}
	present := make([]string, 0, len(seat.Present))
	for _, t := range seat.Present {
		present = append(present, t.ID)
	}
	return cast.JSON{
		"username": seat.Username,
		"deck":     len(seat.Deck.Cards),
		"life":     seat.Life,
		"present":  present,
		"hand":     len(seat.Hand),
		"elements": seat.Karma.JSON(),
		"past":     past,
		"future":   len(seat.Deck.Cards),
	}
}
