package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

type ScriptFunc = func(*game.T, *game.Seat, interface{}) game.Event

var Scripts = make(map[string]ScriptFunc)

func Script(game *game.T, seat *game.Seat, power *vii.Power, target interface{}) game.Event {
	if script := Scripts[power.Script]; script == nil {
		game.Log().Add("Script", power.Script).Warn("script missing")
		return nil
	} else {
		return script(game, seat, target)
	}
}
