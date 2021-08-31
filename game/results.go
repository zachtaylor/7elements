package game

type Resulter interface {
	String() string
	IsDraw() bool
	Winner() string
	Loser() string
}

type WinLossResult struct {
	winner string
	loser  string
}

func (r *WinLossResult) String() string { return "W:" + r.winner + " L:" + r.loser }
func (*WinLossResult) IsDraw() bool     { return false }
func (r *WinLossResult) Winner() string { return r.winner }
func (r *WinLossResult) Loser() string  { return r.loser }

type DrawResult struct{}

func (*DrawResult) String() string { return "draw" }
func (*DrawResult) IsDraw() bool   { return true }
func (*DrawResult) Winner() string { return "" }
func (*DrawResult) Loser() string  { return "" }
