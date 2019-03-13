package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

type CombatEvent struct {
	A *game.Card
	B *game.Card
}

func Combat(game *game.T, acard *game.Card, dcard *game.Card) game.Event {
	return &CombatEvent{acard, dcard}
}

func (event *CombatEvent) Name() string {
	return "combat"
}

// // OnActivate implements game.ActivateEventer
// func (event *CombatEvent) OnActivate(*game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *CombatEvent) OnConnect(*game.T, *game.Seat) {
// }

// // GetStack implements game.StackEventer
// func (event *CombatEvent) GetStack(game *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestEventer
// func (event *CombatEvent) Request(game *game.T, seat *game.Seat, json vii.Json) {
// }

// Finish implements game.FinishEventer
func (event *CombatEvent) Finish(g *game.T) {
	if event.B != nil {
		Damage(g, event.B, event.A.Body.Attack)
		Damage(g, event.A, event.B.Body.Attack)
	} else if seat := g.GetOpponentSeat(event.A.Username); seat == nil {

	} else {
		seat.Life -= event.A.Body.Attack
	}
}

func (event *CombatEvent) GetNext(game *game.T) *game.State {
	if event.A == nil {
	} else if seat := game.GetOpponentSeat(event.A.Username); seat == nil {
	} else if seat.Life < 1 {
		return game.NewState("", End(game, event.A.Username, seat.Username))
	}
	return game.NewState(game.State.Seat, Main(game))
}

func (event *CombatEvent) Json(game *game.T) vii.Json {
	return vii.Json{
		"attack": event.A.Json(),
		"block":  event.B.Json(),
	}
}

func Damage(game *game.T, card *game.Card, n int) {
	card.Body.Health -= n
	if card.Body.Health < 1 {
		seat := game.GetSeat(card.Username)
		delete(seat.Present, card.Id)

		if !card.IsToken {
			seat.Past[card.Id] = card
		}
	}
}
