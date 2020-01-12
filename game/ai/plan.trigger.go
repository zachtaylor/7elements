package ai

import "ztaylor.me/cast"

// TriggerPlan is a plan to trigger an ability
type TriggerPlan struct {
	TID     string
	PowerID int
	Target  interface{}
	score   int
}

func (trigger *TriggerPlan) Score() int {
	return trigger.score
}
func (trigger *TriggerPlan) Submit(ai *AI) {
	ai.Request("trigger", cast.JSON{
		"tid":     trigger.TID,
		"powerid": trigger.PowerID,
		"target":  trigger.Target,
	})
}
func (trigger *TriggerPlan) String() string {
	return "Trigger " + trigger.TID
}
