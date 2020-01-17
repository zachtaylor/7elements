package engine

import (
	"github.com/zachtaylor/7elements/game"
	pkg_request "github.com/zachtaylor/7elements/game/engine/request"
	"ztaylor.me/cast"
)

func request(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Stater {
	if g.State.R.Name() == "main" && g.State.R.Seat() == seat.Username {
		return pkg_request.InMain(g, seat, uri, json)
	}
	return pkg_request.InResponse(g, seat, uri, json)
}
