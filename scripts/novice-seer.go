package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
	"ztaylor.me/js"
)

const NoviceSeerId = "novice-seer"

func init() {
	engine.Scripts[NoviceSeerId] = NoviceSeer
}

func NoviceSeer(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	return &NoviceSeerEvent{
		Card:  seat.Deck.Draw(),
		Stack: game.State.Event,
	}
}

type NoviceSeerEvent struct {
	Destroy *bool
	Card    *vii.GameCard
	Stack   vii.GameEvent
}

func (event *NoviceSeerEvent) Name() string {
	return "choice"
}

func (event *NoviceSeerEvent) Priority(game *vii.Game) bool {
	return event.Destroy == nil
}

func (event *NoviceSeerEvent) OnStart(game *vii.Game) {
	animate.NoviceSeerChoice(game.GetSeat(game.State.Seat), game, event.Card)
}

func (event *NoviceSeerEvent) OnReconnect(game *vii.Game, seat *vii.GameSeat) {
	if game.State.Seat == seat.Username {
		animate.NoviceSeerChoice(seat, game, event.Card)
	}
}

func (event *NoviceSeerEvent) NextEvent(game *vii.Game) vii.GameEvent {
	seat := game.GetSeat(event.Card.Username)
	if destroy := event.Destroy; destroy != nil && *destroy == true {
		seat.Graveyard[event.Card.Id] = event.Card
		animate.GameSeat(game, seat)
	} else if destroy != nil {
		seat.Deck.Prepend(event.Card)
	}
	return event.Stack
}

func (event *NoviceSeerEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	log := game.Log().Add("Username", seat.Username)

	if seat.Username != game.State.Seat {
		log.Add("HotSeat", game.State.Seat).Warn("games.NoviceSeerEvent: not your choice")
		return
	}

	switch json.Sval("choice") {
	case "yes":
		b := true
		event.Destroy = &b
		break
	case "no":
		b := false
		event.Destroy = &b
		break
	default:
		log.Add("Choice", json.Sval("choice")).Warn("games.NoviceSeerEvent: unrecognized choice")
		return
	}
	log.Add("Destroy", event.Destroy).Info("games.NoviceSeer: confirmed destroy choice")

}

func (event *NoviceSeerEvent) Json(game *vii.Game) js.Object {
	return js.Object{
		"gameid": game.Key,
		"choice":   "Novice Seer",
		"username": game.State.Seat,
		"timer":    game.State.Timer.Seconds(),
	}
}
