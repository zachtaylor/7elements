package scripts

import (
	"github.com/zachtaylor/7tcg"
	"github.com/zachtaylor/7tcg/animate"
	"github.com/zachtaylor/7tcg/engine"
	"ztaylor.me/js"
)

const NoviceSeerId = "novice-seer"

func init() {
	engine.Scripts[NoviceSeerId] = NoviceSeer
}

func NoviceSeer(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	return t.Fork(game, &NoviceSeerEvent{
		Card:  seat.Deck.Draw(),
		Stack: t,
	})
}

type NoviceSeerEvent struct {
	Destroy *bool
	Card    *vii.GameCard
	Stack   *engine.Timeline
}

func (event *NoviceSeerEvent) Name() string {
	return "choice"
}

func (event *NoviceSeerEvent) Priority(game *vii.Game, t *engine.Timeline) bool {
	return event.Destroy == nil
}

func (event *NoviceSeerEvent) OnStart(game *vii.Game, t *engine.Timeline) {
	animate.NoviceSeerChoice(game.GetSeat(t.HotSeat), game, event.Card)
}

func (event *NoviceSeerEvent) OnReconnect(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat) {
	if t.HotSeat == seat.Username {
		animate.NoviceSeerChoice(seat, game, event.Card)
	}
}

func (event *NoviceSeerEvent) OnStop(game *vii.Game, t *engine.Timeline) *engine.Timeline {
	seat := game.GetSeat(event.Card.Username)
	if destroy := event.Destroy; destroy != nil && *destroy == true {
		seat.Graveyard[event.Card.Id] = event.Card
		animate.BroadcastSeatUpdate(game, seat)
	} else if destroy != nil {
		seat.Deck.Prepend(event.Card)
	}
	return event.Stack
}

func (event *NoviceSeerEvent) Receive(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, json js.Object) {
	log := game.Log().Add("Username", seat.Username)

	if seat.Username != t.HotSeat {
		log.Add("HotSeat", t.HotSeat).Warn("games.NoviceSeerEvent: not your choice")
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

func (event *NoviceSeerEvent) Json(game *vii.Game, t *engine.Timeline) js.Object {
	seat := game.GetSeat(t.HotSeat)
	return js.Object{
		"gameid":   game,
		"choice":   "Novice Seer",
		"username": seat.Username,
		"timer":    t.Lifetime.Seconds(),
	}
}
