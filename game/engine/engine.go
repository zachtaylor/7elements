package engine

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

// Run the engine
func Run(g *game.T) {
	var tStart cast.Time

	enginelog := g.Log().With(cast.JSON{
		"GameID": g.ID(),
	}).Tag("engine")
	enginelog.Info("start")

	// bootstrap
	g.State.Activate(g)
	timer := cast.NewTimer(g.State.Timer)

loop: // nested break
	for { // event loop
		tStart = cast.Now()
		select { // read request chan or timeout
		case <-timer.C: // timeout
			g.Log().Tag("engine/timeout")
			g.ResolveState()
		case r, ok := <-g.Monitor(): // player noise
			if !timer.Stop() {
				<-timer.C
			}
			g.State.Timer -= cast.Now().Sub(tStart)

			if !ok {
				g.Log().Info("stopping")
				break loop // nested break
			}

			log := g.Log().With(cast.JSON{
				"Path":     r.URI,
				"Username": r.Username,
				"State":    g.State.Name(),
			}).Tag("engine/request")

			if states := request(g, g.GetSeat(r.Username), r.URI, r.Data); len(states) < 1 {
				log.Debug("no stack")
			} else {
				log.Copy().Add("Stack", states).Debug("stacking")
				g.Stack(states)
			}

			if len(g.State.Reacts) == len(g.Seats) {
				log.Info("resolve")
				g.ResolveState()
			} else if g.State.Timer < cast.Second {
				log.Warn("timeout")
				g.ResolveState()
			} else {
				log.Debug()
			}
		}
		timer.Reset(g.State.Timer)
	} // event loop

	enginelog.With(cast.JSON{
		"Time":  cast.Now().Sub(tStart),
		"State": g.State,
	}).Info("end")
	for _, seat := range g.Seats {
		if seat.Player != nil {
			seat.Player.Send("/game", nil)
		}
		seat.Player = nil
	}
}
