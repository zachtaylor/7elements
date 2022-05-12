package ai

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
)

type AI struct {
	Settings
	Name string
	// Request causes the AI to submit input data to the game
	Request RequestFunc
	Game    *game.T
	Seat    *seat.T
	done    chan struct{}
}

func New(name string) *AI {
	return &AI{
		Settings: DefaultSettings(),
		Name:     name,
		done:     make(chan struct{}),
	}
}

func (ai *AI) Connect(game *game.T) {
	ai.Game = game
	ai.Seat = g.Player(ai.Name)
	ai.Request = NewRequestFunc(game, ai.Name)
	ai.Request("connect", nil)
}

func (ai *AI) Entry(version *game.Version) *game.Entry {
	return game.NewEntry(GetDeck(version), &Input{ai})
}

// RequestPass causes the AI to submit "pass" to the current state
func (ai *AI) RequestPass() {
	ai.Request("pass", map[string]any{
		"pass": ai.Game.State.ID(),
	})
}

func (ai *AI) GameState(data map[string]any) {
	if ai.Game.State.Phase.Name() == "start" {
		if ai.Game.State.Reacts[ai.Seat.Username] == "" {
			ai.Request(ai.Game.State.ID(), map[string]any{
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

func (ai *AI) GameChoice(data map[string]any) {
	if ai.Game.State.Phase.Seat() != ai.Seat.Username {
		ai.RequestPass()
		return
	}
	log := ai.Game.Log().With(map[string]any{
		"Data": data,
	})
	if data["prompt"] == `Create a New Element` {
		log.Info("create element")
		ai.Request(ai.Game.State.ID(), map[string]any{
			"choice": int(ai.getNewElement()),
		})
	} else {
		log.Error("wat tho? i can't make choices. i'm just here, man")
	}
}

// func (ai *AI) EventAnimate(data map[string]any) {
// }

// func (ai *AI) EventTrigger(data map[string]any) {
// 	ai.SendGameRequest(map[string]any{
// 		"event": "pass",
// 		"mode":  "trigger",
// 	})
// }

// func (ai *AI) EventSpawn(data map[string]any) {
// }

// func (ai *AI) EventAlert(data map[string]any) {
// }

// func (ai *AI) EventEnd(data map[string]any) {
// }

// func (ai *AI) EventMain(data map[string]any) {
// 	if data["username"] == ai.Seat.Username {
// 		ai.sendPlayEvent()
// 	} else {
// 		ai.SendGameRequest(map[string]any{
// 			"event": "pass",
// 			"mode":  "main",
// 		})
// 	}
// }

// func (ai *AI) EventPlay(data map[string]any) {
// 	ai.SendGameRequest(map[string]any{
// 		"event": "pass",
// 		"mode":  "play",
// 	})
// }

// func (ai *AI) EventPass(data map[string]any) {
// }

// func (ai *AI) EventResolve(data map[string]any) {
// }

// func (ai *AI) EventSunset(data map[string]any) {
// 	ai.SendGameRequest(map[string]any{
// 		"event": "pass",
// 		"mode":  "sunset",
// 	})
// }

// func (ai *AI) EventAttack(data map[string]any) {
// 	ai.SendGameRequest(map[string]any{
// 		"event": "pass",
// 		"mode":  "attack",
// 	})
// }

// func (ai *AI) EventDefend(data map[string]any) {
// 	ai.SendGameRequest(map[string]any{
// 		"event": "pass",
// 		"mode":  "defend",
// 	})
// }

func (ai *AI) requestSunriseElement() {
	ai.Request(ai.Game.State.ID(), map[string]any{
		"choice": int(ai.getNewElement()),
	})
}
