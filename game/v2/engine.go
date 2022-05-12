package game

func Activate(g *G, phase Phaser) (rs []Phaser) {
	rs = TryOnActivate(g, phase)
	if events := OnActivatePhase(g, phase); len(events) > 0 {
		rs = append(rs, events...)
	}
	TryOnConnect(g, phase, nil)
	return
}

func OnActivatePhase(g *G, phase Phaser) (rs []Phaser) {
	priority := phase.Priority().Unique()
	if len(priority) < 1 {
		g.Log().Out("activate no priority")
		return
	}

	for i, playerID := range priority {
		player := g.Player(playerID)

		if i == 0 {
			if events := OnActivatePhaseSeat(g, player, "my-"+phase.Type()); len(events) > 0 {
				rs = append(rs, events...)
			}
		} else {
			if events := OnActivatePhaseSeat(g, player, "op-"+phase.Type()); len(events) > 0 {
				rs = append(rs, events...)
			}
		}

		if events := OnActivatePhaseSeat(g, player, phase.Type()); len(events) > 0 {
			rs = append(rs, events...)
		}
	}
	return
}
func OnActivatePhaseSeat(g *G, player *Player, phase string) (rs []Phaser) {
	for tokenID := range player.T.Present {
		if token := Get[TokenContext](g, tokenID); token == nil {
			g.Log().Error("weird tokenid in present", tokenID)
		} else if events := OnActivatePhaseSeatToken(g, player, token, phase); len(events) > 0 {
			rs = append(rs, events...)
		}
	}
	return
}
func OnActivatePhaseSeatToken(g *G, seat *Player, token *Token, phase string) (rs []Phaser) {
	if triggered := token.T.Powers.GetTrigger(phase); len(triggered) > 0 {
		for _, power := range triggered {
			rs = append(rs, NewTriggerPhase(g, token, power))
		}
	}
	return
}
