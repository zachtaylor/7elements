package event

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func Play(seat string, card *game.Card, target interface{}) game.Event {
	return &PlayEvent{
		Event:  Event(seat),
		Card:   card,
		Target: target,
	}
}

type PlayEvent struct {
	Event
	Card        *game.Card
	Target      interface{}
	IsCancelled bool
}

func (event *PlayEvent) Name() string {
	return "play"
}

// OnActivate implements game.ActivateEventer
func (event *PlayEvent) OnActivate(g *game.T) []game.Event {
	msg := event.Card.Card.Name
	if event.Card.Card.Text != "" {
		msg = event.Card.Card.Text
	}
	go g.GetChat().AddMessage(chat.NewMessage(event.Seat(), msg))
	return nil
}
func (event *PlayEvent) activateEventer() game.ActivateEventer {
	return event
}

// // OnConnect implements game.ConnectEventer
// func (event *PlayEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Request implements game.RequestEventer
// func (event *PlayEvent) Request(*game.T, *game.Seat, cast.JSON) {
// }

// Finish implements game.FinishEventer
func (event *PlayEvent) Finish(g *game.T) []game.Event {
	seat := g.GetSeat(event.Seat())
	log := g.Log().With(log.Fields{
		"Seat": seat.Print(),
		"Card": event.Card.Print(),
	}).Tag("engine/play.finish")
	seat.Past[event.Card.Id] = event.Card

	if event.Card.Card.Type == vii.CTYPbody || event.Card.Card.Type == vii.CTYPitem {
		log.Debug("spawn")
		seat.Present[event.Card.Id] = event.Card // card in present and past
	}
	g.SendSeatUpdate(seat)

	powers := event.Card.Powers.GetTrigger("play")
	events := make([]game.Event, 0)
	for _, power := range powers {
		if script := game.Scripts[power.Script]; script == nil {
			if event.Card.Card.Type == vii.CTYPspell {
				log.Add("Script", power.Script).Error("no script")
			}
		} else if power.Target == "self" {
			if xtraevents := script(g, seat, event.Card); len(xtraevents) > 0 {
				events = append(events, xtraevents...)
			}
		} else if event.Target != nil {
			if xtraevents := script(g, seat, event.Target); len(xtraevents) > 0 {
				events = append(events, xtraevents...)
			}
		} else {
			events = append(events, NewTargetEvent(
				seat.Username,
				power.Target,
				power.Text,
				func(val string) []game.Event {
					return script(g, seat, val)
				},
			))
		}
	}

	return events
}
func (event *PlayEvent) finishEventer() game.FinishEventer {
	return event
}

func (event *PlayEvent) GetNext(g *game.T) game.Event {
	return nil
}

func (event *PlayEvent) JSON() cast.JSON {
	json := cast.JSON{
		"card": event.Card.JSON(),
	}
	if c, ok := event.Target.(*game.Card); ok {
		json["target"] = c.JSON()
	} else {
		json["target"] = event.Target
	}
	return json
}

func (event *PlayEvent) String() string {
	return event.Seat() + " played " + event.Card.Card.Name
}
