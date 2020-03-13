package ai

import "github.com/zachtaylor/7elements/element"

// getNewElement returns the ai preferred new element
func (ai *AI) getNewElement() element.T {
	devotion := ai.Seat.Hand.Devotion()
	for e, stack := range ai.Seat.Karma {
		for _, ok := range stack {
			if ok {
				devotion[e]--
			}
		}
	}
	delete(devotion, 0)

	e, max := element.Nil, 0
	for el, count := range devotion {
		if count > max {
			max = count
			e = el
		}
	}

	if e == element.Nil {
		e = element.Yellow
	}
	return e
}

// getHandCanAfford returns a slice of gcid of cards ai can afford to play
func (ai *AI) getHandCanAfford() (hand []string) {
	elements := ai.Seat.Karma.Active()
	for _, c := range ai.Seat.Hand {
		if elements.Test(c.Card.Costs) {
			hand = append(hand, c.ID)
		}
	}
	return
}

// getPresentCanAttack returns a slice of gcid of cards ai can use to attack
func (ai *AI) getPresentCanAttack() (awake []string) {
	for _, c := range ai.Seat.Present {
		if c.IsAwake {
			awake = append(awake, c.ID)
		}
	}
	return
}
