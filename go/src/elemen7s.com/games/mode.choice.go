package games

import (
	"ztaylor.me/js"
)

type ChoiceMode func(js.Object)

func ChoiceModeFunc(f func(js.Object)) EMode {
	return ChoiceMode(f)
}

func (mode ChoiceMode) Name() string {
	return "choice"
}

func (mode ChoiceMode) Json(e *Event, g *Game, s *Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"username": s.Username,
		"timer":    int(e.Duration.Seconds()),
	}
}

func (mode ChoiceMode) OnActivate(*Event, *Game) {
}

func (mode ChoiceMode) OnSendCatchup(*Event, *Game, *Seat) {
}

func (mode ChoiceMode) OnResolve(*Event, *Game) {
}

func (mode ChoiceMode) OnReceive(e *Event, g *Game, s *Seat, json js.Object) {
	if s.Username != e.Username {
		g.Log().Add("Username", s.Username).Add("HotSeat", e.Username).Warn("games.ChoiceMode: not your choice")
		return
	}

	mode(json)
}
