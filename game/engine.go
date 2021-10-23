package game

import (
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/power"
	"taylz.io/log"
)

type Engine interface {
	Run(syslog *log.T, game *T)
	NewEnding(game *T, results Resulter) Phaser
	NewTrigger(game *T, seat *seat.T, token *token.T, power *power.T) Phaser
	NewToken(*T, *seat.T, *token.T) []Phaser
	RemoveToken(*T, *token.T) []Phaser
	WakeToken(*T, *token.T) []Phaser
	SleepToken(*T, *token.T) []Phaser
	HealToken(*T, *token.T, int) []Phaser
	DamageToken(*T, *token.T, int) []Phaser
	HealSeat(*T, *seat.T, int) []Phaser
	DamageSeat(*T, *seat.T, int) []Phaser
	DrawCard(*T, *seat.T, int) []Phaser
	Script(game *T, seat *seat.T, script string, me interface{}, args []string) []Phaser
}
