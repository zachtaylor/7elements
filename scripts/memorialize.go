package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

const memorializeID = "memorialize"

func init() {
	script.Scripts[memorializeID] = Memorialize
}

func Memorialize(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if targetCard, _err := checktarget.MyPastBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else {
		card := card.New(targetCard.Proto, seat.Username)
		game.RegisterCard(card)
		seat.Hand[card.ID] = card
		game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
		seat.Writer.Write(wsout.GameHand(seat.Hand.Keys()).EncodeToJSON())
	}
	return
}
