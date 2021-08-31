package trigger

// func NewPowerTrigger(game *game.T, seat *seat.T, token *token.T, power *power.T, arg string) game.Phaser {
// 	dirty := false
// 	if power.Costs.Total() > 0 {
// 		dirty = true
// 		seat.Karma.Deactivate(power.Costs)
// 	}
// 	if power.UsesTurn {
// 		token.IsAwake = false
// 		game.Seats.Write(wsout.GameTokenJSON(token.Data()))
// 	}
// 	if power.UsesLife {
// 		dirty = true
// 		delete(seat.Present, token.ID)
// 	}
// 	if dirty {
// 		game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
// 	}

// 	if power.Target == "self" {
// 		return phase.NewTrigger(seat.Username, token, power, token.ID)
// 	}
// 	return phase.NewTarget(seat.Username, power, token)
// }
