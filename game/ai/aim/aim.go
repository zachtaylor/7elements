package aim

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game/ai/view"
)

// EnemyBeing picks a GCID to target for opponent Present beings given an intention string
//
// effect recognizes
// - "sleep"
func EnemyBeing(view view.T, effect string) (target interface{}, score int) {
	for tokenID := range view.Enemy.T.Present {
		t := view.Game.Token(tokenID)
		if t.T.Body == nil {
			// won't target items
		} else if effect == "sleep" && t.T.Awake {
			// won't sleep sleeping bodies
		} else if score < 3*t.T.Body.Attack {
			target = tokenID
			score = 3 * t.T.Body.Attack
		}
	}
	return
}

// EnemyPastBeingItem picks a GCID to target "past-being-item"
func EnemyPastBeingItem(view view.T) (target string, score int) {
	for cardID := range view.Enemy.T.Past {
		pastCard := view.Game.Card(cardID)
		if pastCard.T.Kind == card.Item && score < 2 {
			target = cardID
			score = 2
		} else if pastCard.T.Kind == card.Being && score < pastCard.T.Body.Life {
			target = cardID
			score = pastCard.T.Body.Life
		}
	}
	return
}

// MyPastBeing picks a GCID to target "mypast-being" with score
func MyPastBeing(view view.T) (target interface{}, score int) {
	for cardID := range view.Self.T.Past {
		pastCard := view.Game.Card(cardID)

		if pastCard.T.Kind != card.Being {
			continue
		}

		val := 3 * pastCard.T.Body.Attack

		if score < val {
			target = cardID
			score = val
		}
	}
	return
}

// MyPresentBeing picks a GCID to target "being" with score
//
// effect recognizes
// - "health"
// - "wake"
func MyPresentBeing(view view.T, effect string) (target string, score int) {
	for tokenID := range view.Self.T.Present {
		t := view.Game.Token(tokenID)

		if t.T.Body == nil {
		} else if t.T.Awake && effect == "wake" {
		} else if effect == "health" {
			if score < 6-t.T.Body.Life {
				target = tokenID
				score = 6 - t.T.Body.Life
			}
		} else {
			if score < 2*t.T.Body.Life {
				target = tokenID
				score = 2 * t.T.Body.Life
			}
		}
	}
	return
}

// MyPresentBeingItem picks a GCID to target "being-item" with score
//
// effect recognizes
// - "wake"
// - "health"
func MyPresentBeingItem(view view.T, effect string) (target string, score int) {
	for tokenID := range view.Self.T.Present {
		t := view.Game.Token(tokenID)

		if t.T.Awake && effect == "wake" {
			// score 0 skip
		} else if effect == "health" {
			if t.T.Body == nil {
				continue
			}
			if score < 1-t.T.Body.Life {
				target = tokenID
				score = 1 - t.T.Body.Life
			}
		} else if t.T.Body == nil {
			if len(t.T.Powers.GetTrigger("")) > 0 && score < 4 {
				target = tokenID
				score = 4
			}
		} else if score < 3*t.T.Body.Attack {
			target = tokenID
			score = 3 * t.T.Body.Attack
		}
	}
	return
}
