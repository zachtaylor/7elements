package end

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func New(winner, loser string) game.Stater {
	return &EndStater{
		Winner: winner,
		Loser:  loser,
	}
}

type EndStater game.Results

func (event *EndStater) Name() string {
	return "end"
}

func (event *EndStater) Seat() string {
	return ""
}

// OnActivate implements game.ActivateStaterer
func (event *EndStater) OnActivate(g *game.T) []game.Stater {
	if g.State.R == event {
		g.State.Stack = nil // rip L0L
	}
	return nil
}
func _endActivateStaterer(event *EndStater) game.ActivateStater {
	return event
}

// // // OnConnect implements game.ConnectStater
// func (event *EndStater) OnConnect(*game.T, *game.Seat) {
// }

func (event *EndStater) OnDisconnect(g *game.T, seat *game.Seat) {
	g.Log().Add("Username", seat.Username).Debug("left")
	g.State.Reacts[seat.Username] = "left"
}

// // GetStack implements game.StackStater
// func (event *EndStater) GetStack(g *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestStater
// func (event *EndStater) Request(g*game.T, seat *game.Seat, json cast.JSON) {
// }

func (event *EndStater) GetNext(g *game.T) game.Stater {
	return nil
}

// Finish implements game.FinishStaterer
func (event *EndStater) Finish(g *game.T) []game.Stater {
	log := g.Log().With(cast.JSON{
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
		} else if rcv := seat.Receiver; rcv == nil {
			log.Copy().Add("Error", err).Add("Username", username).Error("no receiver")
		} else {
			g.Runtime.Root.SendAccountUpdate(seat.WriteJSON, account.Username)
		}
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
		g.Runtime.Root.SendAccountUpdate(seat.WriteJSON, account.Username)
	}
	g.Close()
	return nil
}
func (event *EndStater) finishStaterer() game.FinishStater {
	return event
}

func (event *EndStater) JSON() cast.JSON {
	return cast.JSON{
		"winner": event.Winner,
		"loser":  event.Loser,
	}
}
