package seat

// List is an organization of seats
type List struct {
	data map[string]*T
}

func NewList() *List { return &List{data: make(map[string]*T)} }

func (l *List) Keys() []string {
	keys := make([]string, 0)
	for k := range l.data {
		keys = append(keys, k)
	}
	return keys
}

func (l *List) Count() int { return len(l.data) }

func (l *List) Add(seat *T) { l.data[seat.Username] = seat }

func (l *List) Get(name string) *T {
	for k, seat := range l.data {
		if k == name {
			return seat
		}
	}
	return nil
}

func (l *List) GetOpponent(name string) *T {
	for k, seat := range l.data {
		if k != name {
			return seat
		}
	}
	return nil
}

func (l *List) Write(data []byte) { go l.WriteSync(data) }

func (l *List) WriteSync(data []byte) {
	for _, seat := range l.data {
		seat.Writer.Write(data)
	}
}
