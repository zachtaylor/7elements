package games

import (
	"ztaylor.me/js"
)

func Autopass(g *Game) func() {
	return func() {
		autopass := make([]string, 0)
		for _, s := range g.Seats {
			if g.Active.Resp[s.Username] == "pass" {
			} else if s.HasActiveElements() && s.HasCardsInHand() {
				return
			} else {
				autopass = append(autopass, s.Username)
			}
		}

		if len(autopass) > 0 {
			for _, username := range autopass {
				g.Broadcast("pass", js.Object{
					"gameid":   g.Id,
					"target":   g.Active.Target,
					"username": username,
					"auto":     true,
				})
			}
		}
	}
}
