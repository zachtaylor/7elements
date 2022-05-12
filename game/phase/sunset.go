package phase

import "github.com/zachtaylor/7elements/game/v2"

func NewSunset(priority game.Priority) game.Phaser {
	return &Sunset{
		PriorityContext: game.PriorityContext(priority),
	}
}

type Sunset struct {
	game.PriorityContext
}

func (r *Sunset) Type() string { return "sunset" }

// OnConnect implements game.OnConnectPhaser
func (r *Sunset) OnConnect(g *game.G, player *game.Player) {
	// if player == nil {
	// 	g.Chat("sunset", player.T.Username)
	// }
}

func (r *Sunset) JSON() map[string]any { return nil }
