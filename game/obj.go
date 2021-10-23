package game

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game/token"
	"taylz.io/keygen"
)

func (t *T) objSave(i interface{}) (key string) {
	for ok := true; ok; _, ok = t.obj[key] {
		key = keygen.New(7)
	}
	t.obj[key] = i
	return
}

func (t *T) RegisterState(state *State) {
	state.id = t.objSave(state)
}

func (t *T) RegisterCard(card *card.T) {
	card.ID = t.objSave(card)
}

func (t *T) RegisterToken(token *token.T) {
	token.ID = t.objSave(token)
}

func (t *T) GetState(key string) *State {
	if state, ok := t.obj[key].(*State); ok {
		return state
	}
	return nil
}

func (t *T) GetCard(key string) *card.T {
	if card, ok := t.obj[key].(*card.T); ok {
		return card
	}
	return nil
}

func (t *T) GetToken(key string) *token.T {
	if token, ok := t.obj[key].(*token.T); ok {
		return token
	}
	return nil
}
