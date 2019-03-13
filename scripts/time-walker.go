package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
	"ztaylor.me/js"
)

func init() {
	engine.Scripts["time-walker"] = TimeWalker
}

func TimeWalker(game *game.T, seat *game.Seat, target interface{}) game.Event {
	return &TimeWalkerEvent{
		Stack: game.State.Event,
	}
}

type TimeWalkerEvent struct {
	Stack game.Event
	vii.Element
}

func (event *TimeWalkerEvent) Name() string {
	return "choice"
}

// OnActivate implements game.ActivateEventer
func (event *TimeWalkerEvent) OnActivate(game *game.T) {
	animate.NewElementChoice(game.GetSeat(game.State.Seat), game)
}

// OnConnect implements game.ConnectEventer
func (event *TimeWalkerEvent) OnConnect(game *game.T, seat *game.Seat) {
	if game.State.Seat == seat.Username {
		animate.NewElementChoice(game.GetSeat(game.State.Seat), game)
	}
}

// Finish implements game.FinishEventer
func (event *TimeWalkerEvent) Finish(game *game.T) {
	seat := game.GetSeat(game.State.Seat)
	seat.Elements.Append(event.Element)
	animate.GameElement(game, game.State.Seat, int(event.Element))
}

// GetStack implements game.StackEventer
func (event *TimeWalkerEvent) GetStack(game *game.T) game.Event {
	return event.Stack
}

// GetNext implements game.StackEventer
func (event *TimeWalkerEvent) GetNext(game *game.T) *game.State {
	return nil
}

func (event *TimeWalkerEvent) Json(game *game.T) js.Object {
	return js.Object{
		"choice": "Time Walker",
	}
}

func (event *TimeWalkerEvent) Receive(game *game.T, seat *game.Seat, json js.Object) {
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
