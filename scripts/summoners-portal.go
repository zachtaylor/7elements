package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

const summonersportalID = "summoners-portal"

func init() {
	script.Scripts[summonersportalID] = SummonersPortal
}

func SummonersPortal(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else if c := seat.Deck.Draw(); c == nil {
		err = ErrFutureEmpty
	} else if c.Proto.Type == card.BodyType || c.Proto.Type == card.ItemType {
		rs = append(rs, phase.NewPlay(seat.Username, c, ""))
	} else {
		seat.Past[c.ID] = c
		seat.Writer.Write(wsout.ErrorJSON("Summoners Portal", "Next card was "+c.Proto.Name))
		game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
	}
	return
}
