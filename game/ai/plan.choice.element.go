package ai

import "ztaylor.me/cast"

// ChoiceElementPlan is a plan to choose an element
type ChoiceElementPlan byte

func (choice *ChoiceElementPlan) Score() int {
	return 12
}

func (choice *ChoiceElementPlan) Submit(ai *AI) {
	ai.Request(ai.Game.State.ID(), cast.JSON{
		"choice": int(ai.getNewElement()),
	})
}

func (choice *ChoiceElementPlan) String() string {
	return "Choice Element"
}
