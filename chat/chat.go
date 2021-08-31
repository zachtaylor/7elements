package chat

import "errors"

var (
	ErrNoRoom        = errors.New("no room")
	ErrUserNotInRoom = errors.New("user not in room")
)
