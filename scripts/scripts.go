package scripts

import (
	"errors"

	"ztaylor.me/cast"
)

var ErrMeCard = cast.NewError(nil, "script host must be a card")
var ErrMeToken = cast.NewError(nil, "script host must be a token")
var ErrNoTarget = cast.NewError(nil, "no target")
var ErrBadTarget = errors.New("bad target")
var ErrFutureEmpty = errors.New("future is empty")
