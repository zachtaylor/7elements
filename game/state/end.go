package state

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func NewEnd(winner, loser string) game.Stater {
	return &End{
		Winner: winner,
		Loser:  loser,
	}
}

type End game.Results

func (r *End) Name() string {
	return "end"
}

func (r *End) Seat() string {
	return ""
}

// OnActivate implements game.ActivateStaterer
func (r *End) OnActivate(g *game.T) []game.Stater {
	if g.State.Name() == "end" {
		g.State.Stack = nil // rip L0L
	}
	return nil
}
func _endActivateStaterer(r *End) game.ActivateStater {
	return r
}

// // // OnConnect implements game.ConnectStater
// func (r *End) OnConnect(*game.T, *game.Seat) {
// }

func (r *End) OnDisconnect(g *game.T, seat *game.Seat) {
	g.Log().Add("Username", seat.Username).Debug("left")
	g.State.Reacts[seat.Username] = "left"
}

// // GetStack implements game.StackStater
// func (r *End) GetStack(g *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestStater
// func (r *End) Request(g*game.T, seat *game.Seat, json cast.JSON) {
// }

func (r *End) GetNext(g *game.T) game.Stater {
	return nil
}

// Finish implements game.FinishStaterer
func (r *End) Finish(g *game.T) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Winner": r.Winner,
		"Loser":  r.Loser,
	}).Tag("end")

	if username := r.Winner; username == "" {
		log.Warn("winner missing")
	} else if username == "A.I" {
		// skip
	} else if seat := g.GetSeat(username); seat == nil {
		log.Warn("winning seat missing")
	} else if account, err := g.Settings.Accounts.Get(username); err != nil {
		log.Copy().Add("Error", err).Error("account missing")
	} else {
		account.Coins += 2
		if err = g.Settings.Accounts.UpdateCoins(account); err != nil {
			log.Copy().Add("Error", err).Add("Username", username).Error("account service error")
		} else {
			out.Account(seat.Player)
		}
	}

	if username := r.Loser; username == "" {
		log.Warn("loser missing")
	} else if username == "A.I." {
		// skip
	} else if seat := g.GetSeat(username); seat == nil {
		log.Warn("winning seat missing")
	} else if account, err := g.Settings.Accounts.Get(username); err != nil {
		log.Add("Error", err).Error("account missing")
	} else if r.Winner == "" {
		log.Add("Error", "Forfeit!").Warn("no pity coins")
	} else {
		account.Coins++
		if err = g.Settings.Accounts.UpdateCoins(account); err != nil {
			log.Add("Error", err).Error("account service error")
		}
		out.Account(seat.Player)
	}
	g.Close()
	return nil
}
func (r *End) finishStater() game.FinishStater {
	return r
}

func (r *End) JSON() cast.JSON {
	return cast.JSON{
		"winner": r.Winner,
		"loser":  r.Loser,
	}
}
