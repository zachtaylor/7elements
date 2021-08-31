package script

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/power"
)

// Scripts is the injection point
var Scripts = make(map[string]game.Script)

// Run uses the Power, returns additional phasers
func Run(game *game.T, seat *seat.T, p *power.T, me interface{}, args []string) []game.Phaser {
	log := game.Log().Add("Script", p.Script)
	if script := Scripts[p.Script]; script == nil {
		log.Error("script missing")
	} else if events, err := script(game, seat, me, args); err != nil {
		log.Add("Error", err).Error()
	} else {
		log.Add("Events", events).Debug("stack")
		return events
	}
	return nil
}
