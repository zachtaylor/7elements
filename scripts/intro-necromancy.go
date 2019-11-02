package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event/end"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const IntroToNecromancyID = "intro-necromancy"

func init() {
	game.Scripts[IntroToNecromancyID] = IntroToNecromancy
}

func IntroToNecromancy(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + IntroToNecromancyID)

	card, err := target.MyPastBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	seat.Present[card.Id] = card
	card.Body.Health = 1
	// TODO return trigger.DamageSeat()
	seat.Life--
	if seat.Life < 1 {
		return []game.Event{end.New(g.GetOpponentSeat(seat.Username).Username, seat.Username)}
	}
	return nil
}
