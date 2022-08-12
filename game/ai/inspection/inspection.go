package inspection

import "github.com/zachtaylor/7elements/game"

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

func Parse(game *game.G, player *game.Player) (t T) {
	for tokenID := range player.T.Present {
		token := game.Token(tokenID)
		if token.T.Body == nil {
			t.Items++
			if token.T.Awake {
				t.AwakeItems++
			}
		} else {
			t.Beings++
			t.BeingsAttack += token.T.Body.Attack
			t.BeingsLife += token.T.Body.Life
			if token.T.Awake {
				t.AwakeBeings++
				t.AwakeBeingsAttack += token.T.Body.Attack
				t.AwakeBeingsLife += token.T.Body.Life
			}
		}
	}
	return
}
