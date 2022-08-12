package out

import (
	"github.com/zachtaylor/7elements/game"
	"taylz.io/http/websocket"
)

// func Error(player *game.Player, err error) {

// }

// func WritePlayer(g *game.G, player *game.Player) {

// }

func PrivateHand(player *game.Player) {
	player.T.Writer.Write(websocket.NewMessage("/game/hand", map[string]any{
		"ids": player.T.Hand.Slice(),
	}).ShouldMarshal())
}

func PrivateCard(player *game.Player, card *game.Card) {
	player.T.Writer.Write(CardMessage(card.ID(), card.T.Data()))
}

func ErrorMessage(id, error string) []byte {
	return websocket.NewMessage("/error", map[string]any{
		"id":   id,
		"data": error,
	}).ShouldMarshal()
}

func CardMessage(id string, cardData map[string]any) []byte {
	return websocket.NewMessage("/game/card", map[string]any{
		"id":   id,
		"data": cardData,
	}).ShouldMarshal()
}

func TokenMessage(id string, tokenData map[string]any) []byte {
	return websocket.NewMessage("/game/token", map[string]any{
		"id":   id,
		"data": tokenData,
	}).ShouldMarshal()
}

func StateMessage(id string, stateData map[string]any) []byte {
	return websocket.NewMessage("/game/state", map[string]any{
		"id":   id,
		"data": stateData,
	}).ShouldMarshal()
}

func PlayerMessage(id string, playerData map[string]any) []byte {
	return websocket.NewMessage("/game/player", map[string]any{
		"id":   id,
		"data": playerData,
	}).ShouldMarshal()
}

// func PrivateChoice(g *game.G, player *game.Player, text string, choices []map[string]any, data map[string]any) {

// }

// func WritePresent(g *game.G, player *game.Player) {

// }

// func WritePast(g *game.G, player *game.Player) {

// }

// func WriteToken(g *game.G, token *game.Token) {

// }

// func WriteReact(g *game.G, state *game.State, player *game.Player) {

// }
