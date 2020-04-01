package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/power"
)

// Power invokes game script
func Power(g *game.T, s *game.Seat, p *power.T, me interface{}, args []interface{}) []game.Stater {
	log := g.Log().Add("Script", p.Script).Tag("trigger/power")
	if script := game.Scripts[p.Script]; script == nil {
		log.Error("script missing")
	} else if events, err := script(g, s, me, args); err != nil {
		log.Add("Error", err).Error()
	} else {
		log.Add("Events", events).Debug("stack")
		return events
	}
	return nil
}
