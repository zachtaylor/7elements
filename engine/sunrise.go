package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Sunrise(game *vii.Game, past *Timeline, hotseat string) Event {
	if tname := past.Name(); tname != "start" && tname != "sunset" {
		game.Log().Add("Timeline", tname).Error("sunrise can only follow start or sunset")
		return nil
	}
	game.Log().Info("sunrise")
	return new(SunriseEvent)
}

type SunriseEvent struct {
	vii.Element
}

func (event *SunriseEvent) Name() string {
	return "sunrise"
}

func (event *SunriseEvent) Priority(game *vii.Game, t *Timeline) bool {
	return event.Element == vii.ELEMnull
}

func (event *SunriseEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	log := game.Log().WithFields(log.Fields{
		"Seat": seat,
	})
	if event.Element != vii.ELEMnull {
		log.Add("SavedEl", event.Element).Add("ElementId", json["elementid"]).Warn("sunrise: choice already saved")
		return
	} else if json["event"] == "element" {
		log.Add("Event", json["event"]).Warn("element: event unrecognized")
		return
	} else if seat.Username != t.HotSeat {
		log.Add("Expectt.HotSeat", t.HotSeat).Warn("sunrise: username rejected")
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

func (event *SunriseEvent) OnStart(game *vii.Game, t *Timeline) {
}

func (event *SunriseEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *SunriseEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	if event.Element == vii.ELEMnull {
		game.Log().Warn("games.Sunrise: !resolve forfeit")
		game.Results = &vii.GameResults{
			Loser: t.HotSeat,
		}
	} else if seat := game.GetSeat(t.HotSeat); seat == nil {
		game.Log().Warn("sunrise: !resolve seat missing")
	} else {
		card := seat.Deck.Draw()
		seat.Hand[card.Id] = card
		seat.Reactivate()
		animate.Hand(game, seat)
		animate.BroadcastSeatUpdate(game, seat)
	}
	return t.Fork(game, Main(game, t))
}

func (event *SunriseEvent) Json(game *vii.Game, t *Timeline) js.Object {
	return js.Object{
		"gameid":   game,
		"timer":    t.Lifetime.Seconds(),
		"username": t.HotSeat,
	}
}
