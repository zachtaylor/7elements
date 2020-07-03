package out

import "ztaylor.me/cast"

func GameReact(t Target, stateid, username, react string, timer cast.Duration) {
	t.Send("/game/react", NewGameReact(stateid, username, react, timer))
}

func NewGameReact(stateid, username, react string, timer cast.Duration) cast.JSON {
	return cast.JSON{
		"stateid":  stateid,
		"username": username,
		"react":    react,
		"timer":    int(timer.Seconds()),
	}
}
