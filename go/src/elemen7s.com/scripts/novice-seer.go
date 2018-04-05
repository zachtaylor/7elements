package scripts

import (
	"elemen7s.com/games"
	"ztaylor.me/js"
)

func init() {
	games.Scripts["novice-seer"] = NoviceSeer
}

type NoviceSeerMode struct {
	destroy bool
	*games.Card
	*games.Stack
}

func (mode NoviceSeerMode) Name() string {
	return "choice"
}

func (mode NoviceSeerMode) Json(e *games.Event, g *games.Game, s *games.Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"choice":   "Novice Seer",
		"username": s.Username,
		"timer":    int(e.Duration.Seconds()),
	}
}

func (mode NoviceSeerMode) OnActivate(e *games.Event, g *games.Game) {
	games.AnimateNoviceSeerChoice(g.GetSeat(e.Username), g, mode.Card)
}

func (mode NoviceSeerMode) OnSendCatchup(e *games.Event, g *games.Game, s *games.Seat) {
	if e.Username == s.Username {
		games.AnimateNoviceSeerChoice(s, g, mode.Card)
	}
}

func (mode NoviceSeerMode) OnResolve(e *games.Event, g *games.Game) {
	seat := g.GetSeat(mode.Card.Username)
	if mode.destroy {
		seat.Graveyard[mode.Card.Id] = mode.Card
	} else {
		seat.Deck.Prepend(mode.Card)
	}
	mode.Stack.OnResolve(e, g)
}

func (mode NoviceSeerMode) OnReceive(e *games.Event, g *games.Game, s *games.Seat, json js.Object) {
	if s.Username != e.Username {
		g.Log().Add("Username", s.Username).Add("HotSeat", e.Username).Warn("games.NoviceSeerMode: not your choice")
		return
	}

	switch json.Sval("choice") {
	case "yes":
		mode.destroy = true
		break
	case "no":
		break
	default:
		g.Log().Add("Username", s.Username).Add("Choice", json.Sval("choice")).Warn("games.NoviceSeerMode: unrecognized choice")
		return
	}
	g.Log().Add("Username", s.Username).Add("Destroy", mode.destroy).Info("games.NoviceSeer: confirmed destroy choice")

	g.TimelineJoin(nil)
}

func NoviceSeer(g *games.Game, s *games.Seat, target interface{}) {
	event := games.NewEvent(s.Username)
	event.EMode = NoviceSeerMode{
		Card:  s.Deck.Draw(),
		Stack: games.StackEvent(g.Active),
	}
	g.TimelineJoin(event)
}
