package ai

import (
	"time"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

type AI struct {
	Delay time.Duration
	Game  *game.T
	Seat  *game.Seat
}

func ConnectAI(game *game.T, seat *game.Seat) {
	ai := &AI{3 * time.Second, game, seat}
	go seat.Login(game, ai)
}

// Request causes the AI to submit input data to the game
func (ai *AI) Request(uri string, json vii.Json) {
	ai.Game.Log().WithFields(log.Fields{
		"URI":  uri,
		"Data": json.String(),
	}).Debug("ai: request")
	ai.Game.Request(ai.Seat.Username, uri, json)
}

func (ai *AI) RequestPass() {
	ai.Request("pass", vii.Json{
		"pass": ai.Game.State.ID(),
	})
}

// WriteJson receives data from the game engine as a GameSeat.Listener
func (ai *AI) WriteJson(json js.Object) {
	uri := json.Sval("uri")
	data := json.Val("data")
	switch v := data.(type) {
	case js.Object:
		go ai.send(uri, v)
	default:
		ai.Game.Log().Add("data", json).Warn("ai: i don't understand")
	}
}

func (ai *AI) send(name string, json js.Object) {
	<-time.After(ai.Delay)
	if name == "/game" {
		ai.EventGame(json)
	} else if name == "/game/react" {
		ai.EventGameReact(json)
	} else if name == "/game/state" {
		ai.EventGameState(json)
	} else if name == "/game/hand" {

	} else if name == "/alert" {

	} else {
		ai.Game.Log().WithFields(log.Fields{
			"GameId":    ai.Game.ID(),
			"Username":  ai.Seat.Username,
			"EventName": name,
		}).Warn("ai: event not recognized")
	}

	// old ai code
	// if name == "animate" {
	// 	ai.EventAnimate(json)
	// } else if name == "spawn" {
	// 	ai.EventSpawn(json)
	// } else if name == "sunrise" {
	// 	ai.EventSunrise(json)
	// } else if name == "play" {
	// 	ai.EventPlay(json)
	// } else if name == "main" {
	// 	ai.EventMain(json)
	// } else if name == "trigger" {
	// 	ai.EventTrigger(json)
	// } else if name == "pass" {
	// 	ai.EventPass(json)
	// } else if name == "resolve" {
	// 	ai.EventResolve(json)
	// } else if name == "sunset" {
	// 	ai.EventSunset(json)
	// } else if name == "attack" {
	// 	ai.EventAttack(json)
	// } else if name == "defend" {
	// 	ai.EventDefend(json)
	// } else if name == "end" {
	// 	ai.EventEnd(json)
	// }
}

func (ai *AI) EventGame(data js.Object) {
	if ai.Game.State.EventName() == "start" {
		if ai.Game.State.Reacts[ai.Seat.Username] == "" {
			ai.Request(ai.Game.State.ID(), js.Object{
				"choice": "keep",
			})
		}
	}
}

func (ai *AI) EventGameReact(data js.Object) {
}

func (ai *AI) EventGameState(data js.Object) {
	if ai.Game.State.EventName() == "sunrise" {
		if ai.Game.State.Seat == ai.Seat.Username {
			ai.requestSunriseElement()
		} else {
			ai.RequestPass()
		}
	}
}

// func (ai *AI) EventAnimate(data js.Object) {
// }

// func (ai *AI) EventTrigger(data js.Object) {
// 	ai.SendGameRequest(js.Object{
// 		"event": "pass",
// 		"mode":  "trigger",
// 	})
// }

// func (ai *AI) EventSpawn(data js.Object) {
// }

// func (ai *AI) EventAlert(data js.Object) {
// }

// func (ai *AI) EventEnd(data js.Object) {
// }

// func (ai *AI) EventSunrise(data js.Object) {
// 	if data["username"] == ai.Seat.Username {
// 		ai.sendElementEvent()
// 	}
// }

// func (ai *AI) EventMain(data js.Object) {
// 	if data["username"] == ai.Seat.Username {
// 		ai.sendPlayEvent()
// 	} else {
// 		ai.SendGameRequest(js.Object{
// 			"event": "pass",
// 			"mode":  "main",
// 		})
// 	}
// }

// func (ai *AI) EventPlay(data js.Object) {
// 	ai.SendGameRequest(js.Object{
// 		"event": "pass",
// 		"mode":  "play",
// 	})
// }

// func (ai *AI) EventPass(data js.Object) {
// }

// func (ai *AI) EventResolve(data js.Object) {
// }

// func (ai *AI) EventSunset(data js.Object) {
// 	ai.SendGameRequest(js.Object{
// 		"event": "pass",
// 		"mode":  "sunset",
// 	})
// }

// func (ai *AI) EventAttack(data js.Object) {
// 	ai.SendGameRequest(js.Object{
// 		"event": "pass",
// 		"mode":  "attack",
// 	})
// }

// func (ai *AI) EventDefend(data js.Object) {
// 	ai.SendGameRequest(js.Object{
// 		"event": "pass",
// 		"mode":  "defend",
// 	})
// }

func (ai *AI) requestPlay() {
	choices := make([]string, 0)
	elements := ai.Seat.Elements.GetActive()
	for gcid, gc := range ai.Seat.Hand {
		log := ai.Game.Log().WithFields(log.Fields{
			"gcid":     gcid,
			"CardId":   gc.Card.Id,
			"Cost":     gc.Card.Costs,
			"Elements": ai.Seat.Elements,
		})

		if elements.Test(gc.Card.Costs) {
			choices = append(choices, gcid)
			log.Debug("ai: choice saved")
		} else {
			log.Debug("ai: cannot afford")
		}
	}

	if len(choices) == 0 {
		ai.RequestPass()
		return
	}

	ai.Game.Log().Add("Choices", choices).Add("Choice", choices[0]).Info("ai: choose")
	ai.Request("play", vii.Json{
		"gcid": choices[0],
	})
}

func (ai *AI) requestSunriseElement() {
	devo := vii.ElementMap{}
	for _, card := range ai.Seat.Hand {
		for element, amount := range card.Card.Costs {
			devo[element] += amount
		}
	}
	for _, gc := range ai.Seat.Present {
		for _, power := range gc.Powers {
			for element, amount := range power.Costs {
				devo[element] += amount
			}
		}
	}
	delete(devo, 0)

	var element float64
	var highCount int
	for e, c := range devo {
		if c > highCount {
			element = float64(e)
			highCount = c
		}
	}

	if element < 1 {
		element = 4
		ai.Game.Log().Add("CardsInHand", len(ai.Seat.Hand)).Add("Devo", devo).Warn("ai: no element devotion")
	}

	ai.Request(ai.Game.State.ID(), vii.Json{
		"elementid": element,
	})
}
