package engine

import (
	"time"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/out"
	"github.com/zachtaylor/7elements/game/phase"
	"taylz.io/http/websocket"
	"taylz.io/log"
)

type runner struct{}

func NewRunner() game.Runner                    { return runner{} }
func (runner) Run(game *game.G, log log.Writer) { Run(game, log) }

func Run(g *game.G, syslog log.Writer) {
	// bootstrap
	firstTurn := g.Players()[0]
	state := g.NewState(firstTurn, game.NewStateContext(phase.NewStart(g.NewPriority(firstTurn))))

	// for _, playerID := range g.Players() {
	// player := g.Player(playerID)
	// player.WriteMessageData(wsout.Game(game.JSON(seat)))
	// }
	game.OnActivatePhase(g, state.T.Phase)

	run(g, state, syslog)
}

// Run the engine
func run(g *game.G, state *game.State, syslog log.Writer) {
	timer := time.NewTimer(g.Rules().Timeout)
	tStart := time.Now()
	for tLast := tStart; true; tLast = time.Now() { // event loop
		select { // read request chan or timeout
		case <-timer.C: // timeout
			g.Log().Warn("timeout")
			state = resolve(g, state)
		case r, ok := <-g.RequestChan(): // player noise
			if !timer.Stop() {
				<-timer.C
			}
			if !ok {
				g.Log().Info("stopping")
				state = nil
				break
			}

			td := time.Now().Sub(tLast)
			state.T.Timer -= td

			logger := g.Log().With(log.Fields{
				"Path":     r.URI,
				"Username": r.Username,
				"Timer":    int(state.T.Timer.Seconds()),
				"dT":       int(td.Seconds()),
			})
			logger.Info("received")

			player := g.Player(r.Username)
			if player == nil {
				syslog.With(map[string]any{
					"GameID": g.ID(),
				}).Warn("seat missing")
				continue
			}

			if rs := Request(g, state, player, r.URI, r.Data); len(rs) > 0 {
				state = stack(g, state, rs)
			} else if len(state.T.React) == g.PlayerCount() {
				logger.Info("resolve")
				state = resolve(g, state)
			}
		}
		if state == nil {
			break
		}
		updates := g.ReadUpdates()
		for _, objID := range updates {
			if obj := g.Object(objID); obj == nil {

			} else if token, ok := obj.(*game.Token); ok {
				g.Write(out.TokenMessage(objID, token.T.Data()))
			} else if card, ok := obj.(*game.Card); ok {
				g.Write(out.CardMessage(objID, card.T.Data()))
			} else if state, ok := obj.(*game.State); ok {
				g.Write(out.StateMessage(objID, state.T.Phase.JSON()))
			} else if player, ok := obj.(*game.Player); ok {
				g.Write(out.PlayerMessage(objID, player.T.Data()))
			}
			// push update data
		}
		timer.Reset(state.T.Timer)
	} // event loop

	g.Write(websocket.NewMessage("/game", nil).ShouldMarshal())

	for _, playerID := range g.Players() {
		player := g.Player(playerID)
		player.T.Writer = nil
	}

	g.Close()

	syslog.Add("Runtime", time.Since(tStart)).Info("done")
}
