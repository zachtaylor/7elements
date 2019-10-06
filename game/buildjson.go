package game

import (
	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/cast"
)

// BuildPushJSON builds a push object
func BuildPushJSON(uri string, data cast.JSON) cast.JSON {
	return cast.JSON{
		"uri":  uri,
		"data": data,
	}
}

// BuildStateUpdate returns JSON command to update game state
func BuildStateUpdate(s *State) cast.JSON {
	return BuildPushJSON("/game/state", s.JSON())
}

// BuildSpawn returns JSON command to spawn a Card in the Present
func BuildSpawn(g *T, c *Card) cast.JSON {
	return BuildPushJSON("/game/spawn", cast.JSON{
		"username": c.Username,
		"card":     c.JSON(),
	})
}

// BuildUpdate returns JSON command to update game
func BuildGameUpdate(g *T, username string) cast.JSON {
	return BuildPushJSON("/game", g.PerspectiveJSON(username))
}

// BuildCardUpdate returns JSON command to update game card
func BuildCardUpdate(c *Card) cast.JSON {
	return BuildPushJSON("/game/card", c.JSON())
}

// BuildChoiceUpdate returns JSON command to ask player to choose something
func BuildChoiceUpdate(prompt string, choices []cast.JSON, data cast.JSON) cast.JSON {
	return BuildPushJSON("/game/choice", cast.JSON{
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	})
}

// GameChoiceNewElementChoices is sent to display a choice of new element
var GameChoiceNewElementChoices = []cast.JSON{
	cast.JSON{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
	cast.JSON{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
	cast.JSON{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
	cast.JSON{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
	cast.JSON{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
	cast.JSON{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
	cast.JSON{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
}

// BuildElementUpdate returns JSON command to add an element
func BuildElementUpdate(username string, e vii.Element) cast.JSON {
	return BuildPushJSON("/game/element", cast.JSON{
		"username": username,
		"element":  e,
	})
}

// BuildErrorUpdate returns JSON command display an error message
func BuildErrorUpdate(source, message string) cast.JSON {
	return BuildPushJSON("/game/error", cast.JSON{
		"source":  source,
		"message": message,
	})
}

// BuildHandUpdate returns JSON command to update Hand
func BuildHandUpdate(seat *Seat) cast.JSON {
	return BuildPushJSON("/game/hand", cast.JSON{
		"cards": seat.Hand.JSON(),
	})
}

// BuildReactUpdate returns JSON command to update state reaction
func BuildReactUpdate(g *T, username string) cast.JSON {
	return BuildPushJSON("/game/react", cast.JSON{
		"stateid":  g.State.ID(),
		"username": username,
		"react":    g.State.Reacts[username],
		"timer":    int(g.State.Timer.Seconds()),
	})
}

// BuildSeatUpdate returns JSON command to update seat
func BuildSeatUpdate(s *Seat) cast.JSON {
	return BuildPushJSON("/game/seat", s.JSON())
}

// BuildSpawnUpdate returns JSON command to spawn a Card in the Present
func BuildSpawnUpdate(g *T, c *Card) cast.JSON {
	return BuildPushJSON("/game/spawn", cast.JSON{
		"username": c.Username,
		"card":     c.JSON(),
	})
}
