package games

type TurnClock struct {
	Username string
	Next     *TurnClock
}

func BuildTurnClock(names []string) *TurnClock {
	root := &TurnClock{}
	step := root
	for _, name := range names {
		step.Next = &TurnClock{
			Username: name,
		}
		step = step.Next
	}
	step.Next = root.Next
	return root.Next
}
