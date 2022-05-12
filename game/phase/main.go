package phase

import (
	"github.com/zachtaylor/7elements/game/v2"
)

func NewMain(priority game.Priority) game.Phaser {
	return &Main{
		PriorityContext: game.PriorityContext(priority),
	}
}

type Main struct{ game.PriorityContext }

func (r *Main) Type() string { return "main" }

// OnConnect implements game.OnConnectPhaser
func (r *Main) OnConnect(g *game.G, player *game.Player) {
	if player == nil {
		// go game.Chat("main", r.Seat())
	}
}

// func (r *Main) GetNext(game *game.G) game.Phaser { return NewSunset(r.Seat()) }

func (r *Main) JSON() map[string]any { return nil }
