package scripts

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

const WormholeID = "wormhole"

func init() {
	script.Scripts[WormholeID] = Wormhole
}

func Wormhole(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := checktarget.PresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else if token == nil {
		err = ErrBadTarget
	} else if owner := game.Seats.Get(token.Card.User); owner == nil {
		err = errors.New("owner missing: " + token.Card.User)
	} else if ctrlr := game.Seats.Get(token.User); ctrlr == nil {
		err = errors.New("ctrlr missing: " + token.Card.User)
	} else {
		game.Log().With(map[string]interface{}{
			"Target": token.ID,
			"Owner":  owner.Username,
			"Ctrlr":  ctrlr.Username,
		}).Info()
		rs = trigger.RemoveToken(game, token)

		if owner.Past.Has(token.Card.ID) {
			delete(owner.Past, token.Card.ID)
			owner.Hand[token.Card.ID] = token.Card
			game.Seats.Write(wsout.GameSeatJSON(owner.Data()))
			owner.Writer.Write(wsout.GameHand(owner.Hand.Keys()).EncodeToJSON())
		}
	}
	return
}
