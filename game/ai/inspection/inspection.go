package inspection

import (
	"github.com/zachtaylor/7elements/game/seat"
)

type T struct {
	Items             int
	AwakeItems        int
	Beings            int
	AwakeBeings       int
	BeingsAttack      int
	AwakeBeingsAttack int
	BeingsLife        int
	AwakeBeingsLife   int
}

func Parse(seat *seat.T) (t T) {
	for _, token := range seat.Present {
		if token.Body == nil {
			t.Items++
			if token.IsAwake {
				t.AwakeItems++
			}
		} else {
			t.Beings++
			t.BeingsAttack += token.Body.Attack
			t.BeingsLife += token.Body.Life
			if token.IsAwake {
				t.AwakeBeings++
				t.AwakeBeingsAttack += token.Body.Attack
				t.AwakeBeingsLife += token.Body.Life
			}
		}
	}
	return
}
