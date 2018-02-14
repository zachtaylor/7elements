package games

import (
	"elemen7s.com/accounts"
	"elemen7s.com/decks"
	"ztaylor.me/js"
)

type EndMode bool

func (m EndMode) Name() string {
	return "end"
}

func (m EndMode) OnActivate(e *Event, g *Game) {
	log := g.Log()

	for _, username := range g.Results.Winners {
		if seat := g.GetSeat(username); seat == nil {
			log.Clone().Add("Username", username).Warn("games.End: winning seat missing")
		} else if err := decks.UpdateAddWin(username, seat.Deck.DeckId); err != nil {
			log.Clone().Add("Error", err).Add("Username", username).Error("games.End: winning deck missing")
		} else if err := accounts.UpdateAddCoins(username, 2); err != nil {
			log.Clone().Add("Error", err).Add("Username", username).Error("games.End: winning account missing")
		}
	}

	for _, username := range g.Results.Losers {
		if err := accounts.UpdateAddCoins(username, 1); err != nil {
			log.Clone().Add("Error", err).Add("Username", username).Error("games.End: loser account missing")
		}
	}
}

func (m EndMode) Json(e *Event, g *Game, s *Seat) js.Object {
	return js.Object{
		"gameid":  g.Id,
		"winners": g.Results.Winners,
		"losers":  g.Results.Losers,
	}
}

func (m EndMode) OnResolve(e *Event, g *Game) {
	g.Log().Debug("games.End: resolve")
	Cache.Remove(g)
}

func (m EndMode) OnReceive(e *Event, g *Game, seat *Seat, j js.Object) {
	g.Log().Add("Username", seat.Username).Add("Name", j["event"]).Warn("games.End: receive")
}

func End(g *Game) {
	e := NewEvent("end")
	e.EMode = EndMode(true)
	go func() {
		g.Timeline <- e
	}()
}
