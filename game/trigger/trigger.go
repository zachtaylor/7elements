package trigger

import "github.com/zachtaylor/7elements/game"

func PlayerDamage(g *game.G, player *game.Player, n int) (rs []game.Phaser) {
	g.Log().Trace("damage-player", player.T.Writer.Name(), player.T.Life, n)
	player.T.Life -= n
	g.MarkUpdate(player.ID())
	if player.T.Life < 1 {
		return []game.Phaser{
			game.NewLoserPhase(g.NewPriority(player.ID()), player.ID()),
		}
	}
	// game.Chat(card.Proto.Name, strconv.FormatInt(int64(n), 10)+" damage to "+seat.Username)
	return nil
}

func PlayerHeal(g *game.G, player *game.Player, n int) []game.Phaser {
	g.Log().Trace("player-heal", player.T.Writer.Name(), player.T.Life, n)
	player.T.Life += n
	// game.Chat(seat.Username, "gain "+strconv.FormatInt(int64(n), 10)+" Life")
	g.MarkUpdate(player.ID())
	return nil // todo
}

func TokenAwake(g *game.G, token *game.Token) (rs []game.Phaser) {
	wasAwake := token.T.Awake
	token.T.Awake = true
	if !wasAwake {
		if triggered := token.T.Powers.GetTrigger("become-awake"); len(triggered) > 0 {
			for _, power := range triggered {
				rs = append(rs, game.NewTriggerPhase(g, token, power))
			}
		}
		g.MarkUpdate(token.ID())
	}
	return
}

func TokenSleep(g *game.G, token *game.Token) (rs []game.Phaser) {
	wasAwake := token.T.Awake
	token.T.Awake = false
	if wasAwake {
		if triggered := token.T.Powers.GetTrigger("become-asleep"); len(triggered) > 0 {
			for _, power := range triggered {
				rs = append(rs, game.NewTriggerPhase(g, token, power))
			}
		}
		g.MarkUpdate(token.ID())
	}
	return
}

func TokenDamage(g *game.G, token *game.Token, n int) (rs []game.Phaser) {
	g.Log().Trace("damage-token", token.ID(), token.Player(), token.T.Name, n)

	token.T.Body.Life -= n

	// game.Chat("vii", strconv.FormatInt(int64(n), 10)+" damage to "+token.Card.Proto.Name)

	g.MarkUpdate(token.ID())

	if powers := token.T.Powers.GetTrigger("take-damage"); len(powers) > 0 {
		for _, power := range powers {
			rs = append(rs, game.NewTriggerPhase(g, token, power))
		}
	}
	if token.T.Body.Life <= 0 {
		if triggered := TokenRemove(g, token); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}
	return
}

func TokenHeal(g *game.G, token *game.Token, n int) []game.Phaser {
	token.T.Body.Life += n
	g.MarkUpdate(token.ID())
	return nil // todo
}

func TokenAdd(g *game.G, player *game.Player, ctx game.TokenContext) (rs []game.Phaser) {
	token := g.NewToken(player.ID(), ctx)
	player.T.Present.Add(token.ID())
	g.MarkUpdate(player.ID())
	g.MarkUpdate(token.ID())
	return nil // todo
}

func TokenRemove(g *game.G, token *game.Token) (rs []game.Phaser) {
	if token.T.Body != nil {
		token.T.Body.Life = 0
	}
	player := g.Player(token.Player())
	// game.Chat("vii", token.Card.Proto.Name+" died")
	player.T.Present.Remove(token.ID())
	g.MarkUpdate(token.ID())
	g.MarkUpdate(player.ID())

	if triggered := token.T.Powers.GetTrigger("death"); len(triggered) > 0 {
		for _, power := range triggered {
			rs = append(rs, game.NewTriggerPhase(g, token, power))
		}
	}

	return
}
