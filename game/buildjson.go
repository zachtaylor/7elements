package game

import (
	"ztaylor.me/cast"
)

// BuildPushJSON returns a new JSON object with "uri", "data"
func BuildPushJSON(uri string, json cast.JSON) cast.JSON {
	return cast.JSON{
		"uri":  uri,
		"data": json,
	}
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
