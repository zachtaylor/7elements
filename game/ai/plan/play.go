package plan

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai/aim"
	"github.com/zachtaylor/7elements/game/ai/view"
	"github.com/zachtaylor/7elements/power"
)

// Play is a plan to make a play
type Play struct {
	card   string
	pay    string
	target interface{}
	score  int
}

func (play *Play) Score() int {
	return play.score
}

func (play *Play) Submit(request RequestFunc) {
	request("play", map[string]any{
		"id":     play.card,
		"pay":    play.pay,
		"target": play.target,
	})
}

func (play *Play) String() string {
	return "Play " + play.card
}

func ParsePlay(view view.T) (play *Play) {
	for cardID := range view.Self.T.Hand {
		card := view.Game.Card(cardID)
		if !view.Self.T.Karma.Active().Test(card.T.Costs) {
			continue
		}

		p := ParsePlayWith(view, card)
		if p == nil {
			continue
		}

		view.Game.Log().With(map[string]any{
			"Card":  p.card,
			"Score": p.score,
		}).Trace("potential")

		if play == nil {
			play = p
		} else if p.score > play.score {
			play = p
		}
	}

	return
}

func ParsePlayWith(view view.T, c *game.Card) *Play {
	switch c.T.Kind {
	case card.Being:
		if !view.IsMyMain {
			return nil
		}
		return &Play{
			card:   c.ID(),
			pay:    PayString(view, c.T.Costs),
			target: nil,
			score:  1,
		}
	case card.Item:
		if !view.IsMyMain {
			return nil
		}
		return &Play{
			card:   c.ID(),
			pay:    PayString(view, c.T.Costs),
			target: nil,
			score:  1,
		}
	case card.Spell:
		if ps := c.T.Powers.GetTrigger("play"); len(ps) < 1 {
		} else if p := ps[0]; p == nil {
		} else if id, score := playTargetScore(view, c, p); score < 1 {
		} else {
			return &Play{
				card:   c.ID(),
				pay:    PayString(view, c.T.Costs),
				target: id,
				score:  score,
			}
		}
	}
	return nil
}

func playTargetScore(view view.T, card *game.Card, p *power.T) (interface{}, int) {
	switch card.T.ID {
	case 9:
		return aim.EnemyBeing(view, "damage")
	case 10:
		return aim.MyPresentBeingItem(view, "wake")
	case 11:
		return aim.MyPresentBeing(view, "wake")
	case 12:
		return aim.EnemyBeing(view, "")
	case 13:
		return aim.MyPresentBeing(view, "health")
	case 14:
		return aim.MyPastBeing(view)
	}
	return nil, 0
}
