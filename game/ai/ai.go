package ai

import (
	"time"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

const Username = "A.I."

var nullJSON = cast.Stringer(`null`)

type AI struct {
	Game *game.T
	Seat *game.Seat
	Settings
}

func ConnectAI(g *game.T) {
	seat := g.GetSeat(Username)
	ai := &AI{g, seat, DefaultSettings()}
	seat.Receiver = ai
	g.Request(seat.Username, "connect", nil)
}

// Request causes the AI to submit input data to the game
func (ai *AI) Request(uri string, json cast.JSON) {
	ai.Game.Request(ai.Seat.Username, uri, json)
}

func (ai *AI) RequestPass() {
	ai.Request("pass", cast.JSON{
		"pass": ai.Game.State.ID(),
	})
}

// WriteJSON receives data from the game engine as a GameSeat.Listener
func (ai *AI) WriteJSON(json cast.JSON) {
	uri := json.GetS("uri")
	data := json["data"]
	switch v := data.(type) {
	case cast.JSON:
		go ai.receive(uri, v)
	default:
		ai.Game.Log().Add("data", json).Warn("ai: i don't understand")
	}
}

func (ai *AI) receive(name string, json cast.JSON) {
	<-time.After(ai.Delay)
	if name == "/game/state" {
		ai.GameState(json)
	} else if name == "/game/choice" {
		ai.GameChoice(json)
	} else if name == "/game" {
	} else if name == "/game/react" {
	} else if name == "/game/card" {
	} else if name == "/game/hand" {
	} else if name == "/game/seat" {
	} else if name == "/alert" {
	} else {
		ai.Game.Log().With(cast.JSON{
			"GameId":   ai.Game.ID(),
			"Username": ai.Seat.Username,
			"Name":     name,
		}).Warn("ai: event not recognized")
	}
}

func (ai *AI) GameState(data cast.JSON) {
	if ai.Game.State.Name() == "start" {
		if ai.Game.State.Reacts[ai.Seat.Username] == "" {
			ai.Request(ai.Game.State.ID(), cast.JSON{
				"choice": "keep",
			})
		}
	} else if ai.Game.State.Name() == "sunrise" {
		if choice := ai.NewPlan(); choice == nil {
			// ai.requestSunriseElement()
		} else {
			choice.Submit(ai)
		}
	} else if ai.Game.State.Name() == "play" {
		ai.RequestPass()
	} else if ai.Game.State.Name() == "trigger" {
		ai.RequestPass()
	} else if ai.Game.State.Name() == "main" {
		if choice := ai.NewPlan(); choice == nil {
			ai.RequestPass()
		} else {
			choice.Submit(ai)
		}
	} else if ai.Game.State.Name() == "attack" {
		ai.RequestPass()
	} else if ai.Game.State.Name() == "combat" {
		ai.RequestPass()
	} else if ai.Game.State.Name() == "sunset" {
		ai.RequestPass()
	} else if ai.Game.State.Name() == "end" {
		ai.RequestPass()
		ai.Request("disconnect", nil)
	} else {
		ai.Game.Log().With(cast.JSON{
			"State": ai.Game.State.Name(),
		}).Warn("ai: state unrecognized")
	}
}

func (ai *AI) GameChoice(data cast.JSON) {
	if ai.Game.State.R.Seat() != ai.Seat.Username {
		ai.RequestPass()
		return
	}
	log := ai.Game.Log().With(cast.JSON{
		"Data": data.String(),
	}).Tag("ai/gamechoice")
	if data["prompt"] == `Create a New Element` {
		log.Info("create element")
		ai.Request(ai.Game.State.ID(), cast.JSON{
			"choice": int(ai.getNewElement()),
		})
	} else {
		log.Error("wat tho")
	}
}

// func (ai *AI) EventAnimate(data cast.JSON) {
// }

// func (ai *AI) EventTrigger(data cast.JSON) {
// 	ai.SendGameRequest(cast.JSON{
// 		"event": "pass",
// 		"mode":  "trigger",
// 	})
// }

// func (ai *AI) EventSpawn(data cast.JSON) {
// }

// func (ai *AI) EventAlert(data cast.JSON) {
// }

// func (ai *AI) EventEnd(data cast.JSON) {
// }

// func (ai *AI) EventMain(data cast.JSON) {
// 	if data["username"] == ai.Seat.Username {
// 		ai.sendPlayEvent()
// 	} else {
// 		ai.SendGameRequest(cast.JSON{
// 			"event": "pass",
// 			"mode":  "main",
// 		})
// 	}
// }

// func (ai *AI) EventPlay(data cast.JSON) {
// 	ai.SendGameRequest(cast.JSON{
// 		"event": "pass",
// 		"mode":  "play",
// 	})
// }

// func (ai *AI) EventPass(data cast.JSON) {
// }

// func (ai *AI) EventResolve(data cast.JSON) {
// }

// func (ai *AI) EventSunset(data cast.JSON) {
// 	ai.SendGameRequest(cast.JSON{
// 		"event": "pass",
// 		"mode":  "sunset",
// 	})
// }

// func (ai *AI) EventAttack(data cast.JSON) {
// 	ai.SendGameRequest(cast.JSON{
// 		"event": "pass",
// 		"mode":  "attack",
// 	})
// }

// func (ai *AI) EventDefend(data cast.JSON) {
// 	ai.SendGameRequest(cast.JSON{
// 		"event": "pass",
// 		"mode":  "defend",
// 	})
// }

func (ai *AI) requestSunriseElement() {
	ai.Request(ai.Game.State.ID(), cast.JSON{
		"choice": int(ai.getNewElement()),
	})
}

func (ai *AI) NewPlan() (plan Plan) {
	plans := ai.NewPlans()
	log := ai.Game.Log().With(cast.JSON{
		"Plans": plans,
	}).Tag("ai/plan")
	var score int
	for _, p := range plans {
		if p.Score() > score {
			plan = p
			score = p.Score()
		}
	}
	if plan == nil {
		log.Debug("no high score")
	} else {
		log.Add("Plan", plan).Add("Score", score).Debug()
	}
	return
}
