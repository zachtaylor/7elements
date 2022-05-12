package phase

import (
	"reflect"
	"strconv"

	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
)

func NewSunrise(g *game.G, id string) game.Phaser {
	return &Sunrise{
		PriorityContext: game.PriorityContext(g.NewPriority(id)),
		Ans:             make(map[string]element.T),
	}
}

type Sunrise struct {
	game.PriorityContext
	Ans map[string]element.T
}

func (r *Sunrise) Type() string { return "sunrise" }

// OnActivate implements game.OnActivatePhaser
func (r *Sunrise) OnActivate(g *game.G) (rs []game.Phaser) {
	priority := r.Priority()
	g.Log().Trace("activate", priority[0])
	player := g.Player(priority[0])
	player.T.Karma.Reactivate()
	g.MarkUpdate(priority[0])
	for _, playerID := range priority {
		p := g.Player(playerID)
		for tokenID := range p.T.Present {
			if triggered := trigger.TokenAwake(g, g.Token(tokenID)); len(triggered) > 0 {
				rs = append(rs, triggered...)
			}
		}
	}
	return
}

// OnConnect implements game.OnConnectPhaser
func (r *Sunrise) OnConnect(g *game.G, player *game.Player) {
	g.Log().Add("Player", player).Trace("connect")
	if player == nil {
		// go game.Chat("sunrise", r.Seat())
	}
}

// Finish implements game.OnFinishPhaser
func (r *Sunrise) OnFinish(g *game.G, _ *game.State) (rs []game.Phaser) {
	priority := r.Priority()

	if player := g.Player(priority[0]); len(player.T.Future) > 0 {
		if triggered := trigger.DrawCard(g, player, 1); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}

	for _, playerID := range priority {
		player := g.Player(playerID)
		log := g.Log().Add("Username", player.T.Username)

		if el := r.Ans[playerID]; el < element.White || el > element.Black {
			log.Error("el is out of bounds", el)
			continue
		} else {
			player.T.Karma.Append(el, false)
		}
		log.Add("Karma", player.T.Karma).Trace("finish")

		g.MarkUpdate(playerID)
	}

	return
}

func (r *Sunrise) JSON() map[string]any { return nil }

// OnRequest implements game.OnRequestPhaser
func (r *Sunrise) OnRequest(g *game.G, state *game.State, player *game.Player, json map[string]any) {
	if choice, _ := json["choice"].(string); len(choice) < 1 {
		g.Log().Add("choice", json["choice"]).Add("type", reflect.TypeOf(json["choice"])).Warn("choice missing")
	} else if i, err := strconv.ParseInt(choice, 10, 0); err != nil {
		g.Log().Add("choice", i).Add("error", err.Error()).Error("choice parse")
	} else if el := element.T(i); el < element.White || el > element.Black {
		g.Log().Add("choice", i).Warn("invalid element")
	} else {
		g.Log().Add("Player", player).Add("Choice", i).Info("confirm")
		r.Ans[player.ID()] = el
		state.T.React.Set(player.ID())
		g.MarkUpdate(state.ID())
	}
}
