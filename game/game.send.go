package game

import "ztaylor.me/cast"

// SendAll sends data to all seats
func (game *T) SendAll(json cast.JSON) {
	for _, seat := range game.Seats {
		seat.Send(json)
	}
}

// SendUpdate creates a data push object, and uses SendAll
func (game *T) SendUpdate(uri string, data cast.JSON) {
	game.SendAll(BuildPushJSON(uri, data))
}

// SendStateUpdate sends game state data using SendUpdate
func (game *T) SendStateUpdate() {
	game.SendUpdate("/game/state", game.State.JSON())
}

// SendSeatUpdate sends game seat data using SendUpdate
func (game *T) SendSeatUpdate(seat *Seat) {
	game.SendUpdate("/game/seat", seat.JSON())
}

// SendPrivateUpdate sends game private data only to a seat
func (game *T) SendPrivateUpdate(seat *Seat) {
	seat.Send(BuildPushJSON("/game", game.PerspectiveJSON(seat)))
}

// SendCardUpdate sends card data using SendDataUpdate
func (game *T) SendCardUpdate(c *Card) {
	game.SendUpdate("/game/card", c.JSON())
}

// SendReactUpdate sends JSON command to update state reaction
func (game *T) SendReactUpdate(username string) {
	game.SendUpdate("/game/react", cast.JSON{
		"stateid":  game.State.ID(),
		"username": username,
		"react":    game.State.Reacts[username],
		"timer":    int(game.State.Timer.Seconds()),
	})
}
