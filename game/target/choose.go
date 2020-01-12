package target

// import (
// 	vii "github.com/zachtaylor/7elements"
// 	"github.com/zachtaylor/7elements/game"
// 	"ztaylor.me/cast"
// )

// type Choice struct {
// 	Choice  string
// 	Display string
// }

// func (c *Choice) JSON() cast.JSON {
// 	return cast.JSON{}
// }

// var HelpChoose = map[string]ChooseFunc{
// 	"self": ChooseSelf,
// }

// type ChooseFunc func(*game.T, *game.Seat, *game.Card, *vii.Power, interface{}) ([]*Choice, error)

// func ChooseSelf(g *game.T, seat *game.Seat, card *game.Card, p *vii.Power, arg interface{}) ([]*Choice, error) {
// 	return []*Choice{
// 		&Choice{
// 			Choice: card.Id,
// 		},
// 	}, nil
// }
