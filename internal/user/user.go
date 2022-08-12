package user

import "taylz.io/http/user"

type Server interface {
	Users() Manager
}

type Manager = user.Manager

type Settings = user.Settings

var NewServiceHandler = user.NewServiceHandler
