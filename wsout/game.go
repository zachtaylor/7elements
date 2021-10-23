package wsout

import (
	"time"

	"taylz.io/http/websocket"
)

func Game(data websocket.MsgData) []byte { return websocket.NewMessage("/game", data).EncodeToJSON() }

func GameCardJSON(data websocket.MsgData) []byte { return GameCard(data).EncodeToJSON() }

func GameCard(data websocket.MsgData) *websocket.Message {
	return websocket.NewMessage("/game/card", data)
}

func GameHand(ids []string) *websocket.Message {
	return websocket.NewMessage("/game/hand", websocket.MsgData{
		"cards": ids,
	})
}

func GamePresentJSON(seat string, ids []string) []byte { return GamePresent(seat, ids).EncodeToJSON() }

func GamePresent(seat string, ids []string) *websocket.Message {
	return websocket.NewMessage("/game/present", websocket.MsgData{
		"username": seat,
		"present":  ids,
	})
}

func GameReact(stateid, username, react string, time time.Duration) *websocket.Message {
	return websocket.NewMessage("/game/react", websocket.MsgData{
		"stateid":  stateid,
		"username": username,
		"react":    react,
		"timer":    int(time.Seconds()),
	})
}

func GameSeat(data websocket.MsgData) *websocket.Message {
	return websocket.NewMessage("/game/seat", data)
}

func GameState(data websocket.MsgData) *websocket.Message {
	return websocket.NewMessage("/game/state", data)
}

func GameToken(data websocket.MsgData) *websocket.Message {
	return websocket.NewMessage("/game/token", data)
}

func GameTokenJSON(data websocket.MsgData) []byte { return GameToken(data).EncodeToJSON() }

func GameSeatJSON(data websocket.MsgData) []byte { return GameSeat(data).EncodeToJSON() }

func GameChoice(prompt string, choices []websocket.MsgData, data websocket.MsgData) *websocket.Message {
	return websocket.NewMessage("/game/choice", websocket.MsgData{
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	})
}

func GameChoiceElement(prompt string, data websocket.MsgData) *websocket.Message {
	return GameChoice(prompt, GameChoiceElementsData, data)
}

var GameChoiceElementsData = []websocket.MsgData{
	{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
	{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
	{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
	{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
	{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
	{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
	{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
}

func GameChoiceJSON(prompt string, choices []websocket.MsgData, data websocket.MsgData) []byte {
	return GameChoice(prompt, choices, data).EncodeToJSON()
}
