package scripts

// const SymbiosisID = "symbiosis"

// func init() {
// 	script.Scripts[SymbiosisID] = Symbiosis
// }

// func Symbiosis(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
// 	rs = append(rs, phase.NewTarget(
// 		seat.Username,
// 		"being",
// 		"Target Being gains 1 Attack",
// 		func(val string) []game.Phaser {
// 			token, err := checktarget.PresentBeing(g, s, val)
// 			if err != nil {
// 				g.Log().Add("Error", err).Error()
// 			} else {
// 				g.Log().Add("Token", token.String()).Info()
// 				token.Body.Attack++
// 				g.Seats.WriteSync(wsout.GameToken(token.Data()).EncodeToJSON())
// 			}
// 			return nil
// 		},
// 	))
// 	return
// }
