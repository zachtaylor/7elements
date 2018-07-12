package engine

import (
	"elemen7s.com"
	"elemen7s.com/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Play(game *vii.Game, past *Timeline, seat *vii.GameSeat, card *vii.GameCard, target interface{}) Event {
	game.Log().Info("play")

	return &PlayEvent{
		Stack:  past,
		Card:   card,
		Target: target,
	}
}

type PlayEvent struct {
	Stack       *Timeline
	Card        *vii.GameCard
	IsCancelled bool
	Target      interface{}
}

func (event *PlayEvent) Name() string {
	return "play"
}

func (event *PlayEvent) Priority(game *vii.Game, t *Timeline) bool {
	return t.Reacts[game.GetOpponentSeat(t.HotSeat).Username] != "pass"
}

func (event *PlayEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-play: receive")
}

func (event *PlayEvent) OnStart(game *vii.Game, t *Timeline) {
	seat := game.GetSeat(t.HotSeat)
	game.Log().WithFields(log.Fields{
		"Seat":     seat,
		"Elements": seat.Elements,
	}).Warn("engine-play: OnStart")
}

func (event *PlayEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *PlayEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	seat := game.GetSeat(t.HotSeat)
	log := game.Log().WithFields(log.Fields{
		"Seat": seat,
		"Card": event.Card,
		"Hand": seat.Hand,
	})

	game.Send("resolve", js.Object{
		"gameid":   game,
		"username": seat.Username,
		"card":     event.Card.Json(),
	})

	playPower := event.Card.Card.GetPlayPower()
	if playPower == nil {
		animate.Error(game, game, event.Card.CardText.Name, "card does not work yet")
		log.Warn("engine-play: resolve; card does not work")
		return event.Stack
	}

	if event.Card.Card.CardType == vii.CTYPbody || event.Card.Card.CardType == vii.CTYPitem {
		seat.Alive[event.Card.Id] = event.Card
		animate.Spawn(game, event.Card)
	} else if event.Card.Card.CardType == vii.CTYPspell {
		seat.Graveyard[event.Card.Id] = event.Card
	}

	if playPower != nil {
		Script(game, t, seat, playPower, event.Target)
	}

	return event.Stack
}

func (event *PlayEvent) Json(game *vii.Game, t *Timeline) js.Object {
	seat := game.GetSeat(t.HotSeat)
	return js.Object{
		"gameid":   game,
		"timer":    t.Lifetime.Seconds(),
		"username": seat.Username,
		"elements": seat.Elements,
		"hand":     len(seat.Hand),
		"card":     event.Card.Json(),
		"target":   event.Target,
	}
}
