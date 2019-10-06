package event

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

type ChoiceEvent struct {
	Event
	Text     string
	Data     cast.JSON
	Choices  []cast.JSON
	Finisher func(answer interface{})
	answer   interface{}
}

func NewChoiceEvent(seat, text string, data cast.JSON, choices []cast.JSON, fin func(answer interface{})) *ChoiceEvent {
	return &ChoiceEvent{
		Event:    Event(seat),
		Text:     text,
		Data:     data,
		Choices:  choices,
		Finisher: fin,
	}
}
func _choiceIsEvent(event *ChoiceEvent) game.Event {
	return event
}

func (event *ChoiceEvent) Name() string {
	return "choice"
}

// OnActivate implements game.ActivateEventer
func (event *ChoiceEvent) OnActivate(g *game.T) []game.Event {
	s := g.GetSeat(event.Seat())
	s.Send(game.BuildChoiceUpdate(event.Text, event.Choices, nil))
	return nil
}
func _choiceIsActivator(event *ChoiceEvent) game.ActivateEventer {
	return event
}

// OnConnect implements game.ConnectEventer
func (event *ChoiceEvent) OnConnect(g *game.T, s *game.Seat) {
	if s == nil {
		s := g.GetSeat(event.Seat())
		s.Send(game.BuildChoiceUpdate(event.Text, event.Choices, nil))
	} else if event.Seat() == s.Username {
		s.Send(game.BuildChoiceUpdate(event.Text, event.Choices, nil))
	}
}
func _choiceIsConnector(event *ChoiceEvent) game.ConnectEventer {
	return event
}

// Finish implements game.FinishEventer
func (event *ChoiceEvent) Finish(*game.T) []game.Event {
	if event.Finisher != nil {
		event.Finisher(event.answer)
	}
	return nil
}
func _choiceIsFinisher(event *ChoiceEvent) game.FinishEventer {
	return event
}

func (event *ChoiceEvent) GetNext(g *game.T) game.Event {
	return nil
}

func (event *ChoiceEvent) JSON() cast.JSON {
	return cast.JSON{
		"choice":  event.Text,
		"options": event.Choices,
		"data":    event.Data,
	}
}

// Request implements game.RequestEventer
func (event *ChoiceEvent) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	if seat.Username != event.Seat() {
		g.Log().With(log.Fields{
			"Seat": seat,
			"json": json,
		}).Warn("choice: receive")
		return
	}

	event.answer = json["choice"]
	if event.answer != "" {
		for _, seat := range g.Seats {
			g.State.Reacts[seat.Username] = "push"
		}
	}
}
