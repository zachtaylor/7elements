package trigger

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

// Power invokes game script
func Power(g *game.T, s *game.Seat, p *vii.Power, me interface{}, args []interface{}) []game.Stater {
	// return []game.Stater{}
	log := g.Log().Add("Script", p.Script).Tag("trigger/power")
	if script := game.Scripts[p.Script]; script == nil {
		log.Add("Error", game.ErrNotImplemented).Error()
	} else if events, err := script(g, s, me, args); err != nil {
		log.Add("Error", err).Error()
	} else {
		log.Add("Events", events).Debug("stack")
		return events
	}
	return nil
}
