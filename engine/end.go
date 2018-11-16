package engine

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func End(game *vii.Game, winner, loser string) vii.GameEvent {
	return &EndEvent{
		Winner: winner,
		Loser:  loser,
	}
}

type EndEvent vii.GameResults

func (event *EndEvent) Name() string {
	return "end"
}

func (event *EndEvent) Priority(game *vii.Game) bool {
	return len(game.State.Reacts) < len(game.Seats)
}

func (event *EndEvent) OnStart(game *vii.Game) {
	log := game.Logger.WithFields(log.Fields{
		"Winner": event.Winner,
		"Loser":  event.Loser,
	})

	if username := event.Winner; username == "" {
		log.Warn("engine-end: winner missing")
	} else if username == "A.I" {
		// skip
	} else if seat := game.GetSeat(username); seat == nil {
		log.Warn("engine-end: winning seat missing")
	} else if err := vii.AccountDeckService.UpdateTallyWinCount(username, seat.Deck.AccountDeckID); err != nil {
		log.Clone().Add("Error", err).Error("engine-end: winning deck missing")
	} else if account, err := vii.AccountService.Get(username); err != nil {
		log.Clone().Add("Error", err).Error("engine-end: account missing")
	} else {
		account.Coins += 2
		if err = vii.AccountService.UpdateCoins(account); err != nil {
			log.Clone().Add("Error", err).Add("Username", username).Error("engine-end: account service error")
		}
	}

	if username := event.Loser; username == "" {
		log.Warn("engine-end: loser missing")
	} else if username == "A.I." {
		// skip
	} else if account, err := vii.AccountService.Get(username); err != nil {
		log.Add("Error", err).Error("engine-end: account missing")
	} else if event.Winner == "" {
		log.Add("Error", "Forfeit!").Warn("engine-end: no pity coins")
	} else {
		account.Coins++
		if err = vii.AccountService.UpdateCoins(account); err != nil {
			log.Add("Error", err).Error("engine-end: account service error")
		}
	}
}

func (event *EndEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event *EndEvent) NextEvent(game *vii.Game) vii.GameEvent {
	game.Log().Debug("engine-end: resolve")
	close(game.In)
	vii.GameService.Forget(game.Key)
	return nil
}

func (event *EndEvent) Receive(game *vii.Game, seat *vii.GameSeat, json vii.Json) {
	game.Logger.WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-end: receive")
}

func (event *EndEvent) Json(game *vii.Game) vii.Json {
	return js.Object{
		"winner": event.Winner,
		"loser":  event.Loser,
	}
}
