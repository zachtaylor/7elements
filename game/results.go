package game

// type Resulter interface {
// 	String() string
// 	IsDraw() bool
// 	Winner() string
// 	Loser() string
// }

// type WinLossResult struct {
// 	winner string
// 	loser  string
// }

// func NewWinLoss(winner, loser string) Resulter {
// 	return &WinLossResult{
// 		winner: winner,
// 		loser:  loser,
// 	}
// }
// func (*T) NewWinLoss(winner, loser string) Resulter { return NewWinLoss(winner, loser) }

// func (r *WinLossResult) String() string { return "W:" + r.winner + " L:" + r.loser }
// func (*WinLossResult) IsDraw() bool     { return false }
// func (r *WinLossResult) Winner() string { return r.winner }
// func (r *WinLossResult) Loser() string  { return r.loser }

// type DrawResult struct{}

// func NewDraw() Resulter            { return &DrawResult{} }
// func (*T) NewDraw() Resulter       { return NewDraw() }
// func (*DrawResult) String() string { return "draw" }
// func (*DrawResult) IsDraw() bool   { return true }
// func (*DrawResult) Winner() string { return "" }
// func (*DrawResult) Loser() string  { return "" }
