package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
	"ztaylor.me/js"
)

const GraveBirthID = "grave-birth"

func init() {
	engine.Scripts[GraveBirthID] = GraveBirth
}

func GraveBirth(game *game.T, seat *game.Seat, target interface{}) game.Event {
	return &GraveBirthEvent{
		Stack: game.State,
	}
}

type GraveBirthEvent struct {
	Stack *game.State
	Card  *game.Card
}

func (event *GraveBirthEvent) Name() string {
	return "choice"
}

// OnActivate implements game.ActivateEventer
func (event *GraveBirthEvent) OnActivate(game *game.T) {
	seat := game.GetSeat(game.State.Seat)
	animate.GraveBirth(seat, game)
}

// OnConnect implements game.ConnectEventer
func (event *GraveBirthEvent) OnConnect(game *game.T, seat *game.Seat) {
	if game.State.Seat == seat.Username {
		animate.GraveBirth(seat, game)
	}
}

// Finish implements game.FinishEventer
func (event *GraveBirthEvent) Finish(g *game.T) {
	card := game.NewCard(event.Card.Card)
	card.Username = g.State.Seat
	card.IsToken = true
	g.RegisterCard(card)
	animate.GameSeat(g, g.GetSeat(g.State.Seat))
}

// GetStack implements game.StackEventer
func (event *GraveBirthEvent) GetStack(g *game.T) *game.State {
	return event.Stack
}

// GetNext implements game.StackEventer
func (event *GraveBirthEvent) GetNext(g *game.T) *game.State {
	return nil
}

func (event *GraveBirthEvent) Json(game *game.T) js.Object {
	return js.Object{
		"choice": "Grave Birth",
	}
}

func (event *GraveBirthEvent) Request(g *game.T, seat *game.Seat, json js.Object) {
	log := g.Log().Add("Username", seat.Username).Add("Choice", json.Val("choice"))

	if seat.Username != g.State.Seat {
		log.Warn(GraveBirthID + ": not your choice")
		return
	}

	if gcid := json.Sval("choice"); gcid == "" {
		log.Warn(GraveBirthID + ": choice not found")
	} else if card := g.Cards[gcid]; card == nil {
		log.Warn(GraveBirthID + ": gcid not found")
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Warn(GraveBirthID + ": not type body")
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Add("CardOwner", card.Username).Warn(GraveBirthID + ": card owner not found")
	} else if !ownerSeat.HasPastCard(gcid) {
		log.Add("CardOwner", card.Username).Add("Past", ownerSeat.Past.String()).Warn(GraveBirthID + ": card not in past")
	} else {
		log.Add("CardId", event.Card.Card.Id).Info(GraveBirthID + ": confirmed card")
		event.Card = game.NewCard(card.Card)
		g.State.Reacts[seat.Username] = "confirm"
	}
}
