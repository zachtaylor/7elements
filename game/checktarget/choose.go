package checktarget

// import (
// 	vii "github.com/zachtaylor/7elements"
// 	"github.com/zachtaylor/7elements/game"
// 	"ztaylor.me/cast"
// )

// type Choice struct {
// 	Choice  string
// 	Display string
// }

// func (c *Choice) JSON() websocket.MsgData {
// 	return websocket.MsgData{}
// }

// var HelpChoose = map[string]ChooseFunc{
// 	"self": ChooseSelf,
// }

// type ChooseFunc func(*game.T, *seat.T, *game.Card, *vii.Power, interface{}) ([]*Choice, error)

// func ChooseSelf(game *game.T, seat *seat.T, card *game.Card, p *vii.Power, arg interface{}) ([]*Choice, error) {
// 	return []*Choice{
// 		&Choice{
// 			Choice: card.Id,
// 		},
// 	}, nil
// }
