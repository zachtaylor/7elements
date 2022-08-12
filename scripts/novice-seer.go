package scripts

import (
	"reflect"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"taylz.io/yas"
)

const noviceseerId = "novice-seer"

func init() { game.Scripts[noviceseerId] = NoviceSeer }

func NoviceSeer(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	player := g.Player(ctx.Player)
	if player == nil {
		return nil, ErrPlayerID
	} else if len(player.T.Future) < 1 {
		return nil, ErrFutureEmpty
	}
	cardID, future := yas.Shift(player.T.Future)
	player.T.Future = future
	card := g.Card(cardID)
	if card == nil {
		return nil, ErrCardID
	}

	return []game.Phaser{
		phase.NewChoice(
			player.ID(),
			"Novice Seer",
			map[string]any{
				"card": card.T.Data(),
			},
			[]map[string]any{
				map[string]any{
					"choice":  "false",
					"display": "Put on top of your Future",
				},
				map[string]any{
					"choice":  "true",
					"display": "Put into your Past",
				},
			},
			func(val interface{}) {
				g.Log().Add("val", val).Add("type", reflect.TypeOf(val)).Info()
				if destroy, _ := val.(bool); destroy {
					player.T.Past.Add(card.ID())
				} else {
					player.T.Future = yas.Unshift(player.T.Future, card.ID())
				}
				g.MarkUpdate(player.ID())
			},
		),
	}, nil
}
