package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func NewStart(seat string) game.Stater {
	return &Start{
		R: R(seat),
	}
}

type Start struct {
	R
}

func (r *Start) Name() string {
	return "start"
}

// OnActivate implements game.ActivateStater
func (r *Start) OnActivate(g *game.T) []game.Stater {
	for _, seat := range g.Seats {
		seat.Life = 7
		seat.Deck.Shuffle()
		seat.DrawCard(3)
		update.Hand(seat)
	}
	return nil
}
func _startIsActivator(r *Start) game.ActivateStater {
	return r
}

// OnConnect implements game.ConnectStater
func (r *Start) OnConnect(g *game.T, seat *game.Seat) {
	if seat == nil {
		go g.GetChat().AddMessage(chat.NewMessage("sunrise", r.Seat()))
	} else if g.State.Reacts[seat.Username] == "" {
		update.Choice(seat, "Choose Starting Hand", []cast.JSON{
			cast.JSON{"choice": "keep", "display": `<button class="vii">Click here to Keep</button>`},
			cast.JSON{"choice": "mulligan", "display": `<button class="vii-alt">Click here to Mulligan</button>`},
		}, nil)
	}
}
func (r *Start) _isConnectStater() game.ConnectStater {
	return r
}

// // Finish implements game.FinishStater
// func (r *Start) Finish(*game.T) {
// }

// // GetStack implements game.StackEventer
// func (r *Start) GetStack(g *game.T) *game.State {
// 	return nil
// }

func (r *Start) GetNext(g *game.T) game.Stater {
	return NewSunrise(r.Seat())
}

func (r *Start) JSON() cast.JSON {
	return nil
}

// Request implements game.RequestStater
func (r *Start) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	choice := json.GetS("choice")
	log := g.Log().Add("Seat", seat).Add("Choice", choice).Tag("engine/start")

	if react := g.State.Reacts[seat.Username]; react != "" {
		log.Add("React", react).Warn("already recorded")
		return
	} else if choice == "keep" {
		g.State.Reacts[seat.Username] = "keep"
	} else if choice == "mulligan" {
		g.State.Reacts[seat.Username] = "mulligan"
		seat.DiscardHand()
		seat.DrawCard(3)
		update.Hand(seat)
		update.Seat(g, seat)
	} else {
		log.Warn("unrecognized")
		return
	}

	g.State.Reacts[seat.Username] = choice
	update.React(g, seat.Username)
	log.Info("confirm")
}
