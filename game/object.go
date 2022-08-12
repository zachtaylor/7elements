package game

import (
	"time"

	"github.com/zachtaylor/7elements/card"
	"taylz.io/yas"
)

// type Object struct {
// 	id string
// 	Item any
// }

// Object is a managed game target
type Object[T Target] struct {
	id     string
	player string
	T      T
}

func NewObject[T Target](id, player string, t T) *Object[T] {
	return &Object[T]{
		id:     id,
		player: player,
		T:      t,
	}
}

func (obj Object[T]) ID() string { return obj.id }

func (obj Object[T]) Player() string { return obj.player }

// Target is a constraint for Object
type Target interface {
	*card.Prototype | PlayerContext | StateContext | TokenContext
}

type Card = Object[*card.Prototype]

type Player = Object[PlayerContext]

type StateContext struct {
	Phase Phaser
	React yas.Set[string]
	Stack *State
	Timer time.Duration
}

func NewStateContext(phase Phaser) StateContext {
	return StateContext{
		Phase: phase,
		React: yas.NewSet[string](),
	}
}

type State = Object[StateContext]

type Token = Object[TokenContext]
