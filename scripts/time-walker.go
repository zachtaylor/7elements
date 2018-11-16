package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
	"ztaylor.me/js"
)

func init() {
	engine.Scripts["time-walker"] = TimeWalker
}

func TimeWalker(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	return &TimeWalkerEvent{
		Stack: game.State.Event,
	}
}

type TimeWalkerEvent struct {
	vii.Element
	Stack vii.GameEvent
}

func (event *TimeWalkerEvent) Name() string {
	return "choice"
}

func (event *TimeWalkerEvent) Priority(game *vii.Game) bool {
	return event.Element == vii.ELEMnull
}

func (event *TimeWalkerEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	if seat.Username != game.State.Seat {
		game.Log().Add("Username", seat.Username).Add("HotSeat", game.State.Seat).Warn("script-timewalker-event: not your choice")
		return
	}

	if e := json.Ival("choice"); e < 1 || e > 7 {
		game.Log().Add("Username", seat.Username).Add("Element", e).Warn("script-timewalker-event: invalid element")
		return
	} else {
		event.Element = vii.Element(e)
		game.Log().Add("Username", seat.Username).Add("Element", event.Element).Info("script-timewalker-event: confirmed element")
	}
}

func (event *TimeWalkerEvent) OnStart(game *vii.Game) {
	animate.NewElementChoice(game.GetSeat(game.State.Seat), game)
}

func (event *TimeWalkerEvent) OnReconnect(game *vii.Game, seat *vii.GameSeat) {
	if game.State.Seat == seat.Username {
		animate.NewElementChoice(game.GetSeat(game.State.Seat), game)
	}
}

func (event *TimeWalkerEvent) NextEvent(game *vii.Game) vii.GameEvent {
	seat := game.GetSeat(game.State.Seat)
	seat.Elements.Append(event.Element)
	animate.GameElement(game, game.State.Seat, int(event.Element))
	return event.Stack
}

func (event *TimeWalkerEvent) Json(game *vii.Game) js.Object {
	return js.Object{
		"gameid":   game.Key,
		"choice":   "Time Walker",
		"username": game.State.Seat,
		"timer":    game.State.Timer.Seconds(),
	}
}
