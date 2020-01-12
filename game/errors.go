package game

import "errors"

// ErrBadTarget means the target is bad for unknown reasons
var ErrBadTarget = errors.New("bad target")

// ErrNotImplemented is a recursive definition
var ErrNotImplemented = errors.New("not implemented")

// ErrFutureEmpty is future is empty
var ErrFutureEmpty = errors.New("future is empty")

// ErrMeCard expected me to be type *game.Card
var ErrMeCard = errors.New("script host must be a card")

// ErrMeToken expected me to be type *game.Token
var ErrMeToken = errors.New("script host must be a token")

var ErrMeNil = errors.New("script host cannot be nil")

// var ErrNotEnoughKarma = errors.New("not enough karma")

var ErrNoSeat = errors.New("no seat")

var ErrNoTarget = errors.New("no target")

var ErrNotBeing = errors.New("not being")

var ErrNotItem = errors.New("not being")

var ErrNotPresent = errors.New("not present")

var ErrNotPast = errors.New("not past")
