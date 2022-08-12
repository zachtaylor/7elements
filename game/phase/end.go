package phase

// import (
// 	"github.com/zachtaylor/7elements/game"
// 	"github.com/zachtaylor/7elements/game/seat"
// )

// func NewEnd(results game.Resulter) game.Phaser {
// 	return &End{
// 		Resulter: results,
// 	}
// }

// type End struct {
// 	game.Resulter
// }

// func (r *End) Name() string { return "end" }

// func (r *End) Seat() string { return "" }

// func (r *End) String() string { return r.Resulter.String() }

// // OnActivate implements game.OnActivatePhaserer
// func (r *End) OnActivate(game *game.G) []game.Phaser {
// 	if game.State.Phase.Name() == "end" {
// 		game.State.Stack = nil // rip L0L
// 	}
// 	return nil
// }
// func _endActivateStaterer(r *End) game.OnActivatePhaser {
// 	return r
// }

// // // // OnConnect implements game.OnConnectPhaser
// // func (r *End) OnConnect(*game.T, *seat.T) {
// // }

// func (r *End) OnDisconnect(g *game.G, player *game.Player) {
// 	g.Log().Add("Username", seat.Username).Debug("left")
// 	game.State.Reacts[seat.Username] = "left"
// }
// func (r *End) onDisconnectPhaser() game.OnDisconnectPhaser { return r }

// // // GetStack implements game.StackStater
// // func (r *End) GetStack(game *game.G) *game.State {
// // 	return nil
// // }

// // // Request implements Requester
// // func (r *End) Request(g*game.T, player *game.Player, json map[string]any) {
// // }

// func (r *End) GetNext(game *game.G) game.Phaser {
// 	return nil
// }

// // Finish implements game.OnFinishPhaserer
// // func (r *End) OnFinish(game *game.G) []game.Phaser { return nil }
// // func (r *End) onFinishPhaser() game.OnFinishPhaser { return r }

// func (r *End) JSON() map[string]any {
// 	return map[string]any{
// 		"winner": r.Winner,
// 		"loser":  r.Loser,
// 	}
// }
