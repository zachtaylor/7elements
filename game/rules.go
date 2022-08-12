package game

import "time"

// Rules is a rules values sheet
type Rules struct {
	DeckMax    int
	DeckMin    int
	DeckCopy   int
	Timeout    time.Duration
	PlayerLife int
	PlayerHand int
}

func DefaultRules() Rules {
	return Rules{
		DeckMax:    -1,
		DeckMin:    21,
		DeckCopy:   3,
		Timeout:    60 * time.Second,
		PlayerLife: 7,
		PlayerHand: 4,
	}
}
