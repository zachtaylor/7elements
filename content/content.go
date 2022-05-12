package content

import (
	"encoding/json"
	"errors"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/card/pack"
	"github.com/zachtaylor/7elements/db/cards"
	"github.com/zachtaylor/7elements/db/decks"
	"github.com/zachtaylor/7elements/db/packs"
	"github.com/zachtaylor/7elements/deck"
	"taylz.io/db"
)

// T is content data
type T struct {
	cards card.Prototypes
	decks deck.Prototypes
	packs pack.Prototypes
	data  []byte
}

func (*T) Version() string          { return vii.Version }
func (t *T) Cards() card.Prototypes { return t.cards }
func (t *T) Decks() deck.Prototypes { return t.decks }
func (t *T) Packs() pack.Prototypes { return t.packs }
func (t *T) Data() []byte           { return t.data }

func Build(db *db.DB) (*T, error) {
	cards := cards.GetAll(db)
	if cards == nil {
		return nil, errors.New("failed to load cards")
	}
	decks, err := decks.GetAll(db)
	if decks == nil {
		return nil, err
	}
	packs, err := packs.GetAll(db)
	if packs == nil {
		return nil, err
	}
	glob, err := json.Marshal(map[string]any{
		"cards": cards.Data(),
		"decks": decks.Data(),
		"packs": packs.Data(),
	})
	if err != nil {
		return nil, err
	}

	return &T{
		cards: cards,
		decks: decks,
		packs: packs,
		data:  glob,
	}, nil
}
