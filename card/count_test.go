package card_test

import (
	"testing"

	"github.com/zachtaylor/7elements/card"
)

func TestCount(t *testing.T) {
	c := card.Count{
		1: 1,
		4: 4,
		7: 7,
	}
	var cfmt string
	if fmt, err := c.Format(); err != nil {
		t.Log("format error", err)
		t.Fail()
		return
	} else {
		cfmt = fmt
	}
	cans := `1a4d7g`
	if cfmt != cans {
		t.Log("expected", cans)
		t.Log("actual", cfmt)
		t.Fail()
		return
	}
	if copy, err := card.ParseCount(cfmt); err != nil {
		t.Log("parse error", err)
		t.Fail()
		return
	} else {
		for k, v := range copy {
			if c[k] != v {
				t.Log("expected", c[k])
				t.Log("actual", v)
			}
		}
		for k, v := range c {
			if copy[k] != v {
				t.Log("expected", v)
				t.Log("actual", copy[k])
			}
		}

		if fmt, err := copy.Format(); err != nil {
			t.Log("format error", err)
			t.Fail()
		} else {
			// in n out
			if cfmt != fmt {
				t.Log("expected", cans)
				t.Log("actual", cfmt)
				t.Fail()
			}
		}
	}
}
