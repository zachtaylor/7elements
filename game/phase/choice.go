package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

type Choice struct {
	R
	Text     string
	ExData   websocket.MsgData
	Choices  []websocket.MsgData
	Finisher func(answer interface{})
	answer   interface{}
}

func NewChoice(seat, text string, data websocket.MsgData, choices []websocket.MsgData, fin func(answer interface{})) *Choice {
	return &Choice{
		R:        R(seat),
		Text:     text,
		ExData:   data,
		Choices:  choices,
		Finisher: fin,
	}
}
func choiceIsPhaser(r *Choice) game.Phaser { return r }

func (r *Choice) Name() string {
	return "choice"
}

func (r *Choice) String() string {
	return "choice (" + r.Seat() + ")"
}

// // OnActivate implements game.OnActivatePhaser
// func (r *Choice) OnActivate(game *game.T) []game.Phaser {
// 	return nil
// }
// func _activateStater(r *Choice) game.OnActivatePhaser {
// 	return r
// }

// OnConnect implements game.OnConnectPhaser
func (r *Choice) OnConnect(game *game.T, s *seat.T) {
	if s == nil {
		game.Log().Trace("send")
		game.Seats.Get(r.Seat()).Writer.Write(wsout.GameChoiceJSON(r.Text, r.Choices, nil))
	} else if s.Username == r.Seat() {
		game.Log().Trace("resend")
		s.Writer.Write(wsout.GameChoiceJSON(r.Text, r.Choices, nil))
	}
}
func (r *Choice) onConnectPhaser() game.OnConnectPhaser { return r }

// Finish implements game.OnFinishPhaser
func (r *Choice) OnFinish(*game.T) []game.Phaser {
	if r.Finisher != nil {
		r.Finisher(r.answer)
	}
	return nil
}
func (r *Choice) onFinishPhaser() game.OnFinishPhaser { return r }

func (r *Choice) GetNext(game *game.T) game.Phaser { return nil }

func (r *Choice) Data() map[string]interface{} {
	return map[string]interface{}{
		"choice":  r.Text,
		"options": r.Choices,
		"data":    r.ExData,
	}
}

// Request implements Requester
func (r *Choice) OnRequest(game *game.T, seat *seat.T, json map[string]interface{}) {
	if seat.Username != r.Seat() {
		game.Log().With(map[string]interface{}{
			"Seat": seat,
			"json": json,
		}).Warn("choice: receive")
		return
	}

	r.answer = json["choice"]
	if r.answer != "" {
		for _, seat := range game.Seats.Keys() {
			game.State.Reacts[seat] = "push"
		}
	}
}
func (r *Choice) onRequestPhaser() game.OnRequestPhaser { return r }
