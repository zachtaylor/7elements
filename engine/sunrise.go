package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Sunrise(game *vii.Game, hotseat string) vii.GameEvent {
	game.Log().Info("Sunrise")
	return new(SunriseEvent)
}

type SunriseEvent struct {
	vii.Element
}

func (event *SunriseEvent) Name() string {
	return "sunrise"
}

func (event *SunriseEvent) Priority(game *vii.Game) bool {
	return event.Element == vii.ELEMnull
}

func (event *SunriseEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	log := game.Log().WithFields(log.Fields{
		"Seat": seat,
	})
	if event.Element != vii.ELEMnull {
		log.Add("SavedEl", event.Element).Add("ElementId", json["elementid"]).Warn("sunrise: choice already saved")
		return
	} else if json["event"] == "element" {
		log.Add("Event", json["event"]).Warn("element: event unrecognized")
		return
	} else if seat.Username != game.State.Seat {
		log.Add("Expectgame.State.Seat", game.State.Seat).Warn("sunrise: username rejected")
	}
	elementId := json.Ival("elementid")
	if elementId < 1 || elementId > 7 {
		log.Add("ElementId", elementId).Warn("sunrise: elementid out of bounds")
		return
	}
	event.Element = vii.Elements[int(elementId)]
	seat.Elements.Append(event.Element)
	log.Add("Element", event.Element).Info("sunrise: confirm element choice")
}

func (event *SunriseEvent) OnStart(game *vii.Game) {
}

func (event *SunriseEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event *SunriseEvent) NextEvent(game *vii.Game) vii.GameEvent {
	if event.Element == vii.ELEMnull {
		game.Log().Warn("sunrise: forfeit!")
		return End(game, "", game.State.Seat) // only lose, no win
	} else if seat := game.GetSeat(game.State.Seat); seat == nil {
		game.Log().Warn("sunrise: !resolve seat missing")
	} else {
		card := seat.Deck.Draw()
		seat.Hand[card.Id] = card
		seat.Reactivate()
		animate.GameHand(game, seat)
		animate.GameSeat(game, seat)
	}
	return Main(game)
}

func (event *SunriseEvent) Json(game *vii.Game) js.Object {
	return js.Object{
		"gameid": game.Key,
		"timer":    game.State.Timer.Seconds(),
		"username": game.State.Seat,
	}
}
