package deck

import (
	"math/rand"

	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/keygen"
)

var settings = keygen.Settings{
	KeySize: 64,
	CharSet: charset.AlphaCapitalNumeric,
	Rand:    rand.New(rand.NewSource(cast.Now().UnixNano())),
}

// Keygen proposes a new key for the deck key format
func Keygen() string {
	return settings.Keygen()
}
