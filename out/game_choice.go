package out

import "ztaylor.me/cast"

// Choice returns JSON command to ask player to choose something
func Choice(t Target, prompt string, choices []cast.JSON, data cast.JSON) {
	t.Send("/game/choice", NewChoice(prompt, choices, data))
}

// NewChoice returns a JSON spec to game choice data through "/game/choice"
func NewChoice(prompt string, choices []cast.JSON, data cast.JSON) cast.JSON {
	return cast.JSON{
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	}
}
