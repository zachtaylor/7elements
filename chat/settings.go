package chat

import "taylz.io/http/user"

type Settings struct {
	Users  *user.Manager
	Keygen func() string
}
