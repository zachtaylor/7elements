package target

// func MyPresent(g *game.G, player *game.Player, arg interface{}) (*game.Token, error) {
// 	if id, ok := arg.(string); !ok {
// 		return nil, errors.New("no id")
// 	} else if token := g.Token(id); token == nil {
// 		return nil, errors.New("not token: " + id)
// 	} else if token.User != seat.Username {
// 		return nil, errors.New("not mine: " + id)
// 	} else if seat.Present[token.ID] == nil {
// 		return nil, errors.New("not present: " + id)
// 	} else {
// 		return token, nil
// 	}
// }
