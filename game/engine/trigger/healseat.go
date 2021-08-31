package trigger

import (
	"strconv"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

func HealSeat(game *game.T, seat *seat.T, n int) []game.Phaser {
	seat.Life += n
	game.Chat(seat.Username, "gain "+strconv.FormatInt(int64(n), 10)+" Life")
	game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
	return nil // todo
}
