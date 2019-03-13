package ai

import (
	"math/rand"
	"time"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

func Register(game *game.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	decks, err := vii.DeckService.GetAll()
	if err != nil {
		log.Add("Error", err).Error("ai/register: decks")
		return
	}
	i := (r.Int() % len(decks)) + 1
	deck := vii.NewAccountDeckWith(decks[i], "A.I.")
	seat := game.Register(deck)
	ConnectAI(game, seat)
	log.Add("Deck", deck.Name).Info("ai/register: decks")
}
