package vii_test

import (
	"testing"

	vii "github.com/zachtaylor/7elements"
)

func TestKarmaJSON(t *testing.T) {
	elements := vii.Karma{}
	elements[vii.ELEMviolet] = []bool{false, true, true}
	ans := `{"6":[false,true,true]}`
	if str := elements.JSON().String(); str != ans {
		t.Log("Expected(" + ans + ") Actual(" + str + ")")
		t.Fail()
	}
}

func TestKarmaString(t *testing.T) {
	elements := vii.Karma{}
	elements[vii.ELEMyellow] = []bool{false, false, true}
	elements[vii.ELEMviolet] = []bool{false, true, true}
	elements[vii.ELEMblack] = []bool{true}
	ans := `{yyYvVVA}`
	if str := elements.String(); str != ans {
		t.Log("Expected(" + ans + ") Actual(" + str + ")")
		t.Fail()
	}
}

func TestParseKarma(t *testing.T) {
	tests := []string{
		`{xXgg}`,
		`{rRyYGVa}`,
		`{wWbBvA}`,
	}
	for _, ans := range tests {
		karma, _ := vii.ParseKarma(ans)
		if str := karma.String(); str != ans {
			t.Log("Expected(" + ans + ") Actual(" + str + ")")
			t.Fail()
		}
	}
}
