package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func init() {
	engine.Scripts["new-element"] = NewElement
}

func NewElement(game *game.T, seat *game.Seat, target interface{}) game.Event {
	return &NewElementEvent{
		Stack: game.State,
	}
}

type NewElementEvent struct {
	Stack   *game.State
	Element vii.Element
}

func (event *NewElementEvent) Name() string {
	return "choice"
}

// OnActivate implements game.ActivateEventer
func (event *NewElementEvent) OnActivate(game *game.T) {
	animate.NewElementChoice(game.GetSeat(game.State.Seat), game)
}

// OnConnect implements game.ConnectEventer
func (event *NewElementEvent) OnConnect(game *game.T, seat *game.Seat) {
	if game.State.Seat == seat.Username {
		animate.NewElementChoice(seat, game)
	}
}

// Finish implements game.FinishEventer
func (event *NewElementEvent) Finish(game *game.T) {
	if event.Element == vii.ELEMnull {
		game.Log().Warn("scripts/new-element: event: finish: no element")
	} else {
		game.Logger.WithFields(log.Fields{
			"Element": event.Element,
		}).Debug("scripts/new-element: event: finish")
		game.GetSeat(game.State.Seat).Elements.Append(event.Element)
		animate.GameElement(game, game.State.Seat, int(event.Element))
	}
}

// GetStack implements game.StackEventer
func (event *NewElementEvent) GetStack(g *game.T) *game.State {
	return event.Stack
}

// GetNext implements game.StackEventer
func (event *NewElementEvent) GetNext(game *game.T) *game.State {
	return nil
}

func (event *NewElementEvent) Json(game *game.T) js.Object {
	return js.Object{
		"choice": "New Element",
	}
}

func (event *NewElementEvent) Request(game *game.T, seat *game.Seat, json js.Object) {
	if game.State.Seat != seat.Username {
		game.Log().WithFields(log.Fields{
			"Username": seat.Username,
			"Seat":     game.State.Seat,
		}).Warn("scripts/new-element: not your choice")
	} else if e := json.Ival("choice"); e < 1 || e > 7 {
		game.Log().WithFields(log.Fields{
			"choice": e,
		}).Warn("scripts/new-element: invalid element")
	} else {
		event.Element = vii.Element(e)
		game.Log().WithFields(log.Fields{
			"choice": e,
		}).Info("scripts/new-element: confirmed element")
		game.State.Reacts[seat.Username] = "confirm"
	}
}
