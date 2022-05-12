package phase

// import (
// 	"github.com/zachtaylor/7elements/game/v2"
// 	"github.com/zachtaylor/7elements/power"
// )

// func NewTrigger(seat string, token *game.Token, p *power.T, target string) game.Phaser {
// 	return &Trigger{
// 		game.PriorityContext: R(seat),
// 		Token:                token,
// 		Power:                p,
// 		Target:               target,
// 	}
// }

// type Trigger struct {
// 	R
// 	Token  *game.Token
// 	Power  *power.T
// 	Target string
// }

// func (r *Trigger) Name() string {
// 	return "trigger"
// }

// func (r *Trigger) String() string {
// 	return "trigger (" + r.Seat() + ":" + r.Token.ID + ")"
// }

// func (r *Trigger) GetNext(game *game.G) game.Phaser {
// 	return nil
// }

// func (r *Trigger) JSON() map[string]any {
// 	return map[string]any{
// 		"token":  r.Token.JSON(),
// 		"power":  r.Power.JSON(),
// 		"target": r.Target,
// 	}
// }

// // // OnActivate implements game.OnActivatePhaser
// // func (r *Trigger) OnActivate(game *T) []Phaser {
// // 	return nil
// // }
// // func (r *Trigger) activateEventer() game.OnActivatePhaser {
// // 	return r
// // }

// // OnConnect implements OnConnectPhaser
// func (r *Trigger) OnConnect(g *game.G, player *game.Player) {
// 	// if seat == nil {
// 	// go game.Chat(game.State.Phase.Seat(), "Trigger "+r.Token.Card.Proto.Name)
// 	// }
// }
// func (r *Trigger) onConnectPhaser() game.OnConnectPhaser { return r }

// // Finish implements OnFinishPhaser
// func (r *Trigger) OnFinish(game *game.G) []game.Phaser {
// 	seat := g.Player(r.Seat())
// 	g.Log().With(map[string]any{
// 		"Username": seat.Username,
// 		"Token":    r.Token,
// 		"Stack":    game.State.Stack,
// 	}).Debug("engine/trigger: finish")
// 	return game.Engine().Script(game, seat, r.Power.Script, r.Token, []string{r.Target})
// }
// func (r *Trigger) onFinishPhaser() game.OnFinishPhaser { return r }

// // // OnConnect implements OnConnectPhaser
// // func (r *Trigger) OnConnect(*T, *seat.T) {
// // }

// // // Request implements Requester
// // func (r *Trigger) Request(g*T, player *game.Player, json js.Object) {
// // }
