package card

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"taylz.io/types"
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
		v := strings.Index(types.CharsetNumeric, string([]byte{format[i]}))
		k := 1 + strings.Index(types.CharsetAlphaCapital, string([]byte{format[i+1]}))
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
			return "", errors.New("val too low: " + strconv.FormatInt(int64(v), 10))
		} else if v > 9 {
			return "", errors.New("val too high: " + strconv.FormatInt(int64(v), 10))
		} else if k < 1 {
			return "", errors.New("key too low: " + strconv.FormatInt(int64(k), 10))
		} else if k > 52 {
			return "", errors.New("key too high: " + strconv.FormatInt(int64(k), 10))
		}
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	b := strings.Builder{}
	b.Grow(len(c) * 2)
	for i = 0; i < len(keys); i++ {
		b.WriteByte(types.CharsetNumeric[c[keys[i]]])
		b.WriteByte(types.CharsetAlphaCapital[keys[i]-1])
	}
	return b.String(), nil
}

// Data returns a representation of this count of Cards as type map[string]any
func (c Count) Data() map[string]any {
	json := map[string]any{}
	for id, count := range c {
		json[strconv.FormatInt(int64(id), 10)] = count
	}
	return json
}
