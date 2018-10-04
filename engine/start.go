package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Start(game *vii.Game) Event {
	game.Log().Info("starting")
	return new(StartEvent)
}

type StartEvent bool

func (event *StartEvent) Name() string {
	return "start"
}

func (event *StartEvent) Priority(game *vii.Game, t *Timeline) bool {
	for username, _ := range game.Seats {
		if r := t.Reacts[username]; r != "keep" && r != "mulligan" {
			return true
		}
	}
	return false
}

func (event *StartEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	choice := json["choice"]
	log := game.Log().Add("Seat", seat).Add("Choice", choice)

	if react := t.Reacts[seat.Username]; react != "" {
		log.Add("React", react).Warn("engine-start: receive already recorded")
		return
	} else if choice == "keep" {
		t.Reacts[seat.Username] = "keep"
	} else if choice == "mulligan" {
		t.Reacts[seat.Username] = "mulligan"
		seat.DiscardHand()
		seat.DrawCard(3)
		animate.Hand(game, seat)
		animate.BroadcastSeatUpdate(game, seat)
	} else {
		log.Warn("engine-start: receive unrecognized")
		return
	}

	log.Info("engine-start: received")
	animate.BroadcastAlertTip(game, seat.Username, t.Reacts[seat.Username])
}

func (event *StartEvent) OnStart(game *vii.Game, t *Timeline) {
	game.Log().WithFields(log.Fields{
		"Timeline": t,
		"Seats":    game.Seats,
	}).Debug("engine-start: OnStart")
}

func (event *StartEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *StartEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	return t.Fork(game, Sunrise(game, t, t.HotSeat))
}

func (event *StartEvent) Json(game *vii.Game, t *Timeline) js.Object {
	return js.Object{
		"gameid": game,
		"timer":  t.Lifetime,
	}
}
