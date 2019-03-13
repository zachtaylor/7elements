package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
)

const NoviceSeerId = "novice-seer"

func init() {
	engine.Scripts[NoviceSeerId] = NoviceSeer
}

func NoviceSeer(game *game.T, seat *game.Seat, target interface{}) game.Event {
	return &NoviceSeerEvent{
		Stack: game.State,
		Card:  seat.Deck.Draw(),
	}
}

type NoviceSeerEvent struct {
	Stack   *game.State
	Card    *game.Card
	Destroy bool
}

func (event *NoviceSeerEvent) Name() string {
	return "choice"
}

// OnActivate implements game.ActivateEventer
func (event *NoviceSeerEvent) OnActivate(game *game.T) {
	animate.NoviceSeerChoice(game.GetSeat(game.State.Seat), game, event.Card)
}

// OnConnect implements game.ConnectEventer
func (event *NoviceSeerEvent) OnConnect(game *game.T, seat *game.Seat) {
	if game.State.Seat == seat.Username {
		animate.NoviceSeerChoice(seat, game, event.Card)
	}
}

// Finish implements game.FinishEventer
func (event *NoviceSeerEvent) Finish(game *game.T) {
	seat := game.GetSeat(game.State.Seat)
	if event.Destroy {
		seat.Past[event.Card.Id] = event.Card
		animate.GameSeat(game, seat)
	} else {
		seat.Deck.Prepend(event.Card)
	}
}

// GetStack implements game.StackEventer
func (event *NoviceSeerEvent) GetStack(g *game.T) *game.State {
	return event.Stack
}

// GetNext implements game.StackEventer
func (event *NoviceSeerEvent) GetNext(game *game.T) *game.State {
	return nil
}

func (event *NoviceSeerEvent) Json(game *game.T) vii.Json {
	return vii.Json{
		"choice": "Novice Seer",
	}
}

func (event *NoviceSeerEvent) Request(game *game.T, seat *game.Seat, json vii.Json) {
	log := game.Log().Add("Username", seat.Username)

	if seat.Username != game.State.Seat {
		log.Add("HotSeat", game.State.Seat).Warn("scripts/novice-seer: not your choice")
		return
	}

	switch json.Sval("choice") {
	case "yes":
		event.Destroy = true
		fallthrough
	case "no":
		log.Add("Choice", json.Sval("choice")).Debug("scripts/novice-seer: confirm")
		game.State.Reacts[seat.Username] = "confirm"
	default:
		log.Add("Choice", json.Sval("choice")).Warn("scripts/novice-seer: unrecognized choice")
	}
}
