package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func End(game *game.T, winner, loser string) game.Event {
	return &EndEvent{
		Winner: winner,
		Loser:  loser,
	}
}

type EndEvent game.Results

func (event *EndEvent) Name() string {
	return "end"
}

// // OnActivate implements game.ActivateEventer
// func (event *EndEvent) OnActivate(game *game.T) {
// }

// // // OnConnect implements game.ConnectEventer
// func (event *EndEvent) OnConnect(*game.T, *game.Seat) {
// }

// // GetStack implements game.StackEventer
// func (event *EndEvent) GetStack(game *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestEventer
// func (event *EndEvent) Request(game *game.T, seat *game.Seat, json vii.Json) {
// }

func (event *EndEvent) GetNext(game *game.T) *game.State {
	return nil
}

// Finish implements game.FinishEventer
func (event *EndEvent) Finish(game *game.T) {
	log := game.Logger.WithFields(log.Fields{
		"Winner": event.Winner,
		"Loser":  event.Loser,
	})

	if username := event.Winner; username == "" {
		log.Warn("engine/end: winner missing")
	} else if username == "A.I" {
		// skip
	} else if seat := game.GetSeat(username); seat == nil {
		log.Warn("engine/end: winning seat missing")
	} else if err := vii.AccountDeckService.UpdateTallyWinCount(username, seat.Deck.AccountDeckID); err != nil {
		log.Clone().Add("Error", err).Error("engine/end: winning deck missing")
	} else if account, err := vii.AccountService.Get(username); err != nil {
		log.Clone().Add("Error", err).Error("engine/end: account missing")
	} else {
		account.Coins += 2
		if err = vii.AccountService.UpdateCoins(account); err != nil {
			log.Clone().Add("Error", err).Add("Username", username).Error("engine/end: account service error")
		}
	}

	if username := event.Loser; username == "" {
		log.Warn("engine/end: loser missing")
	} else if username == "A.I." {
		// skip
	} else if account, err := vii.AccountService.Get(username); err != nil {
		log.Add("Error", err).Error("engine/end: account missing")
	} else if event.Winner == "" {
		log.Add("Error", "Forfeit!").Warn("engine/end: no pity coins")
	} else {
		account.Coins++
		if err = vii.AccountService.UpdateCoins(account); err != nil {
			log.Add("Error", err).Error("engine/end: account service error")
		}
	}
}

func (event *EndEvent) Json(game *game.T) vii.Json {
	return js.Object{
		"winner": event.Winner,
		"loser":  event.Loser,
	}
}
