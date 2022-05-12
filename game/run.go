package game

// func Run(game *T) {
// 	tBeginning := time.Now()

// 	syslog := logger.Add("GameID", game.ID())
// 	syslog.Info("start")

// 	// bootstrap
// 	game.State = t.NewState(game, phase.NewStart(randomfirstturn(game.Seats)))
// 	for _, name := range game.Seats.Keys() {
// 		seat := g.Player(name)
// 		seat.Writer.WriteMessageData(wsout.Game(game.JSON(seat)))
// 	}
// 	phase.TryOnActivate(game)

// 	state.Run()
// }

// // Run the engine
// func (t *Machine) run(logger *log.T, game *game.T) {
// 	timer := time.NewTimer(game.State.Timer)
// 	for tStart := time.Now(); true; tStart = time.Now() { // event loop
// 		select { // read request chan or timeout
// 		case <-timer.C: // timeout
// 			game.Log().Warn("timeout")
// 			t.resolve(game)
// 		case r, ok := <-game.RequestChan(): // player noise
// 			if !timer.Stop() {
// 				<-timer.C
// 			}
// 			if !ok {
// 				game.Log().Info("stopping")
// 				game.State = nil
// 				break
// 			}

// 			td := time.Now().Sub(tStart)
// 			game.State.Timer -= td

// 			logger := game.Log().With(log.Fields{
// 				"Path":     r.URI,
// 				"Username": r.Username,
// 				"Timer":    int(game.State.Timer.Seconds()),
// 			})
// 			logger.Info("received")
// 			logger.Trace("td:", int(td.Seconds()))

// 			seat := g.Player(r.Username)
// 			if seat == nil {
// 				syslog.With(map[string]any{
// 					"GameID": game.ID(),
// 				}).Warn("seat missing")
// 				continue
// 			}

// 			if rs := Request(game, seat, r.URI, r.Data); len(rs) > 0 {
// 				t.stack(game, rs)
// 			}

// 			if len(game.State.Reacts) == game.Seats.Count() {
// 				logger.Info("resolve")
// 				t.resolve(game)
// 			}
// 		}
// 		if game.State == nil {
// 			break
// 		}
// 		timer.Reset(game.State.Timer)
// 	} // event loop

// 	for _, seatName := range game.Seats.Keys() {
// 		seat := g.Player(seatName)
// 		seat.Writer.WriteMessageData(wsout.Game(nil))
// 		seat.Writer = nil
// 	}

// 	game.Close()

// 	syslog.Add("Runtime", time.Since(tBeginning)).Info("done")
// }
