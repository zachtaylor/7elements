package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
	"ztaylor.me/js"
)

const GraveBirthID = "grave-birth"

func init() {
	engine.Scripts[GraveBirthID] = GraveBirth
}

func GraveBirth(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	return &GraveBirthEvent{
		Stack: game.State.Event,
	}
}

type GraveBirthEvent struct {
	Card  *vii.GameCard
	Stack vii.GameEvent
}

func (event *GraveBirthEvent) Name() string {
	return "choice"
}

func (event *GraveBirthEvent) Priority(game *vii.Game) bool {
	return event.Card == nil
}

func (event *GraveBirthEvent) OnStart(game *vii.Game) {
	seat := game.GetSeat(game.State.Seat)
	animate.GraveBirth(seat, game)
}

func (event *GraveBirthEvent) OnReconnect(game *vii.Game, seat *vii.GameSeat) {
	if game.State.Seat == seat.Username {
		animate.GraveBirth(seat, game)
	}
}

func (event *GraveBirthEvent) NextEvent(game *vii.Game) vii.GameEvent {
	card := vii.NewGameCard(event.Card.Card)
	card.Username = game.State.Seat
	card.IsToken = true
	game.RegisterCard(card)
	animate.GameSeat(game, game.GetSeat(game.State.Seat))
	return event.Stack
}

func (event *GraveBirthEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	log := game.Log().Add("Username", seat.Username).Add("Choice", json.Val("choice"))

	if seat.Username != game.State.Seat {
		log.Warn(GraveBirthID + ": not your choice")
		return
	}

	if gcid := json.Sval("choice"); gcid == "" {
		log.Warn(GraveBirthID + ": choice not found")
	} else if card := game.Cards[gcid]; card == nil {
		log.Warn(GraveBirthID + ": gcid not found")
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Warn(GraveBirthID + ": not type body")
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("CardOwner", card.Username).Warn(GraveBirthID + ": card owner not found")
	} else if !ownerSeat.HasPastCard(gcid) {
		log.Add("CardOwner", card.Username).Add("Past", ownerSeat.Graveyard.String()).Warn(GraveBirthID + ": card not in past")
	} else {
		event.Card = vii.NewGameCard(card.Card)
		log.Add("CardId", event.Card.Card.Id).Info(GraveBirthID + ": confirmed card")
	}
}

func (event *GraveBirthEvent) Json(game *vii.Game) js.Object {
	return js.Object{
		"gameid":   game.Key,
		"choice":   "Grave Birth",
		"username": game.State.Seat,
		"timer":    game.State.Timer.Seconds(),
	}
}
