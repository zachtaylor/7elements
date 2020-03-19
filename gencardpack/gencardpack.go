package gencardpack

import (
	"time"

	"github.com/cznic/mathutil"
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/account"
)

var rand, _ = mathutil.NewFC32(1, 98, true)

func NewPack(rt *vii.Runtime, username string, pack *vii.Pack) []*account.Card {
	cards := make([]*account.Card, pack.Size)
	register := time.Now()
	packDelta := len(pack.Cards)

	for i := 0; i < pack.Size; i++ {
		cardid := 0
		for ok := true; ok; ok = checkInPack(cards, cardid) {
			cardid = pack.Cards[int(rand.Next())%packDelta].CardID
		}
		cards[i] = &account.Card{
			Username: username,
			ProtoID:  cardid,
			Register: register,
		}
	}

	return cards
}

func checkInPack(pack []*account.Card, id int) bool {
	for _, card := range pack {
		if card != nil && card.ProtoID == id {
			return true
		}
	}
	return false
}
