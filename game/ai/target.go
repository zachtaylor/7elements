package ai

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

// TargetEnemyBeing picks a GCID to target for opponent Present beings given an intention string
//
// effect recognizes
// - "sleep"
func (ai *AI) TargetEnemyBeing(effect string) (target interface{}, score int) {
	enemy := ai.Game.GetOpponentSeat(ai.Seat.Username)

	var token *game.Token
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

// TargetEnemyPastBeingItem picks a GCID to target "past-being-item"
//
// effect recognizes
// - "sleep"
func (ai *AI) TargetEnemyPastBeingItem() (target interface{}, score int) {
	enemy := ai.Game.GetOpponentSeat(ai.Seat.Username)

	var c *game.Card
	for _, pc := range enemy.Past {
		if pc.Card.Type == card.ItemType && score < 2 {
			c = pc
			score = 2
		} else if pc.Card.Type == card.BodyType && score < pc.Card.Body.Health {
			c = pc
			score = pc.Card.Body.Health
		}
	}

	if c != nil {
		target = c.ID
	}
	return
}

// TargetMyPastBeing picks a GCID to target "mypast-being" with score
func (ai *AI) TargetMyPastBeing() (target interface{}, score int) {
	var c *game.Card
	for _, pc := range ai.Seat.Past {
		if pc.Card.Type != card.BodyType {
		} else if score < 3*pc.Card.Body.Attack {
			c = pc
			score = 3 * pc.Card.Body.Attack
		}
	}

	if c != nil {
		target = c.ID
	}
	return
}

// TargetMyPresentBeing picks a GCID to target "being" with score
//
// effect recognizes
// - "health"
// - "wake"
func (ai *AI) TargetMyPresentBeing(effect string) (target interface{}, score int) {
	var token *game.Token
	for _, t := range ai.Seat.Present {
		if t.Card.Card.Type != card.BodyType {
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

// TargetMyPresentBeingItem picks a GCID to target "being-item" with score
//
// effect recognizes
// - "wake"
func (ai *AI) TargetMyPresentBeingItem(effect string) (target interface{}, score int) {
	var token *game.Token
	for _, t := range ai.Seat.Present {
		if t.IsAwake && effect == "wake" {
			// score 0 skip
		} else if effect == "health" {
			if t.Body == nil {
				continue
			}

			if ai.Settings.Aggro {
				if score < 1-t.Body.Health {
					token = t
					score = 1 - t.Body.Health
				}
			} else {
				if score < 4-t.Body.Health {
					token = t
					score = 4 - t.Body.Health
				}
			}
		} else if t.Card.Card.Type == card.ItemType {
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
