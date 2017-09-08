package games

import (
	"7elements.ztaylor.me/games/cards"
	"7elements.ztaylor.me/games/turns"
	"ztaylor.me/json"
)

func MakeTurnJson(game *Game, turn *gameturns.GameTurn) json.Json {
	seat := game.GetSeat(turn.Username)
	data := json.Json{
		"username": turn.Username,
		"element":  turn.Element,
		"timer":    int(game.Context.Timer().Seconds() + 1),
		"active":   gamecards.Stack(seat.Active).Json(),
	}

	if game.TurnPhase == TPHSwait {
		data["turnphase"] = "wait"
	} else if game.TurnPhase == TPHSbegin {
		data["turnphase"] = "begin"
	} else if game.TurnPhase == TPHSplay {
		data["turnphase"] = "play"
	} else if game.TurnPhase == TPHSattack {
		data["turnphase"] = "attack"
	} else if game.TurnPhase == TPHSdefend {
		data["turnphase"] = "defend"
	} else if game.TurnPhase == TPHSdone {
		data["turnphase"] = "done"
	}

	return data
}

func MakeDoneJson(game *Game, username string) json.Json {
	return json.Json{
		"username": username,
		"gameid":   game.Id,
		"winners":  game.GetWinners(),
		"losers":   game.GetLosers(),
	}
}

func MakeGameJson(game *Game, username string) json.Json {
	data := json.Json{
		"username": username,
		"gameid":   game.Id,
		"timer":    int(game.Context.Timer().Seconds() + 1),
	}

	seatdata := make([]json.Json, 0)
	for i, seat := range game.Seats {
		if seat.Username == username {
			data["seatid"] = i
			data["hand"] = gamecards.Stack(seat.Hand).Json()
		}

		seatdata = append(seatdata, seat.Json())
	}
	data["seats"] = seatdata

	return data
}
