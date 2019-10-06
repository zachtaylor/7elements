package event

type Event string

func (e Event) Seat() string {
	return string(e)
}
