package engine

import (
	"elemen7s.com"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func End(game *vii.Game, t *Timeline) Event {
	return new(EndEvent)
}

type EndEvent bool

func (event *EndEvent) Name() string {
	return "end"
}

func (event *EndEvent) Priority(*vii.Game, *Timeline) bool {
	return true
}

func (event *EndEvent) OnStart(game *vii.Game, t *Timeline) {
	if username, log := game.Results.Winner, game.Logger.Add("Winner", game.Results.Winner); username == "" {
		log.Warn("engine-end: winner missing")
	} else if seat := game.GetSeat(username); seat == nil {
		log.Warn("engine-end: winning seat missing")
	} else if err := vii.AccountDeckService.UpdateTallyWinCount(username, seat.Deck.AccountDeckId, seat.Deck.AccountDeckVersion); err != nil {
		log.Add("Error", err).Error("engine-end: winning deck missing")
	} else if account, err := vii.AccountService.Get(username); err != nil {
		log.Add("Error", err).Error("engine-end: account missing")
	} else {
		account.Coins += 2
		if err = vii.AccountService.UpdateCoins(account); err != nil {
			log.Clone().Add("Error", err).Add("Username", username).Error("engine-end: account service error")
		}
	}

	if username, log := game.Results.Loser, game.Logger.Add("Loser", game.Results.Loser); username == "" {
		log.Warn("engine-end: loser missing")
	} else if account, err := vii.AccountService.Get(username); err != nil {
		log.Add("Error", err).Error("engine-end: account missing")
	} else {
		account.Coins++
		if err = vii.AccountService.UpdateCoins(account); err != nil {
			log.Add("Error", err).Error("engine-end: account service error")
		}
	}
}

func (event *EndEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *EndEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	game.Log().Debug("engine-end: resolve")
	close(game.In)
	vii.GameService.Forget(game.Key)
	return nil
}

func (event *EndEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	game.Logger.WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("engine-end: receive")
}

func (event *EndEvent) Json(game *vii.Game, t *Timeline) js.Object {
	return js.Object{
		"gameid": game,
		"winner": game.Results.Winner,
		"loser":  game.Results.Loser,
	}
}
