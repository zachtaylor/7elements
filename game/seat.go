package game

import (
	"fmt"

	"github.com/zachtaylor/7elements"
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
	receiver vii.JsonWriter
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

func (seat *Seat) Login(game *T, player vii.JsonWriter) {
	seat.receiver = player
	game.Request(seat.Username, "connect", nil)
}

func (seat *Seat) Logout(game *T) {
	seat.receiver = nil
	game.Request(seat.Username, "disconnect", nil)
}

func (seat *Seat) WriteJson(json vii.Json) {
	if player := seat.receiver; player != nil {
		player.WriteJson(json)
	}
}

func (seat *Seat) Json(showHidden bool) vii.Json {
	json := vii.Json{
		"username": seat.Username,
		"deck":     len(seat.Deck.Cards),
		"life":     seat.Life,
		"active":   seat.Present.Json(),
		"elements": seat.Elements,
		"past":     len(seat.Past),
		"future":   len(seat.Deck.Cards),
	}
	if showHidden {
		json["hand"] = seat.Hand.Json()
	} else {
		json["hand"] = len(seat.Hand)
	}
	return json
}

func (seat *Seat) String() string {
	return fmt.Sprintf("vii.Seat{Name:%v, ♥:%v, Deck:%v}", seat.Username, seat.Life, len(seat.Deck.Cards))
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
