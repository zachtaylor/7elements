package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func NewTarget(seat, helper, text string, finish func(val string) []game.Stater) *Target {
	return &Target{
		R:        R(seat),
		Helper:   helper,
		Text:     text,
		Finisher: finish,
	}
}

type Target struct {
	R
	Helper   string
	Text     string
	Finisher func(val string) []game.Stater
	answer   string
}

func (r *Target) Name() string {
	return "target"
}

func (r *Target) JSON() cast.JSON {
	return cast.JSON{
		"helper": r.Helper,
		"text":   r.Text,
	}
}

func (r *Target) GetNext(g *game.T) game.Stater {
	return nil
}

// OnActivate implements game.ActivateStater
func (r *Target) OnActivate(g *game.T) []game.Stater {
	go g.Settings.Chat.AddMessage(chat.NewMessage(r.Seat(), r.Text))
	return nil
}
func (r *Target) activateEventer() game.ActivateStater {
	return r
}

// OnConnect implements game.ConnectStater
func (r *Target) OnConnect(g *game.T, seat *game.Seat) {
	if seat == nil {
		go g.Settings.Chat.AddMessage(chat.NewMessage("target", r.Seat()))
	}
}
func (r *Target) _isConnectStater() game.ConnectStater {
	return r
}

// Request implements game.RequestStater
func (r *Target) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	if seat.Username != r.Seat() {
		g.Log().With(cast.JSON{
			"seat": seat,
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
