package out

import (
	"github.com/zachtaylor/7elements/game"
	"taylz.io/http/websocket"
)

// SendData populates game data in multiple writes to stay under ws frame limit
func SendData(g *game.G, state *game.State, player *game.Player) {

	player.T.Writer.Write(websocket.NewMessage("/game", map[string]any{
		"id": g.ID(),
		"data": map[string]any{
			"playerID": player.ID(),
			"players":  g.Players(),
		},
	}).ShouldMarshal())
	player.T.Writer.Write(StateMessage(state.ID(), state.T.Phase.JSON()))

	// player.T.Writer.Write(PrivateHand(seat.Hand.Keys()))
	for cardID := range player.T.Hand {
		PrivateCard(player, g.Card(cardID))
	}
	PrivateHand(player)

	for _, playerID := range g.Players() {
		p := g.Player(playerID)
		player.T.Writer.Write(PlayerMessage(playerID, p.T.Data()))
		for tokenID := range p.T.Present {
			t := g.Token(tokenID)
			player.T.Writer.Write(TokenMessage(tokenID, t.T.Data()))
		}
		for cardID := range p.T.Past {
			c := g.Card(cardID)
			player.T.Writer.Write(CardMessage(cardID, c.T.Data()))
		}
	}
}
