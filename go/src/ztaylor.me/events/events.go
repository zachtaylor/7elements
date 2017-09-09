package events

type Bus map[string]map[int]func(...interface{})

var globals = Bus{}

func (bus Bus) On(event string, f func(...interface{})) int {
	if bus[event] == nil {
		bus[event] = make(map[int]func(...interface{}))
	}

	id := len(bus[event])

	bus[event][id] = f

	return id
}

func (bus Bus) Off(event string, id int) {
	if chain := bus[event]; chain != nil {
		chain[id] = nil
	}
}

func On(event string, f func(...interface{})) int {
	return globals.On(event, f)
}

func Off(event string, id int) {
	globals.Off(event, id)
}

func (bus Bus) Fire(event string, args ...interface{}) {
	if chain := bus[event]; chain != nil {
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
	globals.Fire(event, args...)
}
