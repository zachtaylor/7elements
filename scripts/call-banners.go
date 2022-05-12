package scripts

import (
	"strconv"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/power"
)

func init() { game.Scripts["call-banners"] = CallTheBanners }

var callthebannersTokenCardProto = card.Prototype{
	ID:     3, // provides default image for clone
	Kind:   card.Being,
	Name:   "Bannerman",
	Text:   "At your call",
	Costs:  element.Count{},
	Body:   &card.Body{Attack: 2, Life: 2},
	Powers: power.NewSet(),
}

func CallTheBanners(g *game.G, ctx game.ScriptContext) (rs []game.Phaser, err error) {
	player := g.Player(ctx.Player)
	card := g.NewCard(ctx.Player, callthebannersTokenCardProto)

	for i := 0; i < 3; i++ {
		tokenCtx := game.NewTokenContext(card)
		tokenCtx.Image = "17." + strconv.FormatInt(int64(i), 10)

		if triggered := trigger.TokenAdd(g, player, tokenCtx); len(triggered) > 1 {
			rs = append(rs, triggered...)
		}
	}
	return rs, nil
}
