package game

// SendData populates game data in multiple writes to stay under ws frame limit
// func (game *T) SendData(username string) {
// 	seat := g.Player(username)
// 	if seat == nil {
// 		return
// 	}

// 	seat.Writer.WriteMessageData(wsout.Game(game.JSON(seat)))
// 	seat.Writer.WriteMessageData(wsout.GameState(game.State.JSON()))

// 	seat.Writer.WriteMessageData(wsout.GameHand(seat.Hand.Keys()))
// 	for _, c := range seat.Hand {
// 		seat.Writer.WriteMessageData(wsout.GameCard(c.JSON()))
// 	}

// 	for _, name := range game.Seats.Keys() {
// 		s := g.Player(name)
// 		seat.Writer.WriteMessageData(wsout.GameSeat(s.JSON()))
// 		for _, t := range s.Present {
// 			seat.Writer.WriteMessageData(wsout.GameToken(t.JSON()))
// 		}
// 		for _, c := range s.Past {
// 			seat.Writer.WriteMessageData(wsout.GameCard(c.JSON()))
// 		}
// 	}
// }
