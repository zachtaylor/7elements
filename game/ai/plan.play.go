package ai

import "ztaylor.me/cast"

// PlayPlan is a plan to make a play
type PlayPlan struct {
	ID     string
	Target interface{}
	score  int
}

func (play *PlayPlan) Score() int {
	return play.score
}
func (play *PlayPlan) Submit(ai *AI) {
	ai.Request("play", cast.JSON{
		"id":     play.ID,
		"target": play.Target,
	})
}
func (play *PlayPlan) String() string {
	msg := "Play " + play.ID
	msg += "(" + cast.StringI(play.score) + ")"
	if play.Target != nil {
		msg += "[" + cast.String(play.Target) + "]"
	}

	return msg
}
