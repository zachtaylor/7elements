package power

// Set contains Powers, mapped by their ID for easy lookup
type Set map[int]*T

// NewSet returns a new empty Set of Powers
func NewSet() Set {
	return Set{}
}

func (set Set) Keys() []int {
	i, keys := 0, make([]int, len(set))
	for k := range set {
		keys[i] = k
		i++
	}
	return keys
}

// Copy returns a shallow copy of this Set of Powers
func (set Set) Copy() Set {
	cp := NewSet()
	for k, v := range set {
		cp[k] = v.Copy()
	}
	return cp
}

// GetTrigger returns a slice of Powers that match the given trigger name within this Set
func (set Set) GetTrigger(name string) []*T {
	ps := make([]*T, 0)
	for _, p := range set {
		if p.Trigger == name {
			ps = append(ps, p)
		}
	}
	return ps
}

// JSON returns a representation of this Set of Powers as type []map[string]any
func (set Set) JSON() []map[string]any {
	if set == nil {
		return nil
	}
	json := make([]map[string]any, 0)
	for _, power := range set {
		json = append(json, power.JSON())
	}
	return json
}
