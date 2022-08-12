package out

// import (
// 	"time"

// 	"taylz.io/http/websocket"
// )

// func Game(data websocket.JSON) []byte { return encode("/game", data) }

// func GameCard(data websocket.JSON) []byte { return encode("/game/card", data) }

// func GameHand(ids []string) []byte {
// 	return encode("/game/hand", websocket.JSON{
// 		"cards": ids,
// 	})
// }

// func GamePresent(seat string, ids []string) []byte {
// 	return encode("/game/present", websocket.JSON{
// 		"username": seat,
// 		"present":  ids,
// 	})
// }

// func GameReact(stateid, username, react string, time time.Duration) []byte {
// 	return encode("/game/react", websocket.JSON{
// 		"stateid":  stateid,
// 		"username": username,
// 		"react":    react,
// 		"timer":    int(time.Seconds()),
// 	})
// }

// func GameSeat(data websocket.JSON) []byte { return encode("/game/seat", data) }

// func GameState(data websocket.JSON) []byte { return encode("/game/state", data) }

// func GameToken(data websocket.JSON) []byte { return encode("/game/token", data) }

// func GameChoice(prompt string, choices []websocket.JSON, data websocket.JSON) []byte {
// 	return encode("/game/choice", websocket.JSON{
// 		"prompt":  prompt,
// 		"choices": choices,
// 		"data":    data,
// 	})
// }

// func GameChoiceElements(prompt string, data websocket.JSON) []byte {
// 	return GameChoice(prompt, GameChoiceElementsData, data)
// }

// var GameChoiceElementsData = []websocket.JSON{
// 	{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
// 	{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
// 	{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
// 	{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
// 	{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
// 	{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
// 	{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
// }
