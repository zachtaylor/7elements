package state

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func NewSunrise(seat string) game.Stater {
	return &Sunrise{
		R: R(seat),
	}
}

type Sunrise struct {
	R
}

func (r *Sunrise) Name() string {
	return "sunrise"
}

// OnActivate implements game.ActivateStater
func (r *Sunrise) OnActivate(g *game.T) []game.Stater {
	events := make([]game.Stater, 0)
	g.GetSeat(r.Seat()).Elements.Reactivate()
	for _, seat := range g.Seats {
		update.Seat(g, seat)
		for _, token := range seat.Present {
			events = append(events, trigger.Wake(g, token)...)
		}
		events = append(events, trigger.Name(g, seat, "sunrise")...)
	}

	return events
}
func (r *Sunrise) _isActivateStater() game.ActivateStater {
	return r
}

// OnConnect implements game.ConnectStater
func (r *Sunrise) OnConnect(g *game.T, seat *game.Seat) {
	if seat == nil {
		update.Choice(g, "Create a New Element", update.ChoicesNewElement, nil)
		go g.GetChat().AddMessage(chat.NewMessage("sunrise", r.Seat()))
	} else if g.State.Reacts[seat.Username] == "" {
		update.Choice(seat, `Create a New Element`, update.ChoicesNewElement, nil)
	}
}
func (r *Sunrise) _isConnectStater() game.ConnectStater {
	return r
}

// // GetStack implements game.StackStater
// func (r *Sunrise) GetStack(g *game.T) *game.State {
// 	return nil
// }

// Finish implements game.FinishStater
func (r *Sunrise) Finish(g *game.T) []game.Stater {
	for _, seat := range g.Seats {
		log := g.Log().Add("Username", seat.Username)
		if card := seat.Deck.Draw(); card != nil {
			seat.Hand[card.ID] = card
			update.Hand(seat)
			log.Add("Hand", len(seat.Hand))
		}
		if react := g.State.Reacts[seat.Username]; react == "" {
			log.Source().Warn("react is empty")
		} else if el := cast.IntS(react); el < 1 || el > 7 {
			log.Add("React", react).Source().Warn("el is out of bounds")
			continue
		} else {
			seat.Elements.Append(vii.Element(el))
		}
		update.Seat(g, seat)
		log.Source().Debug()
	}
	return nil
}
func (r *Sunrise) _isFinishStater() game.FinishStater {
	return r
}

func (r *Sunrise) GetNext(g *game.T) game.Stater {
	return NewMain(r.Seat())
}

func (r *Sunrise) JSON() cast.JSON {
	return nil
}

// Request implements game.RequestStater
func (r *Sunrise) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	elementID := json.GetI("choice")
	log := g.Log().Add("Seat", seat).Add("Choice", elementID)
	if elementID < 1 || elementID > 7 {
		log.Source().Warn("elementid out of bounds")
		return
	}
	log.Source().Info("confirm")
	g.State.Reacts[seat.Username] = cast.StringI(elementID)
	update.React(g, seat.Username)
}
