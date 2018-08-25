package engine

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Sunset(game *vii.Game, past *Timeline) Event {
	if tname := past.Name(); tname != "defend" {
		game.Log().Add("Timeline", tname).Error("sunset can only follow defend")
		return nil
	}
	game.Log().Info("sunset")

	return new(SunsetEvent)
}

type SunsetEvent bool

func (event *SunsetEvent) Name() string {
	return "sunset"
}

func (event *SunsetEvent) Priority(game *vii.Game, t *Timeline) bool {
	for username, _ := range game.Seats {
		if r := t.Reacts[username]; r != "pass" {
			return true
		}
	}
	return false
}

func (event *SunsetEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-sunset: receive")
}

func (event *SunsetEvent) OnStart(game *vii.Game, t *Timeline) {
}

func (event *SunsetEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *SunsetEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	return t.Fork(game, Sunrise(game, t, game.GetOpponentSeat(t.HotSeat).Username))
}

func (event *SunsetEvent) Json(game *vii.Game, t *Timeline) js.Object {
	return js.Object{
		"gameid":   game,
		"username": t.HotSeat,
		"timer":    t.Lifetime.Seconds(),
	}
}
