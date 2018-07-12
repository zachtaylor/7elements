package engine

// import (
// 	"ztaylor.me/js"
// )

// type ChoiceMode func(js.Object)

// func ChoiceModeFunc(f func(js.Object)) EMode {
// 	return ChoiceMode(f)
// }

// func (mode ChoiceMode) Name() string {
// 	return "choice"
// }

// func (mode ChoiceMode) Json(e *Event, g *vii.Game, s *vii.GameSeat) js.Object {
// 	return js.Object{
// 		"gameid":   g.Id,
// 		"username": s.Username,
// 		"timer":    int(e.Duration.Seconds()),
// 	}
// }

// func (mode ChoiceMode) OnActivate(*Event, *vii.Game) {
// }

// func (mode ChoiceMode) OnSendCatchup(*Event, *vii.Game, *vii.GameSeat) {
// }

// func (mode ChoiceMode) OnResolve(*Event, *vii.Game) {
// }

// func (mode ChoiceMode) OnReceive(e *Event, g *vii.Game, s *vii.GameSeat, json js.Object) {
// 	if s.Username != e.Username {
// 		g.Log().Add("Username", s.Username).Add("HotSeat", e.Username).Warn("games.ChoiceMode: not your choice")
// 		return
// 	}

// 	mode(json)
// }
