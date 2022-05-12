package scripts

// const SymbiosisID = "symbiosis"

// func init() {
// 	game.Scripts[SymbiosisID] = Symbiosis
// }

// func Symbiosis(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
// 	rs = append(rs, phase.NewTarget(
// 		seat.Username,
// 		"being",
// 		"Target Being gains 1 Attack",
// 		func(val string) []game.Phaser {
// 			token, err := target.PresentBeing(g, s, val)
// 			if err != nil {
// 				g.Log().Add("Error", err).Error()
// 			} else {
// 				g.Log().Add("Token", token.String()).Info()
// 				token.Body.Attack++
// 				g.Seats.WriteSync(wsout.GameToken(token.JSON()).EncodeToJSON())
// 			}
// 			return nil
// 		},
// 	))
// 	return
// }
