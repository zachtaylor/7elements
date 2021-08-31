package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

const vinespiritID = "vine-spirit"

func init() {
	script.Scripts[vinespiritID] = VineSpirit
}

func VineSpirit(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if token, ok := me.(*token.T); !ok || token == nil {
		err = ErrMeToken
	} else {
		token.Body.Attack++
		game.Seats.WriteSync(wsout.GameToken(token.Data()).EncodeToJSON())
	}
	return
}
