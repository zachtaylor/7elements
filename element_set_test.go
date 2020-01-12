package vii_test

import (
	"testing"

	vii "github.com/zachtaylor/7elements"
)

func TestElementSetJSON(t *testing.T) {
	elements := vii.ElementSet{}
	elements[vii.ELEMviolet] = []bool{true, true, false}
	ans := `{"6":[true,true,false]}`
	if str := elements.JSON().String(); str != ans {
		t.Log("Expected(" + ans + ") Actual(" + str + ")")
		t.Fail()
	}
}
