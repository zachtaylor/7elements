package scripts

// const HandrailsID = "handrails"

// func init() {
// 	game.Scripts[HandrailsID] = Handrails
// }

// func Handrails(g *game.T, seat *game.Seat, arg interface{}) []game.Stater {
// 	log := g.Log().With(log.Fields{
// 		"Target":   arg,
// 		"Username": seat.Username,
// 	}).Tag(logtag + HandrailsID)
// 	token, err := target.PresentBeing(g, seat, arg)
// 	if err != nil {
// 		log.Add("Error", err).Error()
// 		return nil
// 	}
// 	log.Info()
// 	token.IsAwake = true
// 	update.Token(g, token)
// 	return nil
// }
