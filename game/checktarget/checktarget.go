package checktarget

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

func IsValid(game *game.T, seat *seat.T, target, val string) bool {
	return true // todo
}

func IsToken(me interface{}) bool {
	if token, _ := me.(*token.T); token != nil {
		return true
	}
	return false
}

func IsCard(me interface{}) bool {
	if card, _ := me.(*card.T); card != nil {
		return true
	}
	return false
}
