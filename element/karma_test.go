package element_test

import (
	"encoding/json"
	"testing"

	"github.com/zachtaylor/7elements/element"
)

func TestKarmaJSON(t *testing.T) {
	elements := element.Karma{}
	elements[element.Violet] = []bool{false, true, true}
	ans := `{"6":[false,true,true]}`
	if str, _ := json.Marshal(elements.JSON()); string(str) != ans {
		t.Log("Expected(" + ans + ") Actual(" + string(str) + ")")
		t.Fail()
	}
}

func TestKarmaString(t *testing.T) {
	elements := element.Karma{}
	elements[element.Yellow] = []bool{false, false, true}
	elements[element.Violet] = []bool{false, true, true}
	elements[element.Black] = []bool{true}
	ans := `{yyYvVVA}`
	if str := elements.String(); str != ans {
		t.Log("Expected(" + ans + ") Actual(" + str + ")")
		t.Fail()
	}
}
