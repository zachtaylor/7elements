package game

import "github.com/zachtaylor/7elements/game/seat"

// Script sets a func pointer for game code
//
// - me may be *game.Token or *game.Card
//
// - returns Phasers which create new States to stack
type Script = func(game *T, seat *seat.T, me interface{}, args []string) ([]Phaser, error)
