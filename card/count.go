package card

import (
	"errors"
	"sort"
	"strings"

	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
)

// Count is a map(cardid->quantity)
type Count map[int]int

// ParseCount returns a count from a custom format
func ParseCount(format string) (Count, error) {
	lenfmt := len(format)
	if lenfmt%2 != 0 {
		return nil, errors.New("parse count format must have even length")
	}
	count := Count{}
	for i := 0; i < lenfmt; i += 2 {
		v := strings.Index(charset.Numeric, string([]byte{format[i]}))
		k := 1 + strings.Index(charset.AlphaCapital, string([]byte{format[i+1]}))
		count[k] = v
	}
	return count, nil
}

// Count returns the total number of cards listed
func (c Count) Count() int {
	total := 0
	for _, count := range c {
		total += count
	}
	return total
}

// Format writes this Count to a custom format
func (c Count) Format() (string, error) {
	keys := make([]int, len(c))
	i := 0
	for k, v := range c {
		if v < 1 {
			return "", errors.New("val too low: " + cast.StringI(v))
		} else if v > 9 {
			return "", errors.New("val too high: " + cast.StringI(v))
		} else if k < 1 {
			return "", errors.New("key too low: " + cast.StringI(k))
		} else if k > 52 {
			return "", errors.New("key too high: " + cast.StringI(k))
		}
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	b := strings.Builder{}
	b.Grow(len(c) * 2)
	for i = 0; i < len(keys); i++ {
		b.WriteByte(charset.Numeric[c[keys[i]]])
		b.WriteByte(charset.AlphaCapital[keys[i]-1])
	}
	return b.String(), nil
}

// JSON returns a representation of this count of Cards as type cast.JSON
func (c Count) JSON() cast.JSON {
	json := cast.JSON{}
	for id, count := range c {
		json[cast.StringI(id)] = count
	}
	return json
}
