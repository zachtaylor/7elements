package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func NewTarget(seat, helper, display string, finish func(val string) []game.Stater) *Target {
	return &Target{
		R:        R(seat),
		Helper:   helper,
		Display:  display,
		Finisher: finish,
	}
}

type Target struct {
	R
	Helper   string
	Display  string
	Finisher func(val string) []game.Stater
	answer   string
}

func (r *Target) Name() string {
	return "target"
}

func (r *Target) JSON() cast.JSON {
	return cast.JSON{
		"helper":  r.Helper,
		"display": r.Display,
	}
}

func (r *Target) GetNext(g *game.T) game.Stater {
	return nil
}

// OnActivate implements game.ActivateStater
func (r *Target) OnActivate(g *game.T) []game.Stater {
	go g.GetChat().AddMessage(chat.NewMessage(r.Seat(), r.Display))
	return nil
}
func (r *Target) activateEventer() game.ActivateStater {
	return r
}

// Request implements game.RequestStater
func (r *Target) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	if seat.Username != r.Seat() {
		g.Log().With(cast.JSON{
			"Seat": seat,
			"json": json,
		}).Warn("engine/target: receive")
		return
	}

	r.answer = json.GetS("choice")
	if r.answer != "" {
		for _, seat := range g.Seats {
			g.State.Reacts[seat.Username] = "push"
		}
	}
}
func _targetIsRequester(r *Target) game.RequestStater {
	return r
}

// Finish implements game.FinishStater
func (r *Target) Finish(*game.T) []game.Stater {
	if r.Finisher != nil {
		r.Finisher(r.answer)
	}
	return nil
}
func _targetIsFinisher(r *Target) game.FinishStater {
	return r
}
