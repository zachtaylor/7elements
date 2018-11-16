package engine

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Main(game *vii.Game) vii.GameEvent {
	return new(MainEvent)
}

type MainEvent bool

func (event *MainEvent) Name() string {
	return "main"
}

func (event *MainEvent) Priority(game *vii.Game) bool {
	return game.State.Reacts[game.State.Seat] != "pass"
}

func (event *MainEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-main: receive")
}

func (event *MainEvent) OnStart(game *vii.Game) {
	game.Log().WithFields(log.Fields{
		"Seat": game.GetSeat(game.State.Seat),
	}).Warn("engine-main: OnStart")
}

func (event *MainEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event *MainEvent) NextEvent(game *vii.Game) vii.GameEvent {
	if game.State.Reacts[game.State.Seat] == "attack" {
		return Attack(game)
	}
	return Sunset(game)
}

func (event *MainEvent) Json(game *vii.Game) vii.Json {
	seat := game.GetSeat(game.State.Seat)
	return vii.Json{
		"gameid":   game.Key,
		"timer":    game.State.Timer.Seconds(),
		"username": game.State.Seat,
		"life":     seat.Life,
		"hand":     len(seat.Hand),
		"deck":     len(seat.Deck.Cards),
		"elements": seat.Elements,
		"spent":    len(seat.Graveyard),
	}
}
