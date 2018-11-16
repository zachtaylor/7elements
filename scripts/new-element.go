package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
	"ztaylor.me/js"
)

func init() {
	engine.Scripts["new-element"] = NewElement
}

func NewElement(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	return &NewElementMode{
		Stack: game.State.Event,
	}
}

type NewElementMode struct {
	vii.Element
	Stack vii.GameEvent
}

func (event *NewElementMode) Name() string {
	return "choice"
}

func (event *NewElementMode) Priority(game *vii.Game) bool {
	return event.Element == vii.ELEMnull
}

func (event *NewElementMode) OnStart(game *vii.Game) {
	animate.NewElementChoice(game.GetSeat(game.State.Seat), game)
}

func (event *NewElementMode) OnReconnect(game *vii.Game, seat *vii.GameSeat) {
	if game.State.Seat == seat.Username {
		animate.NewElementChoice(seat, game)
	}
}

func (event *NewElementMode) NextEvent(game *vii.Game) vii.GameEvent {
	seat := game.GetSeat(game.State.Seat)
	seat.Elements.Append(event.Element)
	animate.GameElement(game, game.State.Seat, int(event.Element))
	return event.Stack
}

func (event *NewElementMode) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	if game.State.Seat != seat.Username {
		game.Log().Add("Username", seat.Username).Add("HotSeat", game.State.Seat).Warn("gameseat.NewElementMode: not your choice")
		return
	}

	if e := json.Ival("choice"); e < 1 || e > 7 {
		game.Log().Add("Username", seat.Username).Add("Element", e).Warn("gameseat.NewElementMode: invalid element")
	} else {
		event.Element = vii.Element(e)
		game.Log().Add("Username", seat.Username).Add("Element", event.Element).Info("gameseat.NewElement: confirmed element")
	}
}

func (event *NewElementMode) Json(game *vii.Game) js.Object {
	return js.Object{
		"gameid":   game.Key,
		"choice":   "New Element",
		"username": game.State.Seat,
		"timer":    game.State.Timer.Seconds(),
	}
}
