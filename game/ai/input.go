package ai

import "ztaylor.me/cast"

// Input is where the ai receives input
type Input struct {
	AI *AI
}

// Send implements game.Player
//
// receives data from the game runtime, initiates plans and
func (i *Input) Send(uri string, json cast.JSON) {
	<-cast.After(i.AI.Delay)
	if uri == "/game/state" {
		i.AI.GameState(json)
	} else if uri == "/game/choice" {
		i.AI.GameChoice(json)
	} else if uri == "/game" {
	} else if uri == "/game/react" {
	} else if uri == "/game/card" {
	} else if uri == "/game/hand" {
	} else if uri == "/game/seat" {
	} else if uri == "/alert" {
	} else {
		i.AI.Game.Log().With(cast.JSON{
			"URI":      uri,
			"GameId":   i.AI.Game.ID(),
			"Username": i.AI.Seat.Username,
		}).Warn("uri unknown")
	}
}
