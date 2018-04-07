package games

import (
	"elemen7s.com"
	"fmt"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

type Seat struct {
	Username  string
	Life      int
	Deck      *Deck
	Elements  vii.ElementSet
	Hand      Cards
	Alive     Cards
	Graveyard Cards
	Color     string
	Player
}

func newSeat() *Seat {
	return &Seat{
		Alive:     Cards{},
		Hand:      Cards{},
		Graveyard: Cards{},
		Elements:  vii.ElementSet{},
	}
}

func (s *Seat) Start() {
	s.Life = 7
	s.Deck.Shuffle()
	s.DrawCard(3)
}

func (seat *Seat) DiscardHand() {
	for _, card := range seat.Hand {
		seat.Graveyard[card.Id] = card
	}
	seat.Hand = Cards{}
}

func (seat *Seat) DrawCard(count int) {
	for i := 0; i < count && len(seat.Deck.Cards) > 0; i++ {
		card := seat.Deck.Draw()
		seat.Hand[card.Id] = card
	}
}

func (seat *Seat) RemoveHandAndElements(gcid int) bool {
	if card := seat.Hand[gcid]; card == nil {
		return false
	} else {
		seat.Elements.Deactivate(card.Card.Costs)
		delete(seat.Hand, gcid)
		return true
	}
}

func (seat *Seat) Reactivate() {
	for _, card := range seat.Alive {
		card.Awake = true
	}
	seat.Elements.Reactivate()
}

func (s *Seat) Send(name string, json js.Object) {
	if player := s.Player; player != nil {
		player.Send(name, json)
	} else {
		log.Add("Username", s.Username).Add("Name", name).Warn("games: player not in seat")
	}
}

func (s *Seat) Json(showHidden bool) js.Object {
	json := js.Object{
		"username": s.Username,
		"deck":     len(s.Deck.Cards),
		"life":     s.Life,
		"active":   s.Alive.Json(),
		"elements": s.Elements,
		"spent":    len(s.Graveyard),
	}
	if showHidden {
		json["hand"] = s.Hand.Json()
	} else {
		json["hand"] = len(s.Hand)
	}
	return json
}

func (s *Seat) String() string {
	return fmt.Sprintf("games.Seat{Username:%v, Life:%v, DeckCount:%v}", s.Username, s.Life, len(s.Deck.Cards))
}

func (s *Seat) HasAwakeAliveCards() bool {
	for _, card := range s.Alive {
		if card.Awake {
			return true
		}
	}
	return false
}

func (s *Seat) HasActiveElements() bool {
	return len(s.Elements.GetActive()) > 0
}

func (s *Seat) HasCardsInHand() bool {
	return len(s.Hand) > 0
}

func (s *Seat) HasCardInHand(gcid int) bool {
	for _, card := range s.Hand {
		if card.Id == gcid {
			return true
		}
	}
	return false
}
