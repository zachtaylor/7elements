package match

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"taylz.io/log"
)

type Settings struct {
	Logger log.Writer
	Cards  card.Prototypes
	Decks  deck.Prototypes
	Games  *game.Manager
}

func NewSettings(logger log.Writer, cards card.Prototypes, decks deck.Prototypes, games *game.Manager) Settings {
	return Settings{
		Logger: logger,
		Cards:  cards,
		Decks:  decks,
		Games:  games,
	}
}
