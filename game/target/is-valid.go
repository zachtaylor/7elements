package target

import "github.com/zachtaylor/7elements/game"

func IsValid(g *game.G, player *game.Player, target string, objectID string) bool {
	switch target {
	case "being":
		if t := g.Token(objectID); t != nil {
			return t.T.Body != nil
		}
	case "item":
		if t := g.Token(objectID); t != nil {
			return t.T.Body == nil
		}
	case "play":
		if state := g.State(objectID); state != nil {
			return state.T.Phase.Type() == "play"
		}
	}
	return false
}
