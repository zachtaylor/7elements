package engine

import (
	"elemen7s.com"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Main(game *vii.Game, past *Timeline) Event {
	if tname := past.Name(); tname != "sunrise" {
		game.Log().Add("Timeline", tname).Error("main can only follow sunrise")
		return nil
	}
	game.Log().Info("main")
	return new(MainEvent)
}

type MainEvent bool

func (event *MainEvent) Name() string {
	return "main"
}

func (event *MainEvent) Priority(game *vii.Game, t *Timeline) bool {
	return t.Reacts[t.HotSeat] != "pass"
}

func (event *MainEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-main: receive")
}

func (event *MainEvent) OnStart(game *vii.Game, t *Timeline) {
	game.Log().WithFields(log.Fields{
		"Seat": game.GetSeat(t.HotSeat),
	}).Warn("engine-main: OnStart")
}

func (event *MainEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *MainEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	return t.Fork(game, Attack(game, t))
}

func (event *MainEvent) Json(game *vii.Game, t *Timeline) js.Object {
	seat := game.GetSeat(t.HotSeat)
	return js.Object{
		"gameid":   game,
		"timer":    t.Lifetime.Seconds(),
		"username": t.HotSeat,
		"life":     seat.Life,
		"hand":     len(seat.Hand),
		"deck":     len(seat.Deck.Cards),
		"elements": seat.Elements,
		"spent":    len(seat.Graveyard),
	}
}
