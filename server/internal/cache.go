package internal

import (
	"taylz.io/http/session"
	"taylz.io/yas"
)

type Cache struct {
	Sessions yas.Cache[*session.T]
}
