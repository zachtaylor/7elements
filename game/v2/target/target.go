package target

import "errors"

var (
	ErrNotPresent = errors.New("not present")
	ErrNotPast    = errors.New("not past")

	ErrNotMine  = errors.New("not mine")
	ErrNotEnemy = errors.New("not enemy")

	// ErrNotMyPresent = errors.New("not my present")
	// ErrNotMyPast    = errors.New("not my past")

	ErrNotPlayer = errors.New("not player id")
	ErrNotCard   = errors.New("not card id")
	ErrNotToken  = errors.New("not token id")

	ErrNotBeing = errors.New("not being")
	ErrNotItem  = errors.New("not item")
	ErrNotSpell = errors.New("not spell")

	ErrBadPlayer = errors.New("bad player id")
)
