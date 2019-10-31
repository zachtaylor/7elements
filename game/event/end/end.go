package end

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func New(winner, loser string) game.Event {
	return &EndEvent{
		Winner: winner,
		Loser:  loser,
	}
}

type EndEvent game.Results

func (event *EndEvent) Name() string {
	return "end"
}

func (event *EndEvent) Seat() string {
	return ""
}

// OnActivate implements game.ActivateEventer
func (event *EndEvent) OnActivate(g *game.T) []game.Event {
	if g.State.Event == event {
		g.State.Stack = nil // rip
	}
	return nil
}
func _endActivateEventer(event *EndEvent) game.ActivateEventer {
	return event
}

// // // OnConnect implements game.ConnectEventer
// func (event *EndEvent) OnConnect(*game.T, *game.Seat) {
// }

func (event *EndEvent) OnDisconnect(g *game.T, seat *game.Seat) {
	g.Log().Add("Username", seat.Username).Debug("left")
	g.State.Reacts[seat.Username] = "left"
}

// // GetStack implements game.StackEventer
// func (event *EndEvent) GetStack(g *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestEventer
// func (event *EndEvent) Request(g*game.T, seat *game.Seat, json cast.JSON) {
// }

func (event *EndEvent) GetNext(g *game.T) game.Event {
	return nil
}

// Finish implements game.FinishEventer
func (event *EndEvent) Finish(g *game.T) []game.Event {
	log := g.Log().With(log.Fields{
		"Winner": event.Winner,
		"Loser":  event.Loser,
	}).Tag("end")

	if username := event.Winner; username == "" {
		log.Warn("winner missing")
	} else if username == "A.I" {
		// skip
	} else if seat := g.GetSeat(username); seat == nil {
		log.Warn("winning seat missing")
	} else if account, err := g.Runtime.Root.Accounts.Get(username); err != nil {
		log.Copy().Add("Error", err).Error("account missing")
	} else {
		account.Coins += 2
		if err = g.Runtime.Root.Accounts.UpdateCoins(account); err != nil {
			log.Copy().Add("Error", err).Add("Username", username).Error("account service error")
		}
		seat.Send(game.BuildPushJSON("/data/myaccount", g.Runtime.Root.AccountJSON(account.Username)))
	}

	if username := event.Loser; username == "" {
		log.Warn("loser missing")
	} else if username == "A.I." {
		// skip
	} else if seat := g.GetSeat(username); seat == nil {
		log.Warn("winning seat missing")
	} else if account, err := g.Runtime.Root.Accounts.Get(username); err != nil {
		log.Add("Error", err).Error("account missing")
	} else if event.Winner == "" {
		log.Add("Error", "Forfeit!").Warn("no pity coins")
	} else {
		account.Coins++
		if err = g.Runtime.Root.Accounts.UpdateCoins(account); err != nil {
			log.Add("Error", err).Error("account service error")
		}
		seat.Send(game.BuildPushJSON("/data/myaccount", g.Runtime.Root.AccountJSON(account.Username)))
	}
	g.Close()
	return nil
}
func (event *EndEvent) finishEventer() game.FinishEventer {
	return event
}

func (event *EndEvent) JSON() cast.JSON {
	return cast.JSON{
		"winner": event.Winner,
		"loser":  event.Loser,
	}
}
