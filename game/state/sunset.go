package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/cast"
)

func NewSunset(seat string) game.Stater {
	return &Sunset{
		R: R(seat),
	}
}

type Sunset struct {
	R
}

func (r *Sunset) Name() string {
	return "sunset"
}

// OnActivate implements game.ActivateStater
func (r *Sunset) OnActivate(g *game.T) []game.Stater {
	events := make([]game.Stater, 0)
	for _, seat := range g.Seats {
		events = append(events, trigger.AllPresent(g, seat, "sunset")...)
	}
	return events
}
func _sunsetIsActivateEventer(r *Sunset) game.ActivateStater {
	return r
}

// OnConnect implements game.ConnectStater
func (r *Sunset) OnConnect(g *game.T, seat *game.Seat) {
	if seat == nil {
		go g.Settings.Chat.AddMessage(chat.NewMessage("sunset", r.Seat()))
	}
}
func (r *Sunset) _isConnectStater() game.ConnectStater {
	return r
}

// // // Finish implements game.FinishStater
// func (r *Sunset) Finish(g *game.T) {
// 	// game.State.Seat = game.GetOpponentSeat(game.State.Seat).Username
// }

// // GetStack implements game.StackEventer
// func (r *Sunset) GetStack(g *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestStater
// func (r *Sunset) Request(g*game.T, seat *game.Seat, json cast.JSON) {
// }

func (r *Sunset) GetNext(g *game.T) game.Stater {
	return NewSunrise(g.GetOpponentSeat(r.Seat()).Username)
}

func (r *Sunset) JSON() cast.JSON {
	return nil
}
