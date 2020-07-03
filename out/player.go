package out

import "ztaylor.me/cast"

// Target is an endpoint representing a player
//
// Known to be filled by `/server/runtime/*Player` and `/game/ai/*Input`
type Target interface {
	// Send is the way the game engine sends data to players
	Send(uri string, data cast.JSON)
}
