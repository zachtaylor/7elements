package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
)

func Start(game *vii.Game) vii.GameEvent {
	game.Log().Info("Start")
	for _, seat := range game.Seats {
		seat.Life = 7
	}
	return new(StartEvent)
}

type StartEvent bool

func (event *StartEvent) Name() string {
	return "start"
}

func (event *StartEvent) Priority(game *vii.Game) bool {
	return len(game.State.Reacts) < len(game.Seats)
}

func (event *StartEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	choice := json["choice"]
	log := game.Log().Add("Seat", seat).Add("Choice", choice)

	if react := game.State.Reacts[seat.Username]; react != "" {
		log.Add("React", react).Warn("Start Receive already recorded")
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
		log.Warn("start.Receive unrecognized")
		return
	}

	animate.GameReact(game, seat.Username)
	// animate.GameEvent(game)
	// game.WriteJson(animate.AlertInfo("Start", seat.Username+" "+game.State.Reacts[seat.Username]))
	log.Debug("start.Receive")
}

func (event *StartEvent) OnStart(game *vii.Game) {
	for _, seat := range game.Seats {
		seat.Deck.Shuffle()
		seat.DrawCard(3)
		animate.GameHand(game, seat)
	}
	game.Log().Debug("start.OnStart")
}

func (event *StartEvent) OnReconnect(game *vii.Game, seat *vii.GameSeat) {
}

func (event *StartEvent) NextEvent(game *vii.Game) vii.GameEvent {
	if len(game.State.Reacts) < len(game.Seats) {
		return End(game, "", "") // nobody wins
	}
	return Sunrise(game, game.State.Seat)
}

func (event *StartEvent) Json(game *vii.Game) js.Object {
	return js.Object{
		"gameid": game.Key,
		"timer":  game.State.Timer,
	}
}
