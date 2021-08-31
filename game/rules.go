package game

import "time"

// Rules is a rules values sheet
type Rules struct {
	DeckMaximum  int
	DeckMinimum  int
	DeckDupes    int
	Timeout      time.Duration
	StartingLife int
	StartingHand int
}

func DefaultRules() Rules {
	return Rules{
		DeckMaximum:  -1,
		DeckMinimum:  21,
		DeckDupes:    3,
		StartingLife: 7,
		StartingHand: 3,
	}
}
