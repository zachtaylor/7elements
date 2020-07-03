package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/out"
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

func (r *Sunrise) sendChoice(seat *game.Seat) {
	out.Choice(seat.Player, `Create a New Element`, out.ChoicesElements, nil)
}

// OnActivate implements game.ActivateStater
func (r *Sunrise) OnActivate(g *game.T) []game.Stater {
	events := make([]game.Stater, 0)
	g.GetSeat(r.Seat()).Karma.Reactivate()
	for _, seat := range g.Seats {
		out.GameSeat(g, seat.JSON())
		for _, token := range seat.Present {
			events = append(events, trigger.Wake(g, token)...)
		}
		events = append(events, trigger.AllPresent(g, seat, "sunrise")...)
	}

	return events
}
func (r *Sunrise) _isActivateStater() game.ActivateStater {
	return r
}

// OnConnect implements game.ConnectStater
func (r *Sunrise) OnConnect(g *game.T, seat *game.Seat) {
	if seat == nil {
		out.Choice(g, `Create a New Element`, out.ChoicesElements, nil)
		go g.Settings.Chat.AddMessage(chat.NewMessage("sunrise", r.Seat()))
	} else if g.State.Reacts[seat.Username] == "" {
		out.Choice(seat.Player, `Create a New Element`, out.ChoicesElements, nil)
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
			out.GameHand(seat.Player, seat.Hand.JSON())
			log.Add("Hand", len(seat.Hand))
		}
		if react := g.State.Reacts[seat.Username]; react == "" {
			log.Warn("react is empty")
		} else if el := cast.IntS(react); el < 1 || el > 7 {
			log.Add("React", react).Warn("el is out of bounds")
			continue
		} else {
			seat.Karma.Append(element.T(el), false)
		}
		out.GameSeat(g, seat.JSON())
		log.Debug()
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
		log.Warn("elementid out of bounds")
		return
	}
	log.Info("confirm")
	g.State.Reacts[seat.Username] = cast.StringI(elementID)
	out.GameReact(g, g.State.ID(), seat.Username, "confirm", g.State.Timer)
}
