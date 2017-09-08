package games

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/decks"
	"time"
	"ztaylor.me/log"
	// "ztaylor.me/ctxpert"
)

func WinGame(game *Game, username string) {
	awardWinner(game, username)

	if len(game.Seats) != 2 {
		log.Add("Seats#", len(game.Seats)).Warn("end-game: win in ffa mode???")
		return
	} else if len(game.results) == 2 {
		return
	}

	for _, seat := range game.Seats {
		if !game.results[seat.Username] && seat.Username != username {

			game.MarkLoser(username)
			accounts.Test(username).Skill--
			accounts.Delete(username)
			accounts.Insert(username)
		}
	}
	go HaltGame(game)
}

func awardWinner(game *Game, username string) {
	game.MarkWinner(username)
	accounts.Test(username).Coins++
	accounts.Test(username).Skill++
	accounts.Delete(username)
	accounts.Insert(username)
	deckid := game.GetSeat(username).DeckId
	decks.Test(username)[deckid].Wins++
	decks.Delete(username, deckid)
	decks.Insert(username, deckid)
}

func HaltGame(game *Game) {
	game.Context.Cancel()
	game.GamePhase = GPHSdone
	log.Add("GameId", game.Id).Add("Winners", game.GetWinners()).Add("Losers", game.GetLosers()).Warn("end-game: that's a wrap, let's go home folks")

	go func() {
		<-time.After(15 * time.Minute)
		log.Add("GameId", game.Id).Info("end-game: cache cleared")
		delete(Cache, game.Id)
	}()
}

func LoseGame(game *Game, username string) {
	game.MarkLoser(username)
	accounts.Test(username).Skill--
	accounts.Delete(username)
	accounts.Insert(username)

	if len(game.Seats) != 2 {
		log.Add("Seats#", len(game.Seats)).Warn("end-game: lose in ffa mode???")
		return
	} else if len(game.results) == 2 {
		return
	}

	for _, seat := range game.Seats {
		if !game.results[seat.Username] && seat.Username != username {
			awardWinner(game, seat.Username)
		}
	}
	go HaltGame(game)
}

func ForfeitGame(game *Game, username string) {
	game.MarkLoser(username)
	accounts.Delete(username)
	accounts.Insert(username)

	if len(game.Seats) != 2 {
		log.Add("Seats#", len(game.Seats)).Warn("end-game: forfeited in ffa mode???")
		return
	}

	go HaltGame(game)
}
