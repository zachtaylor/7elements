package state

// func TryOnRequest(g *game.T, seat *seat.T, json map[string]any) {
// 	if requester, ok := g.State.Phase.(game.OnRequestPhaser); ok {
// 		requester.OnRequest(g, seat, json)
// 	}
// }

// func TryOnFinish(g *game.T) (rs []game.Phaser) {
// 	if finisher, _ := g.State.Phase.(game.OnFinishPhaser); finisher != nil {
// 		rs = finisher.OnFinish(g)
// 	}
// 	if rs == nil {
// 		rs = []game.Phaser{}
// 	}
// 	return
// }
