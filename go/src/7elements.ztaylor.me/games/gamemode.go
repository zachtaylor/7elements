package games

// type GameMode interface {
// 	Accessor() string
// 	BuildNext() GameMode
// }

// type GameMode_Nil bool

// func (mode *GameMode_Nil) Accessor() string {
// 	return ""
// }

// type GameMode_Begin bool

// func (mode *GameMode_Begin) Accessor() string {
// 	return "begin"
// }

// func (mode *GameMode_Begin) BuildNext() GameMode {
// 	return &GameMode_Play{}
// }

// type GameMode_Play bool

// func (mode *GameMode_Play) Accessor() string {
// 	return "begin"
// }

// func (mode *GameMode_Play) BuildNext() GameMode {
// 	return &GameMode_Play{}
// }

// type GameMode_Respond bool

// func (mode *GameMode_Respond) Accessor() string {
// 	return "respond"
// }

// func (mode *GameMode_Respond) BuildNext() GameMode {
// 	return &GameMode_Play{}
// }

// type GameMode_Done bool

// func (mode *GameMode_Respond) Accessor() string {
// 	return ""
// }

// func (mode *GameMode_Respond) BuildNext() GameMode {
// 	return &GameMode_Play{}
// }
