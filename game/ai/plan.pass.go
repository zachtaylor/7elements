package ai

import "ztaylor.me/cast"

// PassPlan is a plan to pass
type PassPlan bool

func (pass PassPlan) Score() int {
	return 1
}
func (pass PassPlan) Submit(ai *AI) {
	ai.Request("pass", cast.JSON{
		"id": ai.Game.State.ID(),
	})
}
func (pass *PassPlan) String() string {
	return "pass"
}
