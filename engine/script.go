package engine

import (
	"github.com/zachtaylor/7elements"
)

type ScriptFunc = func(*vii.Game, *vii.GameSeat, interface{}) vii.GameEvent

var Scripts = make(map[string]ScriptFunc)

func Script(game *vii.Game, seat *vii.GameSeat, power *vii.Power, target interface{}) vii.GameEvent {
	if script := Scripts[power.Script]; script == nil {
		game.Log().Add("Script", power.Script).Warn("script missing")
		return nil
	} else {
		return script(game, seat, target)
	}
}
