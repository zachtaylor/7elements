package games

type Tally int

func NewTally() *Tally {
	t := Tally(0)
	return &t
}

func (t *Tally) Current() int {
	return int(*t)
}

func (t *Tally) Next() int {
	(*t)++
	return t.Current()
}
