package scripts

import "errors"

var (
	ErrMeCard      = errors.New("script host must be a card")
	ErrMeToken     = errors.New("script host must be a token")
	ErrNoTarget    = errors.New("no target")
	ErrBadTarget   = errors.New("bad target")
	ErrFutureEmpty = errors.New("future is empty")
)
