package state

type R string

func (r R) Seat() string {
	return string(r)
}
