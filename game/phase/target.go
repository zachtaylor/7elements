package phase

// import (
// 	"github.com/zachtaylor/7elements/game/v2"
// 	"github.com/zachtaylor/7elements/game/v2/seat"
// 	"github.com/zachtaylor/7elements/power"
// )

// func NewTarget(seat string, p *power.T, me interface{}) *Target {
// 	return &Target{
// 		game.PriorityContext: R(seat),
// 		this:                 me,
// 		power:                p,
// 	}
// }

// type Target struct {
// 	R
// 	this   interface{}
// 	power  *power.T
// 	answer string
// }

// func (r *Target) isPhaser() game.Phaser { return r }

// func (r *Target) Name() string {
// 	return "target"
// }

// func (r *Target) String() string {
// 	return "target (" + r.Seat() + ":" + r.power.Target + ")"
// }

// func (r *Target) JSON() map[string]any {
// 	return map[string]any{
// 		"helper": r.power.Target,
// 		"text":   r.power.Text,
// 	}
// }

// func (r *Target) GetNext(game *game.G) game.Phaser {
// 	return nil
// }

// // OnActivate implements game.OnActivatePhaser
// func (r *Target) OnActivate(game *game.G) []game.Phaser {
// 	// go game.Chat(r.Seat(), r.power.Text)
// 	return nil
// }
// func (r *Target) onActivatePhaser() game.OnActivatePhaser { return r }

// // OnConnect implements game.OnConnectPhaser
// func (r *Target) OnConnect(g *game.G, player *game.Player) {
// 	// if seat == nil {
// 	// go game.Chat("target", r.Seat())
// 	// }
// }
// func (r *Target) onConnectPhaser() game.OnConnectPhaser { return r }

// // Request implements Requester
// func (r *Target) OnRequest(g *game.G, player *game.Player, json map[string]any) {
// 	if seat.Username != r.Seat() {
// 		g.Log().With(map[string]any{
// 			"seat": seat,
// 			"json": json,
// 		}).Warn("engine/target: receive")
// 		return
// 	}

// 	if r.answer, _ = json["choice"].(string); r.answer != "" {
// 		game.State.Reacts[seat.Username] = r.answer
// 	}
// }
// func (r *Target) onRequestPhaser() game.OnRequestPhaser { return r }

// // Finish implements game.OnFinishPhaser
// func (r *Target) OnFinish(game *game.G) []game.Phaser {
// 	return game.Engine().Script(game, g.Player(r.Seat()), r.power.Script, r.this, []string{game.State.Reacts[r.Seat()]})
// }
// func (r *Target) onFinishPhaser() game.OnFinishPhaser { return r }
