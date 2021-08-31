package runtime

import (
	"encoding/json"

	"taylz.io/http/websocket"
	"taylz.io/types"
)

func (t *T) GlobalData() types.Bytes {
	if t.glob == nil {
		glob := websocket.MsgData{
			"cards": t.Cards.Data(),
			"packs": t.Packs.JSON(),
			"decks": t.Decks.JSON(),
		}
		t.glob, _ = json.Marshal(glob)
	}
	return t.glob
}
