package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

func Sunrise(game *game.T) game.Event {
	game.Log().Info("Sunrise")
	return new(SunriseEvent)
}

type SunriseEvent struct {
	element vii.Element
}

func (event *SunriseEvent) Name() string {
	return "sunrise"
}

// // OnActivate implements game.ActivateEventer
// func (event *SunriseEvent) OnActivate(game *game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *SunriseEvent) OnConnect(*game.T, *game.Seat) {
// }

// // GetStack implements game.StackEventer
// func (event *SunriseEvent) GetStack(game *game.T) *game.State {
// 	return nil
// }

// Finish implements game.FinishEventer
func (event *SunriseEvent) Finish(game *game.T) {
	seat := game.GetSeat(game.State.Seat)
	seat.Elements.Append(event.element)
	card := seat.Deck.Draw()
	seat.Hand[card.Id] = card
	seat.Reactivate()
	animate.GameHand(game, seat)
	animate.GameSeat(game, seat)
}

func (event *SunriseEvent) GetNext(game *game.T) *game.State {
	return game.NewState(game.State.Seat, Main(game))
}

func (event *SunriseEvent) Json(game *game.T) vii.Json {
	return nil
}

// Request implements game.RequestEventer
func (event *SunriseEvent) Request(game *game.T, seat *game.Seat, json vii.Json) {
	log := game.Logger.WithFields(log.Fields{
		"Seat": seat.Username,
	})
	if seat.Username != game.State.Seat {
		log.Add("MyTurn", game.State.Seat).Warn("engine/sunrise: request")
		return
	}
	elementID := json.Ival("elementid")
	if elementID < 1 || elementID > 7 {
		log.Add("ElementId", elementID).Warn("engine/sunrise: elementid out of bounds")
		return
	}
	event.element = vii.Elements[int(elementID)]
	log.Add("Element", event.element).Info("engine/sunrise: confirm")
	game.State.Reacts[seat.Username] = "confirm"
}
