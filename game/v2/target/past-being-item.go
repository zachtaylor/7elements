package target

// func PastBeingItem(g *game.G, player *game.Player, arg interface{}) (*card.T, error) {
// 	if id, ok := arg.(string); !ok {
// 		return nil, errors.New("no id")
// 	} else if c := game.GetCard(id); c == nil {
// 		return nil, errors.New("no card: " + id)
// 	} else if s := game.Player(c.User); s == nil {
// 		return nil, errors.New("no seat: " + id)
// 	} else if c.Proto.Type != card.BodyType && c.Proto.Type != card.ItemType {
// 		return nil, errors.New("not being or item: " + c.Proto.Name)
// 	} else if !s.Past.Has(c.ID) {
// 		return nil, errors.New("not past: " + c.Proto.Name)
// 	} else {
// 		return c, nil
// 	}
// }
