package ctxpert

type Timeout func(*Context)

type timeouts []Timeout

func (ts timeouts) copy() timeouts {
	if ts == nil {
		return make(timeouts, 0)
	}
	cp := make(timeouts, len(ts))
	copy(cp, ts)
	return cp
}
