package ai

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/gameserver"
	"taylz.io/http/websocket"
	"taylz.io/log"
)

type AI struct {
	Settings
	Name string
	// Request causes the AI to submit input data to the game
	Request RequestFunc
	Game    *game.T
	Seat    *seat.T
}

func New(name string) *AI {
	return &AI{
		Settings: DefaultSettings(),
		Name:     name,
	}
}

func (ai *AI) Connect(game *game.T) {
	ai.Game = game
	ai.Seat = game.Seats.Get(ai.Name)
	ai.Request = NewRequestFunc(game, ai.Name)
	ai.Request("connect", nil)
}

func (ai *AI) Entry(log *log.T, cards card.Prototypes, decks deck.Prototypes) *gameserver.Entry {
	return &gameserver.Entry{
		Deck:   GetDeck(log, cards, decks, ai.Name),
		Writer: &Input{ai},
	}
}

// RequestPass causes the AI to submit "pass" to the current state
func (ai *AI) RequestPass() {
	ai.Request("pass", websocket.MsgData{
		"pass": ai.Game.State.ID(),
	})
}

func (ai *AI) GameState(data websocket.MsgData) {
	if ai.Game.State.Phase.Name() == "start" {
		if ai.Game.State.Reacts[ai.Seat.Username] == "" {
			ai.Request(ai.Game.State.ID(), websocket.MsgData{
				"choice": "keep",
			})
		}
	} else if ai.Game.State.Phase.Name() == "sunrise" {
		if choice := ai.NewPlan(); choice == nil {
			ai.requestSunriseElement()
		} else {
			choice.Submit(ai.Request)
		}
	} else if ai.Game.State.Phase.Name() == "end" {
		ai.RequestPass()
		ai.Request("disconnect", nil)
	} else {
		if choice := ai.NewPlan(); choice == nil {
			ai.RequestPass()
		} else {
			choice.Submit(ai.Request)
		}
	}
}

func (ai *AI) GameChoice(data websocket.MsgData) {
	if ai.Game.State.Phase.Seat() != ai.Seat.Username {
		ai.RequestPass()
		return
	}
	log := ai.Game.Log().With(websocket.MsgData{
		"Data": data,
	})
	if data["prompt"] == `Create a New Element` {
		log.Info("create element")
		ai.Request(ai.Game.State.ID(), websocket.MsgData{
			"choice": int(ai.getNewElement()),
		})
	} else {
		log.Error("wat tho? i can't make choices. i'm just here, man")
	}
}

// func (ai *AI) EventAnimate(data websocket.MsgData) {
// }

// func (ai *AI) EventTrigger(data websocket.MsgData) {
// 	ai.SendGameRequest(websocket.MsgData{
// 		"event": "pass",
// 		"mode":  "trigger",
// 	})
// }

// func (ai *AI) EventSpawn(data websocket.MsgData) {
// }

// func (ai *AI) EventAlert(data websocket.MsgData) {
// }

// func (ai *AI) EventEnd(data websocket.MsgData) {
// }

// func (ai *AI) EventMain(data websocket.MsgData) {
// 	if data["username"] == ai.Seat.Username {
// 		ai.sendPlayEvent()
// 	} else {
// 		ai.SendGameRequest(websocket.MsgData{
// 			"event": "pass",
// 			"mode":  "main",
// 		})
// 	}
// }

// func (ai *AI) EventPlay(data websocket.MsgData) {
// 	ai.SendGameRequest(websocket.MsgData{
// 		"event": "pass",
// 		"mode":  "play",
// 	})
// }

// func (ai *AI) EventPass(data websocket.MsgData) {
// }

// func (ai *AI) EventResolve(data websocket.MsgData) {
// }

// func (ai *AI) EventSunset(data websocket.MsgData) {
// 	ai.SendGameRequest(websocket.MsgData{
// 		"event": "pass",
// 		"mode":  "sunset",
// 	})
// }

// func (ai *AI) EventAttack(data websocket.MsgData) {
// 	ai.SendGameRequest(websocket.MsgData{
// 		"event": "pass",
// 		"mode":  "attack",
// 	})
// }

// func (ai *AI) EventDefend(data websocket.MsgData) {
// 	ai.SendGameRequest(websocket.MsgData{
// 		"event": "pass",
// 		"mode":  "defend",
// 	})
// }

func (ai *AI) requestSunriseElement() {
	ai.Request(ai.Game.State.ID(), websocket.MsgData{
		"choice": int(ai.getNewElement()),
	})
}
