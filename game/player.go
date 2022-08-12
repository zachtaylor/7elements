package game

import (
	"github.com/zachtaylor/7elements/element"
	"taylz.io/yas"
)

type PlayerContext struct {
	Writer  Writer
	Life    int
	Future  []string
	Karma   element.Karma
	Hand    yas.Set[string]
	Present yas.Set[string]
	Past    yas.Set[string]
}

func NewPlayerContext(w Writer, life int) PlayerContext {
	return PlayerContext{
		Writer:  w,
		Life:    life,
		Karma:   element.Karma{},
		Hand:    yas.NewSet[string](),
		Present: yas.NewSet[string](),
		Past:    yas.NewSet[string](),
	}
}

func (p *PlayerContext) Data() map[string]any {
	return map[string]any{
		"life":    p.Life,
		"future":  len(p.Future),
		"karma":   p.Karma.JSON(),
		"hand":    len(p.Hand),
		"present": p.Present.Slice(),
		"past":    p.Past.Slice(),
	}
}
