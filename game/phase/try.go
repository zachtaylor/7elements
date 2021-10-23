package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

func TryOnRequest(g *game.T, seat *seat.T, json map[string]interface{}) {
	if requester, ok := g.State.Phase.(game.OnRequestPhaser); ok {
		requester.OnRequest(g, seat, json)
	}
}

func TryOnFinish(g *game.T) (rs []game.Phaser) {
	if finisher, _ := g.State.Phase.(game.OnFinishPhaser); finisher != nil {
		rs = finisher.OnFinish(g)
	}
	if rs == nil {
		rs = []game.Phaser{}
	}
	return
}

func TryOnConnect(g *game.T, seat *seat.T) {
	if seat == nil {
		g.Log().Add("State", g.State).Trace("reconnect broadcast state")
		g.Seats.Write(wsout.GameState(g.State.Data()).EncodeToJSON())
	}
	if connector, ok := g.State.Phase.(game.OnConnectPhaser); ok {
		connector.OnConnect(g, seat)
	}
}

func TryOnActivate(g *game.T) (rs []game.Phaser) {
	if activator, _ := g.State.Phase.(game.OnActivatePhaser); activator != nil {
		rs = activator.OnActivate(g)

		if len(g.State.Phase.Seat()) < 1 {
			return
		}
	}
	if events := tryName(g); len(events) > 0 {
		rs = append(rs, tryName(g)...)
	}
	TryOnConnect(g, nil)
	if rs == nil {
		rs = []game.Phaser{}
	}
	return
}

func tryName(game *game.T) (rs []game.Phaser) {
	rs = append(rs, trySeatName(game.Seats.Get(game.State.Phase.Seat()), "my-"+game.Phase())...)
	for _, seatName := range game.Seats.Keys() {
		seat := game.Seats.Get(seatName)
		rs = append(rs, trySeatName(seat, game.Phase())...)
		if seatName != game.State.Phase.Seat() {
			rs = append(rs, trySeatName(seat, "op-"+game.Phase())...)
		}
	}
	return
}

func trySeatName(seat *seat.T, name string) (rs []game.Phaser) {
	for _, token := range seat.Present {
		if e := tryTokenName(token, name); len(e) > 0 {
			rs = append(rs, e...)
		}
	}
	return
}

func tryTokenName(token *token.T, name string) (rs []game.Phaser) {
	powers := token.Powers.GetTrigger(name)
	if len(powers) < 1 {
		return nil
	}
	for _, p := range powers {
		if p.Target != "self" {
			rs = append(rs, NewTarget(token.User, p, token))
		} else {
			rs = append(rs, NewTrigger(token.User, token, p, token.ID))
		}
		// }
	}
	return
}
