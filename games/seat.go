package games

// import (
// 	"github.com/zachtaylor/7tcg"
// 	"fmt"
// 	"ztaylor.me/js"
// 	"ztaylor.me/log"
// )

// type Seat struct {
// 	Username  string
// 	Life      int
// 	Deck      *vii.GameDeck
// 	Elements  vii.ElementSet
// 	Hand      vii.GameCards
// 	Alive     vii.GameCards
// 	Graveyard vii.GameCards
// 	Color     string
// 	Player
// }

// func newSeat() *vii.GameSeat {
// 	return &Seat{
// 		Elements:  vii.ElementSet{},
// 		Alive:     vii.GameCards{},
// 		Hand:      vii.GameCards{},
// 		Graveyard: vii.GameCards{},
// 	}
// }

// func (s *vii.GameSeat) Start() {
// 	s.Life = 7
// 	s.Deck.Shuffle()
// 	s.DrawCard(3)
// }

// func (seat *vii.GameSeat) DiscardHand() {
// 	for _, card := range seat.Hand {
// 		seat.Graveyard[card.Id] = card
// 	}
// 	seat.Hand = vii.GameCards{}
// }

// func (seat *vii.GameSeat) DrawCard(count int) {
// 	for i := 0; i < count && len(seat.Deck.Cards) > 0; i++ {
// 		card := seat.Deck.Draw()
// 		seat.Hand[card.Id] = card
// 	}
// }

// func (seat *vii.GameSeat) RemoveHandAndElements(gcid string) bool {
// 	if card := seat.Hand[gcid]; card == nil {
// 		return false
// 	} else {
// 		seat.Elements.Deactivate(card.Card.Costs)
// 		delete(seat.Hand, gcid)
// 		return true
// 	}
// }

// func (seat *vii.GameSeat) Reactivate() {
// 	for _, card := range seat.Alive {
// 		card.IsAwake = true
// 	}
// 	seat.Elements.Reactivate()
// }

// func (s *vii.GameSeat) Send(name string, json js.Object) {
// 	if player := s.Player; player != nil {
// 		player.Send(name, json)
// 	} else {
// 		log.Add("Username", s.Username).Add("Name", name).Warn("games: player not in seat")
// 	}
// }

// func (s *vii.GameSeat) Json(showHidden bool) js.Object {
// 	json := js.Object{
// 		"username": s.Username,
// 		"deck":     len(s.Deck.Cards),
// 		"life":     s.Life,
// 		"active":   s.Alive.Json(),
// 		"elements": s.Elements,
// 		"spent":    len(s.Graveyard),
// 	}
// 	if showHidden {
// 		json["hand"] = s.Hand.Json()
// 	} else {
// 		json["hand"] = len(s.Hand)
// 	}
// 	return json
// }

// func (s *vii.GameSeat) String() string {
// 	return fmt.Sprintf("games.Seat{Username:%v, Life:%v, DeckCount:%v}", s.Username, s.Life, len(s.Deck.Cards))
// }

// func (s *vii.GameSeat) HasAwakeAliveCards() bool {
// 	for _, card := range s.Alive {
// 		if card.IsAwake {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (s *vii.GameSeat) HasAliveCard(gcid string) bool {
// 	for _, card := range s.Alive {
// 		if card.Id == gcid {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (s *vii.GameSeat) HasPastCard(gcid string) bool {
// 	for _, card := range s.Graveyard {
// 		if card.Id == gcid {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (s *vii.GameSeat) HasActiveElements() bool {
// 	return len(s.Elements.GetActive()) > 0
// }

// func (s *vii.GameSeat) HasCardsInHand() bool {
// 	return len(s.Hand) > 0
// }

// func (s *vii.GameSeat) HasCardInHand(gcid string) bool {
// 	for _, card := range s.Hand {
// 		if card.Id == gcid {
// 			return true
// 		}
// 	}
// 	return false
// }
