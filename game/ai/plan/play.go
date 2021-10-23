package plan

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai/aim"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/power"
	"taylz.io/http/websocket"
)

// Play is a plan to make a play
type Play struct {
	card   *card.T
	pay    string
	target interface{}
	score  int
}

func (play *Play) Score() int {
	return play.score
}

func (play *Play) Submit(request RequestFunc) {
	request("play", websocket.MsgData{
		"id":     play.card.ID,
		"pay":    play.pay,
		"target": play.target,
	})
}

func (play *Play) String() string {
	return "Play " + play.card.ID
}

func ParsePlay(game *game.T, seat *seat.T) (play *Play) {
	for _, card := range seat.Hand {
		if !seat.Karma.Active().Test(card.Proto.Costs) {
			continue
		}

		p := ParsePlayWith(game, seat, card)
		if p == nil {
			continue
		}

		game.Log().With(websocket.MsgData{
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

func ParsePlayWith(game *game.T, seat *seat.T, c *card.T) *Play {
	switch c.Proto.Type {
	case card.BodyType:
		if game.Phase() != "main" || game.State.Phase.Seat() != seat.Username {
			return nil
		}
		return &Play{
			card:   c,
			pay:    PayString(seat, c.Proto.Costs),
			target: nil,
			score:  1,
		}
	case card.ItemType:
		if game.Phase() != "main" || game.State.Phase.Seat() != seat.Username {
			return nil
		}
		return &Play{
			card:   c,
			pay:    PayString(seat, c.Proto.Costs),
			target: nil,
			score:  1,
		}
	case card.SpellType:
		if ps := c.Proto.Powers.GetTrigger("play"); len(ps) < 1 {
		} else if p := ps[0]; p == nil {
		} else if id, score := playTargetScore(game, seat, c, p); score < 1 {
		} else {
			return &Play{
				card:   c,
				pay:    PayString(seat, c.Proto.Costs),
				target: id,
				score:  score,
			}
		}
	}
	return nil
}

func playTargetScore(game *game.T, seat *seat.T, card *card.T, p *power.T) (interface{}, int) {
	switch card.Proto.ID {
	case 9:
		return aim.EnemyBeing(game, seat, "damage")
	case 10:
		return aim.MyPresentBeingItem(game, seat, "wake")
	case 11:
		return aim.MyPresentBeing(game, seat, "wake")
	case 12:
		return aim.EnemyBeing(game, seat, "")
	case 13:
		return aim.MyPresentBeing(game, seat, "health")
	case 14:
		return aim.MyPastBeing(game, seat)
	}
	return nil, 0
}
