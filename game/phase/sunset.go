package phase

import (
	"github.com/zachtaylor/7elements/game"
	"taylz.io/yas"
)

func NewSunset(priority game.Priority) game.Phaser {
	return &Sunset{
		PriorityContext: game.PriorityContext(priority),
	}
}

type Sunset struct {
	game.PriorityContext
}

func (*Sunset) Type() string { return "sunset" }

func index[T comparable](slice []T, t T) int {
	for i, v := range slice {
		if t == v {
			return i
		}
	}
	return -1
}

func (r *Sunset) Next() game.Phaser {
	prio := r.Priority()
	cur, next := yas.Shift(prio)
	if index(next, cur) < 0 {
		// reuse buffer
		copy(prio, prio[1:])
		prio[len(prio)-1] = cur
	} else {
		prio = prio[:len(prio)-1]
	}
	return NewSunrise(prio)
}

// OnConnect implements game.OnConnectPhaser
func (r *Sunset) OnConnect(g *game.G, player *game.Player) {
	// if player == nil {
	// 	g.Chat("sunset", player.T.Username)
	// }
}

func (r *Sunset) JSON() map[string]any { return nil }
