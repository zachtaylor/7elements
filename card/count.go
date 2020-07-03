package card

import "ztaylor.me/cast"

// Count is a map(cardid->quantity)
type Count map[int]int

// Count returns the total number of cards listed
func (c Count) Count() int {
	total := 0
	for _, count := range c {
		total += count
	}
	return total
}

// JSON returns a representation of this count of Cards as type cast.JSON
func (c Count) JSON() cast.JSON {
	json := cast.JSON{}
	for id, count := range c {
		json[cast.StringI(id)] = count
	}
	return json
}
