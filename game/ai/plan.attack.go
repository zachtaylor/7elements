package ai

import "ztaylor.me/cast"

// AttackPlan is a plan to make an attack
type AttackPlan struct {
	TID   string // token id
	score int    // plan value
}

func (attack *AttackPlan) Score() int {
	return attack.score
}
func (attack *AttackPlan) Submit(ai *AI) {
	ai.Request("attack", cast.JSON{
		"id": attack.TID,
	})
}
func (attack *AttackPlan) String() string {
	return "Attack " + attack.TID
}
