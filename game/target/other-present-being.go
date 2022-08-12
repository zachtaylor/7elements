package target

// func OtherPresentBeing(g *game.G, player *game.Player, me *game.Token, id string) (*game.Token, error) {
// 	if token := g.Token(id); token == nil {
// 		return nil, errors.New("no token: " + id)
// 	} else if s := game.Player(token.User); s == nil {
// 		return nil, errors.New("no seat: " + token.Card.Proto.Name)
// 	} else if !s.Present.Has(token.ID) {
// 		return nil, errors.New("not present: " + token.Card.Proto.Name)
// 	} else if token.Body == nil {
// 		return nil, errors.New("not being: " + token.Card.Proto.Name)
// 	} else if me.ID == token.ID {
// 		return nil, errors.New("is me: " + token.Card.Proto.Name)
// 	} else {
// 		return token, nil
// 	}
// }
