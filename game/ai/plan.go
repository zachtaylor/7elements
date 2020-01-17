package ai

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

// Plan is a kind of thing an AI may elect to do in game
type Plan interface {
	Score() int
	Submit(*AI)
}

func (ai *AI) NewPlans() []Plan {
	var plans []Plan
	if ai.Game.State.Name() == "sunrise" {
		plans = append(plans, new(ChoiceElementPlan))
		return plans
	}
	for _, t := range ai.Seat.Present {
		if cs := ai.plansFromPresent(t); len(cs) > 0 {
			plans = append(plans, cs...)
		}
	}
	for _, c := range ai.Seat.Hand {
		if ai.Game.State.Name() != `main` && ai.Game.State.R.Seat() != ai.Seat.Username && c.Card.Type != vii.CTYPspell {
		} else if !ai.Seat.Karma.Active().Test(c.Card.Costs) {
		} else if cs := ai.plansFromHand(c); len(cs) > 0 {
			plans = append(plans, cs...)
		}
	}
	return plans
}

func (ai *AI) plansFromHand(card *game.Card) []Plan {
	switch card.Card.Type {
	case vii.CTYPbody:
		if ai.Game.State.Name() != "main" || ai.Game.State.R.Seat() != ai.Seat.Username {
			return nil
		}
		return []Plan{&PlayPlan{
			ID:     card.ID,
			Target: nullJSON,
			score:  1,
		}}
	case vii.CTYPitem:
		if ai.Game.State.Name() != "main" || ai.Game.State.R.Seat() != ai.Seat.Username {
			return nil
		}
		return []Plan{&PlayPlan{
			ID:     card.ID,
			Target: nullJSON,
			score:  1,
		}}
	case vii.CTYPspell:
		if ps := card.Card.Powers.GetTrigger("play"); len(ps) < 1 {
		} else if p := ps[0]; p == nil {
		} else if id, score := ai.ScoreCardPower(card, p); score < 1 {
		} else {
			return []Plan{&PlayPlan{
				ID:     card.ID,
				Target: id,
			}}
		}
	}
	return nil
}

func (ai *AI) plansFromPresent(t *game.Token) []Plan {
	var plans []Plan

	// add attack option
	if t.Body == nil {

	} else if t.IsAwake && ai.Game.State.Name() == `main` && ai.Game.State.R.Seat() == ai.Seat.Username {
		if ai.Settings.Aggro {
			plans = append(plans, &AttackPlan{
				TID:   t.ID,
				score: 5 * t.Body.Attack,
			})
		} else {
			plans = append(plans, &AttackPlan{
				TID:   t.ID,
				score: 3 * t.Body.Attack,
			})
		}
	}

	activeElements := ai.Seat.Karma.Active()

	// add triggerable powers option
	if ps := t.Powers.GetTrigger(``); len(ps) < 1 {
	} else if p := ps[0]; p == nil {
	} else if !activeElements.Test(p.Costs) {
	} else if p.UsesTurn && !t.IsAwake {
	} else if target, score := ai.ScoreTokenPower(t, p); score < 1 {
	} else {
		plans = append(plans, &TriggerPlan{
			TID:     t.ID,
			PowerID: p.Id,
			Target:  target,
			score:   score,
		})
	}

	return plans
}
