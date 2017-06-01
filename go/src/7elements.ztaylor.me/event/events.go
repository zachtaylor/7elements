package event

var events = make(map[string]map[int]func(...interface{}))

func On(event string, f func(...interface{})) int {
	if events[event] == nil {
		events[event] = make(map[int]func(...interface{}))
	}

	id := len(events[event])

	events[event][id] = f

	return id
}

func Off(event string, id int) {
	if chain := events[event]; chain != nil {
		chain[id] = nil
	}
}

func Fire(event string, args ...interface{}) {
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
