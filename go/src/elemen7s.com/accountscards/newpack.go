package accountscards

import (
	"github.com/cznic/mathutil"
	"time"
)

var cardRandomFC32s [7]*mathutil.FC32

func init() {
	var seeds = []int64{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(cardRandomFC32s); i++ {
		cardRandomFC32s[i], _ = mathutil.NewFC32(1, 50, true)
		cardRandomFC32s[i].Seed(seeds[i])
	}
}

func NewPack(username string) []*AccountCard {
	pack := make([]*AccountCard, 7)
	register := time.Now()

	for i, cardRandomFC32 := range cardRandomFC32s {
		cardid := int(cardRandomFC32.Next())
		for ; checkInPack(pack, cardid); cardid = int(cardRandomFC32.Next()) {
		}
		pack[i] = &AccountCard{
			Username: username,
			CardId:   cardid,
			Register: register,
		}
	}
	return pack
}

func checkInPack(pack []*AccountCard, id int) bool {
	for _, card := range pack {
		if card != nil && card.CardId == id {
			return true
		}
	}
	return false
}
