package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

var logtag = "scripts/"

func newlog(g *game.T, s *game.Seat, me interface{}, args []interface{}) *log.Entry {
	return g.Log().New().With(cast.JSON{
		"seat": s.Username,
		"me":   me,
		"args": args,
	})
}
