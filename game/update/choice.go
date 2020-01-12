package update

import "ztaylor.me/cast"

// Choice returns JSON command to ask player to choose something
func Choice(writer Writer, prompt string, choices []cast.JSON, data cast.JSON) {
	writer.WriteJSON(Build("/game/choice", cast.JSON{
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	}))
}

// ChoicesNewElement is sent to display a choice of new element
var ChoicesNewElement = []cast.JSON{
	cast.JSON{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
	cast.JSON{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
	cast.JSON{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
	cast.JSON{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
	cast.JSON{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
	cast.JSON{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
	cast.JSON{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
}
