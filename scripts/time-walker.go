package scripts

import (
	"github.com/zachtaylor/7tcg"
	"github.com/zachtaylor/7tcg/animate"
	"github.com/zachtaylor/7tcg/engine"
	"ztaylor.me/js"
)

func init() {
	engine.Scripts["time-walker"] = TimeWalker
}

func TimeWalker(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	return t.Fork(game, &TimeWalkerEvent{
		Stack: t,
	})
}

type TimeWalkerEvent struct {
	vii.Element
	Stack *engine.Timeline
}

func (event *TimeWalkerEvent) Name() string {
	return "choice"
}

func (event *TimeWalkerEvent) Priority(game *vii.Game, t *engine.Timeline) bool {
	return event.Element == vii.ELEMnull
}

func (event *TimeWalkerEvent) Receive(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, json js.Object) {
	if seat.Username != t.HotSeat {
		game.Log().Add("Username", seat.Username).Add("HotSeat", t.HotSeat).Warn("script-timewalker-event: not your choice")
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

func (event *TimeWalkerEvent) OnStart(game *vii.Game, t *engine.Timeline) {
	animate.NewElementChoice(game.GetSeat(t.HotSeat), game)
}

func (event *TimeWalkerEvent) OnReconnect(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat) {
	if t.HotSeat == seat.Username {
		animate.NewElementChoice(game.GetSeat(t.HotSeat), game)
	}
}

func (event *TimeWalkerEvent) OnStop(game *vii.Game, t *engine.Timeline) *engine.Timeline {
	seat := game.GetSeat(t.HotSeat)
	seat.Elements.Append(event.Element)
	animate.BroadcastAddElement(game, t.HotSeat, int(event.Element))
	return event.Stack
}

func (event *TimeWalkerEvent) Json(game *vii.Game, t *engine.Timeline) js.Object {
	seat := game.GetSeat(t.HotSeat)
	return js.Object{
		"gameid":   game,
		"choice":   "Time Walker",
		"username": seat.Username,
		"timer":    t.Lifetime.Seconds(),
	}
}
