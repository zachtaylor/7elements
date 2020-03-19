package state

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func NewPlay(seat string, card *card.T, target interface{}) game.Stater {
	return &Play{
		R:      R(seat),
		Card:   card,
		Target: target,
	}
}

type Play struct {
	R
	Card        *card.T
	Target      interface{}
	IsCancelled bool
}

func (r *Play) Name() string {
	return "play"
}

// OnActivate implements game.ActivateStater
func (r *Play) OnActivate(g *game.T) []game.Stater {
	msg := r.Card.Proto.Name
	if r.Card.Proto.Text != "" {
		msg = r.Card.Proto.Text
	}
	go g.GetChat().AddMessage(chat.NewMessage(r.Seat(), msg))
	return nil
}
func (r *Play) activateEventer() game.ActivateStater {
	return r
}

// // OnConnect implements game.ConnectStater
// func (r *Play) OnConnect(*game.T, *game.Seat) {
// }

// // Request implements game.RequestStater
// func (r *Play) Request(*game.T, *game.Seat, cast.JSON) {
// }

// Finish implements game.FinishStater
func (r *Play) Finish(g *game.T) []game.Stater {
	seat := g.GetSeat(r.Seat())
	g.Log().With(cast.JSON{
		"Seat": seat.String(),
		"Card": r.Card.String(),
	}).Debug("engine/play: finish")
	seat.Past[r.Card.ID] = r.Card

	if r.Card.Proto.Type == card.BodyType || r.Card.Proto.Type == card.ItemType {
		trigger.Spawn(g, seat, r.Card)
	}
	update.Seat(g, seat)

	powers := r.Card.Proto.Powers.GetTrigger("play")
	events := make([]game.Stater, 0)
	for _, power := range powers {
		// trigger.Power(g, seat, power, r.Card, )

		if power.Target == "self" {
			if e := trigger.Power(g, seat, power, r.Card, cast.NewArray(r.Card)); e != nil {
				events = append(events, e...)
			}
		} else if r.Target != nil {
			if e := trigger.Power(g, seat, power, r.Card, cast.NewArray(r.Target)); e != nil {
				events = append(events, e...)
			}
		} else {
			events = append(events, NewTarget(
				seat.Username,
				power.Target,
				power.Text,
				func(val string) []game.Stater {
					return trigger.Power(g, seat, power, r.Card, cast.NewArray(val))
				},
			))
		}
	}

	return events
}
func (r *Play) FinishStater() game.FinishStater {
	return r
}

func (r *Play) GetNext(g *game.T) game.Stater {
	return nil
}

func (r *Play) JSON() cast.JSON {
	json := cast.JSON{
		"card": r.Card.JSON(),
	}
	if r.Target == nil {
		json["target"] = cast.Stringer(`null`)
	} else {
		json["target"] = r.Target
	}
	return json
}

func (r *Play) String() string {
	return r.Seat() + " played " + r.Card.Proto.Name
}
