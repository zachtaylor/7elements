package event

type Eventer map[string]map[int]func(...interface{})

var events = Eventer{}

func (events Eventer) On(event string, f func(...interface{})) int {
	if events[event] == nil {
		events[event] = make(map[int]func(...interface{}))
	}

	id := len(events[event])

	events[event][id] = f

	return id
}

func On(event string, f func(...interface{})) int {
	return events.On(event, f)
}

func (events Eventer) Off(event string, id int) {
	if chain := events[event]; chain != nil {
		chain[id] = nil
	}
}

func Off(event string, id int) {
	events.Off(event, id)
}

func (events Eventer) Fire(event string, args ...interface{}) {
	if chain := events[event]; chain != nil {
		if event != "Fire" {
			Fire("Fire", event)
		}

		for _, f := range chain {
			if f == nil {
				continue
			}

			go f(args...)
		}
	}
}

func Fire(event string, args ...interface{}) {
	events.Fire(event, args...)
}
