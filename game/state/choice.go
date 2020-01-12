package state

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

type Choice struct {
	R
	Text     string
	Data     cast.JSON
	Choices  []cast.JSON
	Finisher func(answer interface{})
	answer   interface{}
}

func NewChoice(seat, text string, data cast.JSON, choices []cast.JSON, fin func(answer interface{})) *Choice {
	return &Choice{
		R:        R(seat),
		Text:     text,
		Data:     data,
		Choices:  choices,
		Finisher: fin,
	}
}
func _choiceIsEvent(r *Choice) game.Stater {
	return r
}

func (r *Choice) Name() string {
	return "choice"
}

// OnActivate implements game.ActivateStater
func (r *Choice) OnActivate(g *game.T) []game.Stater {
	update.Choice(g.GetSeat(r.Seat()), r.Text, r.Choices, nil)
	return nil
}
func _activateStater(r *Choice) game.ActivateStater {
	return r
}

// OnConnect implements game.ConnectStater
func (r *Choice) OnConnect(g *game.T, s *game.Seat) {
	if s == nil || s.Username == r.Seat() {
		update.Choice(g.GetSeat(r.Seat()), r.Text, r.Choices, nil)
	}
}
func _choiceIsConnector(r *Choice) game.ConnectStater {
	return r
}

// Finish implements game.FinishStater
func (r *Choice) Finish(*game.T) []game.Stater {
	if r.Finisher != nil {
		r.Finisher(r.answer)
	}
	return nil
}
func _choiceIsFinisher(r *Choice) game.FinishStater {
	return r
}

func (r *Choice) GetNext(g *game.T) game.Stater {
	return nil
}

func (r *Choice) JSON() cast.JSON {
	return cast.JSON{
		"choice":  r.Text,
		"options": r.Choices,
		"data":    r.Data,
	}
}

// Request implements game.RequestStater
func (r *Choice) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	if seat.Username != r.Seat() {
		g.Log().With(cast.JSON{
			"Seat": seat,
			"json": json,
		}).Warn("choice: receive")
		return
	}

	r.answer = json["choice"]
	if r.answer != "" {
		for _, seat := range g.Seats {
			g.State.Reacts[seat.Username] = "push"
		}
	}
}
