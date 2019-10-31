package game

import (
	"fmt"

	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/cast"
)

type Seat struct {
	Username string
	Life     int
	Deck     *Deck
	Elements vii.ElementSet
	Hand     Cards
	Present  Cards
	Past     Cards
	Color    string
	Receiver vii.JSONWriter
}

func (game *T) NewSeat() *Seat {
	return &Seat{
		Elements: vii.ElementSet{},
		Present:  Cards{},
		Hand:     Cards{},
		Past:     Cards{},
	}
}

func (seat *Seat) DrawCard(count int) {
	for i := 0; i < count && len(seat.Deck.Cards) > 0; i++ {
		card := seat.Deck.Draw()
		seat.Hand[card.Id] = card
	}
}

func (seat *Seat) DiscardHand() {
	for _, card := range seat.Hand {
		seat.Past[card.Id] = card
	}
	seat.Hand = Cards{}
}

func (seat *Seat) Reactivate() {
	for _, card := range seat.Present {
		card.IsAwake = true
	}
	seat.Elements.Reactivate()
}

// Send sends data to player agent if available
func (seat *Seat) Send(json cast.JSON) {
	if r := seat.Receiver; r != nil {
		r.WriteJSON(json)
	}
}

func (seat *Seat) SendHandUpdate() {
	seat.Send(BuildPushJSON("/game/hand", cast.JSON{
		"cards": seat.Hand.JSON(),
	}))
}

func (seat *Seat) SendError(source, message string) {
	seat.Send(BuildPushJSON("/game/error", cast.JSON{
		"source":  source,
		"message": message,
	}))
}

func (seat *Seat) HasAwakePresentCards() bool {
	for _, card := range seat.Present {
		if card.IsAwake {
			return true
		}
	}
	return false
}

func (seat *Seat) HasPresentCard(gcid string) bool {
	for _, card := range seat.Present {
		if card.Id == gcid {
			return true
		}
	}
	return false
}

func (seat *Seat) HasPastCard(gcid string) bool {
	for _, card := range seat.Past {
		if card.Id == gcid {
			return true
		}
	}
	return false
}

func (seat *Seat) HasActiveElements() bool {
	return len(seat.Elements.GetActive()) > 0
}

func (seat *Seat) HasCardsInHand() bool {
	return len(seat.Hand) > 0
}

func (seat *Seat) HasCardInHand(gcid string) bool {
	for _, card := range seat.Hand {
		if card.Id == gcid {
			return true
		}
	}
	return false
}

func (seat *Seat) String() string {
	return fmt.Sprintf("vii.Seat{%s}", seat.Print())
}

// Print returns a detailed compressed string representation
func (seat *Seat) Print() string {
	return fmt.Sprintf("%s ☼:%s ◘:%d ♥:%d ♣:%d",
		seat.Username,
		seat.Elements.String(),
		len(seat.Deck.Cards),
		seat.Life,
		len(seat.Hand),
	)
}

// JSON returns JSON representation of a game seat
func (seat *Seat) JSON() cast.JSON {
	return cast.JSON{
		"username": seat.Username,
		"deck":     len(seat.Deck.Cards),
		"life":     seat.Life,
		"active":   seat.Present.JSON(),
		"hand":     len(seat.Hand),
		"elements": seat.Elements.JSON(),
		"past":     seat.Past.JSON(),
		"future":   len(seat.Deck.Cards),
	}
}
