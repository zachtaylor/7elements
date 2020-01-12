package ai

import vii "github.com/zachtaylor/7elements"

// getNewElement returns the ai preferred new element
func (ai *AI) getNewElement() vii.Element {
	devotion := ai.Seat.Hand.Devotion()
	for e, stack := range ai.Seat.Elements {
		for _, ok := range stack {
			if ok {
				devotion[e]--
			}
		}
	}
	delete(devotion, 0)

	element, max := vii.ELEMnil, 0
	for e, count := range devotion {
		if count > max {
			max = count
			element = e
		}
	}

	if element == vii.ELEMnil {
		element = vii.ELEMyellow
	}
	return element
}

// getHandCanAfford returns a slice of gcid of cards ai can afford to play
func (ai *AI) getHandCanAfford() (hand []string) {
	elements := ai.Seat.Elements.GetActive()
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
