package aim

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

// EnemyBeing picks a GCID to target for opponent Present beings given an intention string
//
// effect recognizes
// - "sleep"
func EnemyBeing(game *game.T, seat *seat.T, effect string) (target interface{}, score int) {
	enemy := game.Seats.GetOpponent(seat.Username)

	var token *token.T
	for _, t := range enemy.Present {
		if t.Body == nil {
			// won't target items
		} else if effect == "sleep" && t.IsAwake {
			// won't sleep sleeping bodies
		} else {
			if score < 3*t.Body.Attack {
				token = t
				score = 3 * t.Body.Attack
			}
		}
	}

	if token != nil {
		target = token.ID
	}
	return
}

// EnemyPastBeingItem picks a GCID to target "past-being-item"
//
// effect recognizes
// - "sleep"
func EnemyPastBeingItem(game *game.T, seat *seat.T) (target interface{}, score int) {
	enemy := game.Seats.GetOpponent(seat.Username)

	var c *card.T
	for _, pc := range enemy.Past {
		if pc.Proto.Type == card.ItemType && score < 2 {
			c = pc
			score = 2
		} else if pc.Proto.Type == card.BodyType && score < pc.Proto.Body.Health {
			c = pc
			score = pc.Proto.Body.Health
		}
	}

	if c != nil {
		target = c.ID
	}
	return
}

// MyPastBeing picks a GCID to target "mypast-being" with score
func MyPastBeing(game *game.T, seat *seat.T) (target interface{}, score int) {
	var c *card.T
	for _, pc := range seat.Past {
		if pc.Proto.Type != card.BodyType {
		} else if score < 3*pc.Proto.Body.Attack {
			c = pc
			score = 3 * pc.Proto.Body.Attack
		}
	}

	if c != nil {
		target = c.ID
	}
	return
}

// MyPresentBeing picks a GCID to target "being" with score
//
// effect recognizes
// - "health"
// - "wake"
func MyPresentBeing(game *game.T, seat *seat.T, effect string) (target interface{}, score int) {
	var token *token.T
	for _, t := range seat.Present {
		if t.Card.Proto.Type != card.BodyType {
		} else if t.IsAwake && effect == "wake" {
		} else if effect == "health" {
			if score < 5-t.Body.Health {
				token = t
				score = 5 - t.Body.Health
			}
		} else {
			if score < 2*t.Body.Health {
				token = t
				score = 2 * t.Body.Health
			}
		}
	}

	if token != nil {
		target = token.ID
	}
	return
}

// MyPresentBeingItem picks a GCID to target "being-item" with score
//
// effect recognizes
// - "wake"
func MyPresentBeingItem(game *game.T, seat *seat.T, effect string) (target interface{}, score int) {
	var token *token.T
	for _, t := range seat.Present {
		if t.IsAwake && effect == "wake" {
			// score 0 skip
		} else if effect == "health" {
			if t.Body == nil {
				continue
			}
			if score < 1-t.Body.Health {
				token = t
				score = 1 - t.Body.Health
			}
		} else if t.Card.Proto.Type == card.ItemType {
			if len(t.Powers.GetTrigger("")) > 0 && score < 4 {
				token = t
				score = 4
			}
		} else if score < 3*t.Body.Attack {
			token = t
			score = 3 * t.Body.Attack
		}
	}

	if token != nil {
		target = token.ID
	}
	return
}
