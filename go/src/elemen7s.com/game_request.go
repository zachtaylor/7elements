package vii

type GameRequest struct {
	Username string
	Data     Json
}

func (r GameRequest) String() string {
	return r.Username + ":" + r.Data.String()
}
