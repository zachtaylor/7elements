package engine

import (
	"time"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"taylz.io/http/websocket"
	"taylz.io/log"
)

func randomfirstturn(seats *seat.List) string { return seats.Keys()[0] }

// Run the engine
func Run(syslog log.Writer, game *game.T) {
	tBeginning := time.Now()

	syslog = syslog.Add("GameID", game.ID())
	syslog.Info("start")

	// bootstrap

	game.State = game.NewState(phase.NewStart(randomfirstturn(game.Seats)))
	for _, name := range game.Seats.Keys() {
		seat := game.Seats.Get(name)
		seat.Writer.Write(websocket.NewMessage("/game", game.Data(seat)).EncodeToJSON())
	}
	phase.TryOnActivate(game)

	timer := time.NewTimer(game.State.Timer)

	for tStart := time.Now(); true; tStart = time.Now() { // event loop
		select { // read request chan or timeout
		case <-timer.C: // timeout
			game.Log().Warn("timeout")
			Resolve(game)
		case r, ok := <-game.RequestChan(): // player noise
			if !timer.Stop() {
				<-timer.C
			}
			if !ok {
				game.Log().Info("stopping")
				game.State = nil
				break
			}

			logger := game.Log().With(log.Fields{
				"Path":     r.URI,
				"Username": r.Username,
			})
			logger.Trace("received")

			td := time.Now().Sub(tStart)
			game.State.Timer -= td
			game.Log().With(log.Fields{
				"Elapsed":   td.Seconds(),
				"Remaining": game.State.Timer.Seconds(),
			}).Trace("updated timer")

			seat := game.Seats.Get(r.Username)
			if seat == nil {
				syslog.With(map[string]interface{}{
					"GameID": game.ID(),
				}).Warn("seat missing")
				continue
			}

			if rs := Request(game, seat, r.URI, r.Data); len(rs) > 0 {
				Stack(game, rs)
			}

			if len(game.State.Reacts) == game.Seats.Count() {
				logger.Info("resolve")
				Resolve(game)
			}
		}
		if game.State == nil {
			break
		}
		timer.Reset(game.State.Timer)
	} // event loop

	for _, seatName := range game.Seats.Keys() {
		seat := game.Seats.Get(seatName)
		seat.Message("/game", nil)
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
