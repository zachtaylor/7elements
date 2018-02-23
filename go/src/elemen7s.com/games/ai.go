package games

import (
	"elemen7s.com/decks"
	"elemen7s.com/elements"
	"time"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

type AI struct {
	Delay time.Duration
	*Game
	*Seat
}

func ConnectAI(game *Game, seat *Seat) *AI {
	ai := &AI{7 * time.Second, game, seat}
	seat.Player = ai
	return ai
}

func BuildAIGame() *Game {
	g := New()

	aideck := decks.New()
	aideck.Username = "A.I."
	aideck.Cards[1] = 3
	aideck.Cards[2] = 3
	aideck.Cards[3] = 3
	aideck.Cards[4] = 3
	aideck.Cards[5] = 3
	aideck.Cards[6] = 3
	aideck.Cards[7] = 3

	aiseat := g.Register(aideck, "en-US")

	ConnectAI(g, aiseat)

	return g
}

func (ai *AI) SendGameRequest(j js.Object) {
	go ai.Game.Receive(ai.Seat.Username, j)
}

func (ai *AI) Send(name string, json js.Object) {
	go ai.send(name, json)
}

func (ai *AI) send(name string, json js.Object) {
	<-time.After(ai.Delay)
	if name == "start" {
		ai.EventStart(json)
	} else if name == "animate" {
		ai.EventAnimate(json)
	} else if name == "spawn" {
		ai.EventSpawn(json)
	} else if name == "alert" {
		ai.EventAlert(json)
	} else if name == "game" {
		ai.EventGame(json)
	} else if name == "sunrise" {
		ai.EventSunrise(json)
	} else if name == "main" {
		ai.EventMain(json)
	} else if name == "play" {
		ai.EventPlay(json)
	} else if name == "pass" {
		ai.EventPass(json)
	} else if name == "resolve" {
		ai.EventResolve(json)
	} else if name == "sunset" {
		ai.EventSunset(json)
	} else if name == "attack" {
		ai.EventAttack(json)
	} else if name == "defend" {
		ai.EventDefend(json)
	} else if name == "end" {
		ai.EventEnd(json)
	} else {
		log.Add("GameId", ai.Game.Id).Add("Username", ai.Seat.Username).Add("EventName", name).Warn("ai: event not recognized")
	}
}

func (ai *AI) EventStart(data js.Object) {
	ai.SendGameRequest(js.Object{
		"event": "start",
		"resp":  "keep",
	})
}

func (ai *AI) EventAnimate(data js.Object) {
}

func (ai *AI) EventSpawn(data js.Object) {
}

func (ai *AI) EventAlert(data js.Object) {
}

func (ai *AI) EventGame(data js.Object) {
}

func (ai *AI) EventEnd(data js.Object) {
}

func (ai *AI) EventSunrise(data js.Object) {
	if data["username"] == ai.Seat.Username {
		ai.sendElementEvent()
	}
}

func (ai *AI) EventMain(data js.Object) {
	if data["username"] == ai.Seat.Username {
		ai.sendPlayEvent()
	} else {
		ai.SendGameRequest(js.Object{
			"event": "main",
			"resp":  "pass",
		})
	}
}

func (ai *AI) EventPlay(data js.Object) {
	ai.SendGameRequest(js.Object{
		"event": "play",
		"resp":  "pass",
	})
}

func (ai *AI) EventPass(data js.Object) {
}

func (ai *AI) EventResolve(data js.Object) {
}

func (ai *AI) EventSunset(data js.Object) {
	ai.SendGameRequest(js.Object{
		"event": "sunset",
		"resp":  "pass",
	})
}

func (ai *AI) EventAttack(data js.Object) {
	if ai.Username == data["username"] {
		ai.SendGameRequest(js.Object{
			"event": "attack",
			"resp":  "pass",
		})
	}
}

func (ai *AI) EventDefend(data js.Object) {
	ai.SendGameRequest(js.Object{
		"event": "defend",
		"resp":  "pass",
	})
}

func (ai *AI) sendPlayEvent() {
	choices := make([]int, 0)
	for gcid, gc := range ai.Hand {
		if ai.Seat.Elements.TestStack(gc.Card.Costs) {
			choices = append(choices, gcid)
			ai.Game.Log().WithFields(log.Fields{
				"CardId":  gc.Card.Id,
				"GCID":    gcid,
				"MyEl":    ai.Seat.Elements,
				"Costs":   gc.Card.Costs,
				"Choices": choices,
			}).Debug("ai: choice saved")
		}
	}

	if len(choices) == 0 {
		ai.SendGameRequest(js.Object{
			"event": "main",
			"resp":  "pass",
		})
		return
	}

	ai.Game.Log().Add("#Choices", len(choices)).Add("Choice", choices[0]).Info("ai: choose")
	ai.SendGameRequest(js.Object{
		"event": "main",
		"gcid":  choices[0],
	})
}

func (ai *AI) sendElementEvent() {
	devo := elements.Stack{}
	for _, card := range ai.Seat.Hand {
		for element, amount := range card.Card.Costs {
			devo[element] += amount
		}
	}
	for _, gc := range ai.Seat.Alive {
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

	ai.SendGameRequest(js.Object{
		"event":     "sunrise",
		"elementid": element,
	})
}