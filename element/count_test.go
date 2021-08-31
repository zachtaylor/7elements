package element_test

import (
	"testing"

	"github.com/zachtaylor/7elements/element"
)

func TestCountString(t *testing.T) {
	c := element.Count{
		element.White: 3,
	}

	if str := c.String(); str != "WWW" {
		t.Log("Expected", "WWW")
		t.Log("Actual", str)
		t.Fail()
	}
}
