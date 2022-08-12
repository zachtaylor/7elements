package ai

import "github.com/zachtaylor/7elements/game/ai/plan"

// Plan is a kind of thing an AI may elect to do in game
type Plan interface {
	Score() int
	Submit(func(string, map[string]any))
}

func (ai *AI) NewPlan() (do Plan) {
	if ai.View.State.T.Phase.Type() == "sunrise" {
		return &plan.NewElement{
			StateID: ai.View.State.ID(),
			Element: ai.getNewElement(),
		}
	}

	plans := ai.NewPlans()
	log := ai.View.Game.Log()
	var score int
	for _, p := range plans {
		if p.Score() > score {
			do = p
			score = p.Score()
		}
	}
	if do == nil {
		log.Debug("no high score")
	} else {
		log.Add("Plan", do).Add("Score", score).Debug()
	}
	return
}

func (ai *AI) NewPlans() (plans []Plan) {
	if attack := plan.ParseAttack(ai.View); attack != nil {
		plans = append(plans, attack)
	}
	if play := plan.ParsePlay(ai.View); play != nil {
		plans = append(plans, play)
	}
	if trigger := plan.ParseTrigger(ai.View); trigger != nil {
		plans = append(plans, trigger)
	}
	return
}
