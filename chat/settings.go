package chat

import "taylz.io/http/user"

type Settings struct {
	Users  *user.Manager
	Keygen func() string
}

func NewSettings(users *user.Manager, keygen func() string) Settings {
	return Settings{
		Users:  users,
		Keygen: keygen,
	}
}
