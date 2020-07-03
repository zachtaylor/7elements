package ai

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/power"
)

// ScoreTokenPower picks best Target for given Token and Power, and returns score for the choice
func (ai *AI) ScoreTokenPower(t *game.Token, p *power.T) (target interface{}, score int) {
	switch t.Card.Proto.ID {
	case 1:
		target, score = t.ID, 10
	case 2:
		if ai.Settings.Aggro {
			target, score = ai.Game.GetOpponentSeat(ai.Seat.Username).Username, 10
		} else if target, score = ai.TargetEnemyBeing("damage"); score < 1 {
			target, score = ai.Game.GetOpponentSeat(ai.Seat.Username).Username, 1
		}
	case 5:
		if ai.Settings.Aggro {
			if ai.Game.State.Name() != `sunrise` || !ai.myturn() {
			} else {
				target, score = ai.TargetEnemyBeing("sleep")
			}
		} else {
			if ai.Game.State.Name() != `main` || !ai.myturn() {
			} else {
				target, score = ai.TargetEnemyBeing("sleep")
			}
		}
	case 7:
		target, score = ai.TargetEnemyBeing("damage")
	case 15:
		target, score = ai.TargetEnemyPastBeingItem()
	case 20:
		target, score = ai.Seat.Username, 8-ai.Seat.Life
		if targetT, scoreT := ai.TargetMyPresentBeing("health"); scoreT > score {
			target, score = targetT, scoreT
		}
	case 21:
		target, score = ai.TargetMyPastBeing()
	}
	return
}

// ScoreCardPower picks best Target for given Card and Power, and returns score for the choice
func (ai *AI) ScoreCardPower(card *card.T, p *power.T) (target interface{}, score int) {
	switch card.Proto.ID {
	case 9:
		target, score = ai.TargetEnemyBeing("damage")
	case 10:
		target, score = ai.TargetMyPresentBeingItem("wake")
	case 11:
		target, score = ai.TargetMyPresentBeing("wake")
	case 12:
		target, score = ai.TargetEnemyBeing("")
	case 13:
		target, score = ai.TargetMyPresentBeing("health")
	case 14:
		target, score = ai.TargetMyPastBeing()
	}
	return
}
