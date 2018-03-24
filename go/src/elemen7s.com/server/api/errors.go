package api

import (
	"errors"
)

var ErrSessionRequired = errors.New("session required")
var ErrGameIdRequired = errors.New("game id required")
var ErrGameMissing = errors.New("game missing")
var ErrInsufficientFunds = errors.New("insufficient funds")
