package engine

import (
	"time"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/log"
)

func randomfirstturn(seats *seat.List) string { return seats.Keys()[0] }

// Run the engine
func (t *T) run(logger *log.T, game *game.T) {
	tBeginning := time.Now()

	syslog := logger.Add("GameID", game.ID())
	syslog.Info("start")

	// bootstrap
	game.State = t.NewState(game, phase.NewStart(randomfirstturn(game.Seats)))
	for _, name := range game.Seats.Keys() {
		seat := game.Seats.Get(name)
		seat.Writer.Write(wsout.Game(game.Data(seat)))
	}
	phase.TryOnActivate(game)

	timer := time.NewTimer(game.State.Timer)

	for tStart := time.Now(); true; tStart = time.Now() { // event loop
		select { // read request chan or timeout
		case <-timer.C: // timeout
			game.Log().Warn("timeout")
			t.resolve(game)
		case r, ok := <-game.RequestChan(): // player noise
			if !timer.Stop() {
				<-timer.C
			}
			if !ok {
				game.Log().Info("stopping")
				game.State = nil
				break
			}

			td := time.Now().Sub(tStart)
			game.State.Timer -= td

			logger := game.Log().With(log.Fields{
				"Path":     r.URI,
				"Username": r.Username,
				"Timer":    int(game.State.Timer.Seconds()),
			})
			logger.Info("received")
			logger.Trace("td:", int(td.Seconds()))

			seat := game.Seats.Get(r.Username)
			if seat == nil {
				syslog.With(map[string]interface{}{
					"GameID": game.ID(),
				}).Warn("seat missing")
				continue
			}

			if rs := Request(game, seat, r.URI, r.Data); len(rs) > 0 {
				t.stack(game, rs)
			}

			if len(game.State.Reacts) == game.Seats.Count() {
				logger.Info("resolve")
				t.resolve(game)
			}
		}
		if game.State == nil {
			break
		}
		timer.Reset(game.State.Timer)
	} // event loop

	for _, seatName := range game.Seats.Keys() {
		seat := game.Seats.Get(seatName)
		seat.Writer.Write(wsout.Game(nil))
		seat.Writer = nil
	}

	game.Close()

	syslog.Add("Runtime", time.Since(tBeginning)).Info("done")
}

// func end(syslog *log.T) {
// if username := game.Winner; username == "" {
// 	log.Warn("winner missing")
// } else if username == "A.I" {
// 	// skip
// } else if seat := game.GetSeat(username); seat == nil {
// 	log.Warn("winning seat missing")
// } else if account, err := accounts.Get(game.Runtime.DB, username); err != nil {
// 	log.Copy().Add("Error", err).Error("account missing")
// } else {
// 	account.Coins += 2
// 	if err = accounts.UpdateCoins(g.Runtime.DB, account); err != nil {
// 		log.Copy().Add("Error", err).Add("Username", username).Error("account service error")
// 	} else {
// 		out.Account(seat.Player)
// 	}
// }

// if username := r.Loser; username == "" {
// 	log.Warn("loser missing")
// } else if username == "A.I." {
// 	// skip
// } else if seat := g.GetSeat(username); seat == nil {
// 	log.Warn("winning seat missing")
// } else if account, err := accounts.Get(g.Runtime.DB, username); err != nil {
// 	log.Add("Error", err).Error("account missing")
// } else if r.Winner == "" {
// 	log.Add("Error", "Forfeit!").Warn("no pity coins")
// } else {
// 	account.Coins++
// 	if err = g.Runtime.Accounts.UpdateCoins(account); err != nil {
// 		log.Add("Error", err).Error("account service error")
// 	}
// 	out.Account(seat.Player)
// }
// g.Close()
// }
