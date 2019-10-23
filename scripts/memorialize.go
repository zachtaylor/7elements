package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const MemorializeID = "memorialize"

func init() {
	game.Scripts[MemorializeID] = Memorialize
}

func Memorialize(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
	}).Tag(logtag + MemorializeID)
	card, err := target.MyPastBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	c := game.NewCard(card.Card)
	c.Username = seat.Username
	g.RegisterCard(c)
	seat.Hand[c.Id] = c
	g.SendSeatUpdate(seat)
	seat.SendHandUpdate()
	return nil
}
