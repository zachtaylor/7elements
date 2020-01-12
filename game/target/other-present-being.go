package target

import (
	"errors"

	"ztaylor.me/cast"

	"github.com/zachtaylor/7elements/game"
)

func OtherPresentBeing(g *game.T, seat *game.Seat, token *game.Token, arg interface{}) (*game.Token, error) {
	if tid, ok := arg.(string); !ok {
		return nil, errors.New("no tid")
	} else if t := g.Objects[tid]; t == nil {
		return nil, errors.New("no token: " + tid)
	} else if target, ok := t.(*game.Token); !ok {
		return nil, errors.New("not token: " + cast.String(t))
	} else if s := g.GetSeat(target.Username); s == nil {
		return nil, errors.New("no seat")
	} else if !s.HasPresent(target.ID) {
		return nil, errors.New("not present")
	} else if target.Body == nil {
		return nil, errors.New("not being")
	} else if target.ID == token.ID {
		return nil, errors.New("not other")
	} else {
		return target, nil
	}
}
