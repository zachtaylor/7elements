package engine

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Sunset(game *vii.Game) vii.GameEvent {
	game.Log().Info("sunset")
	return new(SunsetEvent)
}

type SunsetEvent bool

func (event *SunsetEvent) Name() string {
	return "sunset"
}

func (event *SunsetEvent) Priority(game *vii.Game) bool {
	for username, _ := range game.Seats {
		if r := game.State.Reacts[username]; r != "pass" {
			return true
		}
	}
	return false
}

func (event *SunsetEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-sunset: receive")
}

func (event *SunsetEvent) OnStart(game *vii.Game) {
}

func (event *SunsetEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event *SunsetEvent) NextEvent(game *vii.Game) vii.GameEvent {
	return Sunrise(game, game.GetOpponentSeat(game.State.Seat).Username)
}

func (event *SunsetEvent) Json(game *vii.Game) js.Object {
	return js.Object{
		"gameid":   game.Key,
		"username": game.State.Seat,
		"timer":    game.State.Timer.Seconds(),
	}
}
