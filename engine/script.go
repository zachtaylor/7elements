package engine

import (
	"github.com/zachtaylor/7tcg"
)

type ScriptFunc = func(*vii.Game, *Timeline, *vii.GameSeat, interface{}) *Timeline

var Scripts = make(map[string]ScriptFunc)

func Script(game *vii.Game, t *Timeline, seat *vii.GameSeat, power *vii.Power, target interface{}) *Timeline {
	if script := Scripts[power.Script]; script == nil {
		game.Log().Add("Script", power.Script).Warn("script missing")
		return nil
	} else {
		return script(game, t, seat, target)
	}
}
