package game

import "taylz.io/yas"

// Priority is a []string containing seatid priority
type Priority []string

func (p Priority) Unique() []string {
	set := yas.Set[string]{}
	for _, id := range p {
		set.Add(id)
	}
	return set.Slice()
}

func (p Priority) Reverse() Priority {
	len := len(p)
	r := make(Priority, len)
	for i, id := range p {
		r[len-1-i] = id
	}
	return r
}

// func (p Priority) WithLeader(id string) Priority {
// 	if len(p) < 1 {
// 		return Priority{id}
// 	} else if id == p[0] {
// 		return p
// 	} else if i := slices.Index(p, id); i < 1 {
// 		return append([]string{id}, p...)
// 	} else {
// 		return append([]string{id}, append(p[:i], p[i+i:]...)...)
// 	}
// }
