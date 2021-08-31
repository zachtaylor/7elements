package phase

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func NewPlay(seat string, card *card.T, target string) game.Phaser {
	return &Play{
		R:      R(seat),
		Card:   card,
		Target: target,
	}
}

type Play struct {
	R
	Card        *card.T
	Target      string
	IsCancelled bool
}

func (r *Play) Name() string { return "play" }

func (r *Play) String() string {
	return "play (" + r.Seat() + ":" + r.Card.String() + ")"
}

// OnActivate implements game.OnActivatePhaser
func (r *Play) OnActivate(game *game.T) []game.Phaser {
	msg := r.Card.Proto.Name
	if r.Card.Proto.Text != "" {
		msg = r.Card.Proto.Text
	}
	go game.Chat(r.Seat(), msg)
	return nil
}
func (r *Play) onActivatePhaser() game.OnActivatePhaser { return r }

// Finish implements game.OnFinishPhaser
func (r *Play) OnFinish(g *game.T) (rs []game.Phaser) {
	seat := g.Seats.Get(r.Seat())
	g.Log().With(websocket.MsgData{
		"Seat":        seat.String(),
		"Card":        r.Card.String(),
		"IsCancelled": r.IsCancelled,
	}).Debug("engine/play: finish")
	seat.Past[r.Card.ID] = r.Card
	g.Seats.Write(wsout.GameSeatJSON(seat.Data()))

	if r.IsCancelled {
		return nil
	}

	if r.Card.Proto.Type == card.BodyType || r.Card.Proto.Type == card.ItemType {
		if events := g.Engine().NewToken(g, seat, token.New(r.Card, seat.Username)); events != nil {
			rs = append(rs, events...)
		}
	} else {
		g.Seats.Write(wsout.GameSeatJSON(seat.Data()))
	}

	powers := r.Card.Proto.Powers.GetTrigger("play")
	for _, power := range powers {
		if power.Target == "self" {
			if events := script.Run(g, seat, power, r.Card, []string{r.Card.ID}); events != nil {
				rs = append(rs, events...)
			}
		} else if events := script.Run(g, seat, power, r.Card, []string{r.Target}); len(events) > 0 {
			rs = append(rs, events...)
		}
		// } else if r.Target != nil {
		// 	if events := script.Run(g, seat, power, r.Card, []string{r.Target}); events != nil {
		// 		rs = append(rs, events...)
		// 	}
		// } else {
		// 	rs = append(rs, NewTarget(
		// 		seat.Username,
		// 		power.Target,
		// 		power.Text,
		// 		func(val string) []game.Phaser {
		// 			return script.Run(g, seat, power, r.Card, []string{val})
		// 		},
		// 	))
		// }
	}

	return
}
func (r *Play) onFinishPhaser() game.OnFinishPhaser { return r }

func (r *Play) GetNext(game *game.T) game.Phaser {
	return nil
}

func (r *Play) Data() websocket.MsgData {
	json := websocket.MsgData{
		"card":   r.Card.Data(),
		"target": r.Target,
	}
	return json
}
