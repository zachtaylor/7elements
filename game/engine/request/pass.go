package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/out"
)

func Pass(game *game.G, state *game.State, player *game.Player, json map[string]any) {
	log := game.Log().With(map[string]any{
		"State":    game.State,
		"Username": player.T.Writer.Name(),
	})
	if pass, _ := json["pass"].(string); pass == "" {
		log.Warn("target missing")
	} else if pass != state.ID() {
		log.Add("PassID", pass).Warn("target mismatch")
	} else if state.T.React.Has(player.T.Writer.Name()) {
		player.T.Writer.Write(out.ErrorMessage("pass", "response already recorded"))
	} else {
		state.T.React.Add(player.T.Writer.Name())
		game.MarkUpdate(state.ID())
	}
}
