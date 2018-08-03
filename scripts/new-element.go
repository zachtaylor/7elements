package scripts

import (
	"github.com/zachtaylor/7tcg"
	"github.com/zachtaylor/7tcg/animate"
	"github.com/zachtaylor/7tcg/engine"
	"ztaylor.me/js"
)

func init() {
	engine.Scripts["new-element"] = NewElement
}

func NewElement(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	return t.Fork(game, &NewElementMode{
		Stack: t,
	})
}

type NewElementMode struct {
	vii.Element
	Stack *engine.Timeline
}

func (event *NewElementMode) Name() string {
	return "choice"
}

func (event *NewElementMode) Priority(game *vii.Game, t *engine.Timeline) bool {
	return event.Element == vii.ELEMnull
}

func (event *NewElementMode) Json(game *vii.Game, t *engine.Timeline) js.Object {
	seat := game.GetSeat(t.HotSeat)
	return js.Object{
		"gameid":   game,
		"choice":   "New Element",
		"username": seat.Username,
		"timer":    t.Lifetime.Seconds(),
	}
}

func (event *NewElementMode) OnStart(game *vii.Game, t *engine.Timeline) {
	animate.NewElementChoice(game.GetSeat(t.HotSeat), game)
}

func (event *NewElementMode) OnReconnect(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat) {
	if t.HotSeat == seat.Username {
		animate.NewElementChoice(seat, game)
	}
}

func (event *NewElementMode) OnStop(game *vii.Game, t *engine.Timeline) *engine.Timeline {
	seat := game.GetSeat(t.HotSeat)
	seat.Elements.Append(event.Element)
	animate.BroadcastAddElement(game, t.HotSeat, int(event.Element))
	return event.Stack
}

func (event *NewElementMode) Receive(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, json js.Object) {
	if t.HotSeat != seat.Username {
		game.Log().Add("Username", seat.Username).Add("HotSeat", t.HotSeat).Warn("gameseat.NewElementMode: not your choice")
		return
	}

	if e := json.Ival("choice"); e < 1 || e > 7 {
		game.Log().Add("Username", seat.Username).Add("Element", e).Warn("gameseat.NewElementMode: invalid element")
		return
	} else {
		event.Element = vii.Element(e)
		game.Log().Add("Username", seat.Username).Add("Element", event.Element).Info("gameseat.NewElement: confirmed element")
	}

}
