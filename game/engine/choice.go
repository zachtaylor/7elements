package engine

// import (
// 	"github.com/zachtaylor/7elements"
// 	"github.com/zachtaylor/7elements/game"
// 	"ztaylor.me/js"
// 	"ztaylor.me/log"
// )

// type ChoiceEvent struct {
// 	Stack    *game.State
// 	Text     string
// 	Choices  []string
// 	Finisher func(answer string)
// 	answer   string
// }

// func Choice(stack *game.State, text string, choices []string, fin func(answer string)) game.ZEvent {
// 	return &ChoiceEvent{
// 		Stack:    stack,
// 		Text:     text,
// 		Choices:  choices,
// 		Finisher: fin,
// 	}
// }

// func (event *ChoiceEvent) Name() string {
// 	return "choice"
// }

// // OnActivate implements game.ActivateEventer
// func (event ChoiceEvent) OnActivate(*game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event ChoiceEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Finish implements game.FinishEventer
// func (event *ChoiceEvent) Finish(*game.T) {
// }

// // GetStack implements game.StackEventer
// func (event *ChoiceEvent) GetStack(game *game.T) *game.State {
// 	return event.Stack
// }

// func (event *ChoiceEvent) GetNext(game *game.T) *game.State {
// 	return nil
// }

// func (event *ChoiceEvent) Json(g *game.T) vii.Json {
// 	return js.Object{
// 		"choice":  event.Text,
// 		"options": event.Choices,
// 	}
// }

// // Request implements game.RequestEventer
// func (event *ChoiceEvent) Request(game *game.T, seat *game.Seat, json vii.Json) {
// 	if seat.Username != game.State.Seat {
// 		game.Log().WithFields(log.Fields{
// 			"Seat": seat,
// 			"json": json,
// 		}).Warn("engine/choice: receive")
// 		return
// 	}

// 	choice := json.Sval("choice")

// 	for _, v := range event.Choices {
// 		if v == choice {
// 			event.Finisher(v)
// 		}
// 	}
// }
