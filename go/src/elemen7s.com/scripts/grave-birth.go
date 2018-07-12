package scripts

import (
	"elemen7s.com"
	"elemen7s.com/animate"
	"elemen7s.com/engine"
	"ztaylor.me/js"
)

const GraveBirthID = "grave-birth"

func init() {
	engine.Scripts[GraveBirthID] = GraveBirth
}

func GraveBirth(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	return t.Fork(game, &GraveBirthEvent{
		Stack: t,
	})
}

type GraveBirthEvent struct {
	Card  *vii.GameCard
	Stack *engine.Timeline
}

func (event *GraveBirthEvent) Name() string {
	return "choice"
}

func (event *GraveBirthEvent) Priority(game *vii.Game, t *engine.Timeline) bool {
	return event.Card == nil
}

func (event *GraveBirthEvent) OnStart(game *vii.Game, t *engine.Timeline) {
	animate.GraveBirth(game.GetSeat(t.HotSeat), game)
}

func (event *GraveBirthEvent) OnReconnect(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat) {
	if t.HotSeat == seat.Username {
		animate.GraveBirth(seat, game)
	}
}

func (event *GraveBirthEvent) OnStop(game *vii.Game, t *engine.Timeline) *engine.Timeline {
	seat := game.GetSeat(t.HotSeat)
	card := vii.NewGameCard(event.Card.Card, event.Card.CardText)
	card.Username = seat.Username
	card.IsToken = true
	game.RegisterCard(card)
	animate.BroadcastSeatUpdate(game, seat)
	return event.Stack
}

func (event *GraveBirthEvent) Receive(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, json js.Object) {
	log := game.Log().Add("Username", seat.Username).Add("Choice", json.Val("choice"))

	if seat.Username != t.HotSeat {
		log.Warn(GraveBirthID + ": not your choice")
		return
	}

	if gcid := json.Sval("choice"); gcid == "" {
		log.Warn(GraveBirthID + ": choice not found")
	} else if card := game.Cards[gcid]; card == nil {
		log.Warn(GraveBirthID + ": gcid not found")
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Warn(GraveBirthID + ": not type body")
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("CardOwner", card.Username).Warn(GraveBirthID + ": card owner not found")
	} else if !ownerSeat.HasPastCard(gcid) {
		log.Add("CardOwner", card.Username).Add("Past", ownerSeat.Graveyard.String()).Warn(GraveBirthID + ": card not in past")
	} else {
		event.Card = vii.NewGameCard(card.Card, card.CardText)
		log.Add("CardId", event.Card.Card.Id).Info(GraveBirthID + ": confirmed card")
	}
}

func (event *GraveBirthEvent) Json(game *vii.Game, t *engine.Timeline) js.Object {
	seat := game.GetSeat(t.HotSeat)
	return js.Object{
		"gameid":   game,
		"choice":   "Grave Birth",
		"username": seat.Username,
		"timer":    t.Lifetime.Seconds(),
	}
}
