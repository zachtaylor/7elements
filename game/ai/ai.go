package ai

import (
	"github.com/zachtaylor/7elements/content"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai/view"
)

type AI struct {
	Settings
	Name string
	View view.T
	done chan struct{}
}

func New(name string) *AI {
	return &AI{
		Settings: DefaultSettings(),
		Name:     name,
		done:     make(chan struct{}),
	}
}

func (ai *AI) Connect(g *game.G) {
	ai.View.Game = g
	ai.View.Self = g.Player(ai.Name)
	for _, playerID := range g.Players() {
		if playerID != ai.View.Self.ID() { // todo err nil
			ai.View.Enemy = g.Player(playerID)
		}
	}
	ai.Request("connect", nil)
}

func (ai *AI) Request(uri string, json map[string]any) {
	ai.View.Game.AddRequest(game.NewReq(ai.Name, uri, json))
}

func (ai *AI) Entry(content content.T) game.Entry {
	return game.NewEntry(&Writer{AI: ai}, GetDeck(content).Cards)
}

// RequestPass causes the AI to submit "pass" to the current state
func (ai *AI) RequestPass() {
	ai.Request("pass", map[string]any{
		"pass": ai.View.State.ID(),
	})
}

func (ai *AI) GameState(data map[string]any) {
	ai.View.Update(ai.View.Game.State(data["id"].(string)))

	if ai.View.State.T.Phase.Type() == "start" {
		ai.Request(ai.View.State.ID(), map[string]any{
			"choice": "keep",
		})
	} else if ai.View.State.T.Phase.Type() == "sunrise" {
		if choice := ai.NewPlan(); choice == nil {
			ai.requestSunriseElement()
		} else {
			choice.Submit(ai.Request)
		}
	} else if ai.View.State.T.Phase.Type() == "end" {
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
	if ai.View.State.T.Phase.Priority()[0] != ai.Name {
		ai.RequestPass()
		return
	}
	log := ai.View.Game.Log().With(map[string]any{
		"Data": data,
	})
	if data["prompt"] == `Create a New Element` {
		log.Info("create element")
		ai.Request(ai.View.State.ID(), map[string]any{
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
	ai.Request(ai.View.State.ID(), map[string]any{
		"choice": int(ai.getNewElement()),
	})
}
