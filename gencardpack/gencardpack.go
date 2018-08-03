package gencardpack

import (
	"github.com/zachtaylor/7tcg"
	"github.com/cznic/mathutil"
	"time"
)

var rand [7]*mathutil.FC32

func init() {
	var seeds = []int64{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(rand); i++ {
		rand[i], _ = mathutil.NewFC32(1, 50, true)
		rand[i].Seed(seeds[i])
	}
}

func NewPack(username string) []*vii.AccountCard {
	pack := make([]*vii.AccountCard, 7)
	register := time.Now()

	for i, cardRandomFC32 := range rand {
		cardid := int(cardRandomFC32.Next())
		for ; checkInPack(pack, cardid); cardid = int(cardRandomFC32.Next()) {
		}
		pack[i] = &vii.AccountCard{
			Username: username,
			CardId:   cardid,
			Register: register,
		}
	}
	return pack
}

func checkInPack(pack []*vii.AccountCard, id int) bool {
	for _, card := range pack {
		if card != nil && card.CardId == id {
			return true
		}
	}
	return false
}
