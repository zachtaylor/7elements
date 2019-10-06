package ai

import (
	"time"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const Username = "A.I."

type AI struct {
	Delay time.Duration
	Game  *game.T
	Seat  *game.Seat
}

func ConnectAI(g *game.T) {
	seat := g.GetSeat(Username)
	ai := &AI{2 * time.Second, g, seat}
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
		ai.Game.Log().With(log.Fields{
			"GameId":    ai.Game.ID(),
			"Username":  ai.Seat.Username,
			"EventName": name,
		}).Warn("ai: event not recognized")
	}
}

func (ai *AI) GameState(data cast.JSON) {
	if ai.Game.State.EventName() == "start" {
		if ai.Game.State.Reacts[ai.Seat.Username] == "" {
			ai.Request(ai.Game.State.ID(), cast.JSON{
				"choice": "keep",
			})
		}
	} else if ai.Game.State.EventName() == "sunrise" {
		if ai.Game.State.Event.Seat() == ai.Seat.Username {
			ai.requestSunriseElement()
		} else {
			ai.RequestPass()
		}
	} else if ai.Game.State.EventName() == "play" {
		ai.RequestPass()
	} else if ai.Game.State.EventName() == "trigger" {
		ai.RequestPass()
	} else if ai.Game.State.EventName() == "main" {
		ai.GameStateMain()
	} else if ai.Game.State.EventName() == "attack" {
		ai.RequestPass()
	} else if ai.Game.State.EventName() == "combat" {
		ai.RequestPass()
	} else if ai.Game.State.EventName() == "sunset" {
		ai.RequestPass()
	} else if ai.Game.State.EventName() == "end" {
		ai.RequestPass()
		ai.Request("disconnect", nil)
	} else {
		ai.Game.Log().With(log.Fields{
			"State": ai.Game.State.EventName(),
		}).Warn("ai: state unrecognized")
	}
}

func (ai *AI) GameChoice(data cast.JSON) {
	if ai.Game.State.Event.Seat() != ai.Seat.Username {
		ai.RequestPass()
		return
	}
	log := ai.Game.Log().With(log.Fields{
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

// func (ai *AI) EventSunrise(data cast.JSON) {
// 	if data["username"] == ai.Seat.Username {
// 		ai.sendElementEvent()
// 	}
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
		"elementid": int(ai.getNewElement()),
	})
}

func (ai *AI) GameStateMain() {
	if ai.Game.State.Event.Seat() != ai.Seat.Username {
	} else if hand := ai.getHandCanAfford(); len(hand) > 0 {
		ai.gameStateMainPlay(hand)
		return
	} else if awake := ai.getPresentCanAttack(); len(awake) > 0 {
		ai.gameStateMainAttack(awake)
		return
	}
	ai.RequestPass()
}

func (ai *AI) gameStateMainPlay(hand []string) {
	ai.Game.Log().Add("Choices", hand).Add("Choice", hand[0]).Info("ai: choose")
	ai.Request("play", cast.JSON{
		"gcid": hand[0],
	})
}

func (ai *AI) gameStateMainAttack(awake []string) {
	for _, gcid := range awake {
		card := ai.Game.Cards[gcid]
		if card.IsAwake && card.Body != nil && card.Body.Attack > 0 {
			ai.Request("attack", cast.JSON{
				"gcid": gcid,
			})
			return
		}
	}
	ai.RequestPass()
}
