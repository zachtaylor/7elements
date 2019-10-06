package event

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

type TargetEvent struct {
	Event
	Helper   string
	Display  string
	Finisher func(val string) []game.Event
	answer   string
}

func NewTargetEvent(seat, helper, display string, finish func(val string) []game.Event) *TargetEvent {
	return &TargetEvent{
		Event:    Event(seat),
		Helper:   helper,
		Display:  display,
		Finisher: finish,
	}
}
func _targetIsEvent(event *TargetEvent) game.Event {
	return event
}

func (event *TargetEvent) Name() string {
	return "target"
}

func (event *TargetEvent) JSON() cast.JSON {
	return cast.JSON{
		"helper":  event.Helper,
		"display": event.Display,
	}
}

func (event *TargetEvent) GetNext(g *game.T) game.Event {
	return nil
}

// OnActivate implements game.ActivateEventer
func (event *TargetEvent) OnActivate(g *game.T) []game.Event {
	go g.GetChat().AddMessage(chat.NewMessage(event.Seat(), event.Display))
	return nil
}
func (event *TargetEvent) activateEventer() game.ActivateEventer {
	return event
}

// Request implements game.RequestEventer
func (event *TargetEvent) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	if seat.Username != event.Seat() {
		g.Log().With(log.Fields{
			"Seat": seat,
			"json": json,
		}).Warn("engine/target: receive")
		return
	}

	event.answer = json.GetS("choice")
	if event.answer != "" {
		for _, seat := range g.Seats {
			g.State.Reacts[seat.Username] = "push"
		}
	}
}
func _targetIsRequester(event *TargetEvent) game.RequestEventer {
	return event
}

// Finish implements game.FinishEventer
func (event *TargetEvent) Finish(*game.T) []game.Event {
	if event.Finisher != nil {
		event.Finisher(event.answer)
	}
	return nil
}
func _targetIsFinisher(event *TargetEvent) game.FinishEventer {
	return event
}
