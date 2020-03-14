package power

import "ztaylor.me/cast"

// Set contains Powers, mapped by their ID for easy lookup
type Set map[int]*T

// NewSet returns a new empty Set of Powers
func NewSet() Set {
	return Set{}
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

// JSON returns a representation of this Set of Powers as type []cast.JSON
func (set Set) JSON() []cast.JSON {
	if set == nil {
		return nil
	}
	json := make([]cast.JSON, 0)
	for _, power := range set {
		json = append(json, power.JSON())
	}
	return json
}
