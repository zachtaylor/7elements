package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Play(game *vii.Game, seat *vii.GameSeat, card *vii.GameCard, target interface{}) vii.GameEvent {
	game.Log().Info("play")

	return &PlayEvent{
		Stack:  game.State.Event,
		Card:   card,
		Target: target,
	}
}

type PlayEvent struct {
	Stack       vii.GameEvent
	Card        *vii.GameCard
	IsCancelled bool
	Target      interface{}
}

func (event *PlayEvent) Name() string {
	return "play"
}

func (event *PlayEvent) Priority(game *vii.Game) bool {
	return game.State.Reacts[game.GetOpponentSeat(game.State.Seat).Username] != "pass"
}

func (event *PlayEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-play: receive")
}

func (event *PlayEvent) OnStart(game *vii.Game) {
	seat := game.GetSeat(game.State.Seat)
	game.Log().WithFields(log.Fields{
		"Seat":     seat,
		"Elements": seat.Elements,
	}).Warn("engine-play: OnStart")
}

func (event *PlayEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event *PlayEvent) NextEvent(game *vii.Game) vii.GameEvent {
	seat := game.GetSeat(game.State.Seat)
	log := game.Log().WithFields(log.Fields{
		"Seat": seat,
		"Card": event.Card,
		"Hand": seat.Hand,
	})

	game.WriteJson(animate.Build("/game/resolve", js.Object{
		"gameid":   game.Key,
		"username": seat.Username,
		"card":     event.Card.Json(),
	}))

	playPower := event.Card.Card.GetPlayPower()
	if playPower == nil {
		animate.GameError(game, game, event.Card.Card.Name, "card does not work yet")
		log.Warn("engine-play: resolve; card does not work")
		return event.Stack
	}

	if event.Card.Card.Type == vii.CTYPbody || event.Card.Card.Type == vii.CTYPitem {
		seat.Alive[event.Card.Id] = event.Card
		animate.GameSpawn(game, event.Card)
	} else if event.Card.Card.Type == vii.CTYPspell {
		seat.Graveyard[event.Card.Id] = event.Card
	}

	if playPower != nil {
		Script(game, seat, playPower, event.Target)
	}

	return event.Stack
}

func (event *PlayEvent) Json(game *vii.Game) vii.Json {
	seat := game.GetSeat(game.State.Seat)
	return js.Object{
		"gameid":   game.Key,
		"timer":    game.State.Timer.Seconds(),
		"username": seat.Username,
		"elements": seat.Elements,
		"hand":     len(seat.Hand),
		"card":     event.Card.Json(),
		"target":   event.Target,
	}
}
