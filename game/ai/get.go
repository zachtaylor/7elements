package ai

import "github.com/zachtaylor/7elements/element"

// getNewElement returns the ai preferred new element
func (ai *AI) getNewElement() element.T {
	devotion := element.Count{}
	for cardID := range ai.View.Self.T.Hand {
		card := ai.View.Game.Card(cardID)
		devotion.Add(card.T.Costs)
	}
	for e, stack := range ai.View.Self.T.Karma {
		devotion[e] -= len(stack)
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
	elements := ai.View.Self.T.Karma.Active()
	for cardID := range ai.View.Self.T.Hand {
		c := ai.View.Game.Card(cardID)
		if elements.Test(c.T.Costs) {
			hand = append(hand, cardID)
		}
	}
	return
}

// getPresentCanAttack returns a slice of gcid of cards ai can use to attack
func (ai *AI) getPresentCanAttack() (awake []string) {
	for tokenID := range ai.View.Self.T.Present {
		t := ai.View.Game.Token(tokenID)
		if t.T.Awake {
			awake = append(awake, tokenID)
		}
	}
	return
}
