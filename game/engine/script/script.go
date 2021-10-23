package script

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
)

// Scripts is the injection point
var Scripts = make(map[string]game.Script)

// Run uses the Power, returns additional phasers
func Run(game *game.T, seat *seat.T, scriptName string, me interface{}, args []string) []game.Phaser {
	if script := Scripts[scriptName]; script == nil {
	} else if stack, err := script(game, seat, me, args); err != nil {
		game.Log().With(map[string]interface{}{
			"GameID": game.ID(),
			"Script": scriptName,
			"Me":     me,
			"Args":   args,
			"Error":  err,
		}).Error("script failure")
	} else {
		return stack
	}
	return nil
}
