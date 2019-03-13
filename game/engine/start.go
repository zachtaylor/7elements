package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/js"
)

func Start(game *game.T) game.Event {
	game.Log().Debug("engine/start: Start")
	for _, seat := range game.Seats {
		seat.Life = 7
	}
	return new(StartEvent)
}

type StartEvent bool

func (event *StartEvent) Name() string {
	return "start"
}

// OnActivate implements game.ActivateEventer
func (event *StartEvent) OnActivate(game *game.T) {
	game.Log().Debug("engine/start: OnStart")
	for _, seat := range game.Seats {
		seat.Deck.Shuffle()
		seat.DrawCard(3)
		animate.GameHand(game, seat)
	}
}

// // OnConnect implements game.ConnectEventer
// func (event *StartEvent) OnConnect(game *game.T, seat *game.Seat) {
// }

// // Finish implements game.FinishEventer
// func (event *StartEvent) Finish(*game.T) {
// }

// // GetStack implements game.StackEventer
// func (event *StartEvent) GetStack(game *game.T) *game.State {
// 	return nil
// }

func (event *StartEvent) GetNext(game *game.T) *game.State {
	return game.NewState(game.State.Seat, Sunrise(game))
}

func (event *StartEvent) Json(game *game.T) vii.Json {
	return nil
}

// Request implements game.RequestEventer
func (event *StartEvent) Request(game *game.T, seat *game.Seat, json js.Object) {
	choice := json.Sval("choice")
	log := game.Log().Add("Seat", seat).Add("Choice", choice)

	if react := game.State.Reacts[seat.Username]; react != "" {
		log.Add("React", react).Warn("engine/start: receive already recorded")
		return
	} else if choice == "keep" {
		game.State.Reacts[seat.Username] = "keep"
	} else if choice == "mulligan" {
		game.State.Reacts[seat.Username] = "mulligan"
		seat.DiscardHand()
		seat.DrawCard(3)
		animate.GameHand(game, seat)
		animate.GameSeat(game, seat)
	} else {
		log.Warn("engine/start: receive unrecognized")
		return
	}

	animate.GameReact(game, seat.Username)
	log.Debug("engine/start: receive")
	game.State.Reacts[seat.Username] = choice
}
