package gameseats

import (
	"7elements.ztaylor.me/games/cards"
	"7elements.ztaylor.me/games/decks"
	"7elements.ztaylor.me/games/elements"
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
)

type GameSeat struct {
	Username string
	DeckId   int
	Deckname string
	Life     int
	Deck     *gamedecks.GameDeck
	Hand     []*gamecards.GameCard
	Active   []*gamecards.GameCard
	Elements gameelements.GameElements
	Spent    []*gamecards.GameCard
	*sessions.Socket
}

func New() *GameSeat {
	return &GameSeat{
		Deck:     gamedecks.New(),
		Hand:     make([]*gamecards.GameCard, 0),
		Active:   make([]*gamecards.GameCard, 0),
		Elements: gameelements.New(),
		Spent:    make([]*gamecards.GameCard, 0),
	}
}

func (seat *GameSeat) Reactivate() {
	for _, card := range seat.Active {
		card.Active = true
	}
	seat.Elements.Reactivate()
}

func (seat *GameSeat) CardHandPositionGCID(gcid int) int {
	for i, card := range seat.Hand {
		if gcid == card.GameCardId {
			return i
		}
	}
	return -1
}

func (seat *GameSeat) ActiveCardGCID(gcid int) *gamecards.GameCard {
	for _, card := range seat.Active {
		if gcid == card.GameCardId {
			return card
		}
	}
	return nil
}

func (seat *GameSeat) RemoveHandPosition(i int) *gamecards.GameCard {
	card := seat.Hand[i]
	copy(seat.Hand[i:], seat.Hand[i+1:])
	seat.Hand[len(seat.Hand)-1] = nil
	seat.Hand = seat.Hand[:len(seat.Hand)-1]
	return card
}

func (seat *GameSeat) Json() json.Json {
	return json.Json{
		"username": seat.Username,
		"deckname": seat.Deckname,
		"deck":     len(seat.Deck.Cards),
		"life":     seat.Life,
		"active":   gamecards.Stack(seat.Active).Json(),
		"elements": seat.Elements.Copy(),
		"spent":    gamecards.Stack(seat.Spent).Json(),
	}
}
