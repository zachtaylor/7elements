package game

import (
	"github.com/zachtaylor/7elements/element"
	"taylz.io/yas"
)

type PlayerContext struct {
	Username string
	Life     int
	Future   []string
	Karma    element.Karma
	Hand     yas.Set[string]
	Present  yas.Set[string]
	Past     yas.Set[string]
}
