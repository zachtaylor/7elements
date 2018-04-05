package games

type Stack struct {
	*Event
}

func StackEvent(e *Event) *Stack {
	return &Stack{e}
}

func (s *Stack) OnResolve(e *Event, g *Game) {
	g.Active = s.Event
	g.Active.Activate(g)
}
