package element_test

import (
	"testing"

	"github.com/zachtaylor/7elements/element"
)

func TestParseCount(t *testing.T) {
	tests := []string{
		`GV`,
		`XWW`,
		`XXB`,
	}
	for _, ans := range tests {
		c := make(element.Count)

		t.Log("Parse string", ans)
		for _, char := range ans {
			t.Log("parse char", char, "as byte", byte(char))
			element, _, err := element.Parse(byte(char))
			if err != nil {
				t.Log("Error:", err)
				t.Fail()
			}
			c[element]++
		}

		t.Log("Completed", c, `"`+c.String()+`"`)

		if str := c.String(); str != ans {
			t.Log("Expected(" + ans + ") Actual(" + str + ")")
			t.Fail()
		}
	}
}

func TestParseKarma(t *testing.T) {
	tests := []string{
		`{xXgg}`,
		`{rRyYGVa}`,
		`{wWbBvA}`,
	}
	for _, ans := range tests {
		karma, _ := element.ParseKarma(ans)
		if str := karma.String(); str != ans {
			t.Log("Expected(" + ans + ") Actual(" + str + ")")
			t.Fail()
		}
	}
}
