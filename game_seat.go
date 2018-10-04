package vii

import "fmt"

type GameSeat struct {
	GameKey   string
	Username  string
	Life      int
	Deck      *GameDeck
	Elements  ElementSet
	Hand      GameCards
	Alive     GameCards
	Graveyard GameCards
	Color     string
	Receiver  JsonWriter
}

type Receiver interface {
	Send(string, Json)
}

func (game *Game) NewSeat() *GameSeat {
	return &GameSeat{
		GameKey:   game.Key,
		Elements:  ElementSet{},
		Alive:     GameCards{},
		Hand:      GameCards{},
		Graveyard: GameCards{},
	}
}

func (seat *GameSeat) Start() {
	seat.Life = 7
	seat.Deck.Shuffle()
	seat.DrawCard(3)
}

func (seat *GameSeat) DiscardHand() {
	for _, card := range seat.Hand {
		seat.Graveyard[card.Id] = card
	}
	seat.Hand = GameCards{}
}

func (seat *GameSeat) DrawCard(count int) {
	for i := 0; i < count && len(seat.Deck.Cards) > 0; i++ {
		card := seat.Deck.Draw()
		seat.Hand[card.Id] = card
	}
}

func (seat *GameSeat) Reactivate() {
	for _, card := range seat.Alive {
		card.IsAwake = true
	}
	seat.Elements.Reactivate()
}

func (seat *GameSeat) WriteJson(json Json) {
	if player := seat.Receiver; player != nil {
		player.WriteJson(json)
	}
}

func (seat *GameSeat) Json(showHidden bool) Json {
	json := Json{
		"username": seat.Username,
		"deck":     len(seat.Deck.Cards),
		"life":     seat.Life,
		"active":   seat.Alive.Json(),
		"elements": seat.Elements,
		"spent":    len(seat.Graveyard),
	}
	if showHidden {
		json["hand"] = seat.Hand.Json()
	} else {
		json["hand"] = len(seat.Hand)
	}
	return json
}

func (seat *GameSeat) String() string {
	return fmt.Sprintf("vii.GameSeat{Name:%v, ♥:%v, Deck:%v}", seat.Username, seat.Life, len(seat.Deck.Cards))
}

func (seat *GameSeat) HasAwakeAliveCards() bool {
	for _, card := range seat.Alive {
		if card.IsAwake {
			return true
		}
	}
	return false
}

func (seat *GameSeat) HasAliveCard(gcid string) bool {
	for _, card := range seat.Alive {
		if card.Id == gcid {
			return true
		}
	}
	return false
}

func (seat *GameSeat) HasPastCard(gcid string) bool {
	for _, card := range seat.Graveyard {
		if card.Id == gcid {
			return true
		}
	}
	return false
}

func (seat *GameSeat) HasActiveElements() bool {
	return len(seat.Elements.GetActive()) > 0
}

func (seat *GameSeat) HasCardsInHand() bool {
	return len(seat.Hand) > 0
}

func (seat *GameSeat) HasCardInHand(gcid string) bool {
	for _, card := range seat.Hand {
		if card.Id == gcid {
			return true
		}
	}
	return false
}
