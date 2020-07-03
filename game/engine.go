package game

import "github.com/zachtaylor/7elements/power"

// Engine is a plugin to execute game logic
type Engine interface {
	Run(*T)
	Start(seat string) Stater
	End(winner, loser string) Stater
	TriggerTokenEvent(seat *Seat, token *Token, trigger string) []Stater
	TriggerTokenPower(seat *Seat, token *Token, power *power.T, arg interface{}) []Stater
	Target(seat *Seat, target string, text string, finish func(val string) []Stater) Stater
}
