package gencardpack

import (
	"time"

	"github.com/cznic/mathutil"
	vii "github.com/zachtaylor/7elements"
)

var rand, _ = mathutil.NewFC32(1, 98, true)

func NewPack(rt *vii.Runtime, username string, pack *vii.Pack) []*vii.AccountCard {
	cards := make([]*vii.AccountCard, pack.Size)
	register := time.Now()
	packDelta := len(pack.Cards)

	for i := 0; i < pack.Size; i++ {
		cardid := 0
		for ok := true; ok; ok = checkInPack(cards, cardid) {
			cardid = pack.Cards[int(rand.Next())%packDelta].CardID
		}
		cards[i] = &vii.AccountCard{
			Username: username,
			CardId:   cardid,
			Register: register,
		}
	}

	return cards
}

func checkInPack(pack []*vii.AccountCard, id int) bool {
	for _, card := range pack {
		if card != nil && card.CardId == id {
			return true
		}
	}
	return false
}
