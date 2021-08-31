package token

// Map is a map of tid->T
type Map map[string]*T

func (tokens Map) Keys() []string {
	i, keys := 0, make([]string, len(tokens))
	for k := range tokens {
		keys[i] = k
		i++
	}
	return keys
}

func (tokens Map) Has(id string) (ok bool) {
	_, ok = tokens[id]
	return
}

// Data returns a representation of a set of game tokens
func (tokens Map) Data() []map[string]interface{} {
	data := make([]map[string]interface{}, 0)
	for _, c := range tokens {
		data = append(data, c.Data())
	}
	return data
}
