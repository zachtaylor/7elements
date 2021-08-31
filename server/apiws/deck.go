package apiws

import (
	"reflect"
	"strconv"

	"github.com/zachtaylor/7elements/db/accounts_decks"
	"github.com/zachtaylor/7elements/db/accounts_decks_items"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func UpdateDeck(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		updateDeck(rt, socket, m)
	})
}

func updateDeck(rt *runtime.T, socket *websocket.T, m *websocket.Message) {
	log := rt.Log()

	if socket.Name() == "" {
		log.Add("Socket", socket.ID()).Warn("no user")
		socket.Write(wsout.ErrorJSON("vii", "no user"))
		return
	}

	account := rt.Accounts.Get(socket.Name())
	if account == nil {
		log.Add("User", socket.Name()).Error("no account")
		socket.Write(wsout.ErrorJSON("vii", "no account"))
		return
	}
	log = log.Add("User", account.Username)

	var id int
	if idbuff, _ := m.Data["id"].(float64); idbuff < 1 {
		log.Add("id", m.Data["id"]).Add("type", reflect.TypeOf(m.Data["id"])).Warn("id missing")
		socket.Write(wsout.ErrorJSON("vii", "no deck id"))
		return
	} else {
		id = int(idbuff)
	}
	log = log.Add("ID", id)

	deck := account.Decks[id]
	if deck == nil {
		log.Warn("invalid id")
		socket.Write(wsout.ErrorJSON("vii", "invalid deck id"))
		return
	}

	if newname, _ := m.Data["name"].(string); newname != "" {
		log.Add("Name", newname).Trace("newname")
		deck.Name = newname
	}
	if newcover, _ := m.Data["cover"].(float64); int(newcover) > 0 {
		log.Add("Cover", newcover).Trace("newcover")
		deck.Cover = int(newcover)
	}

	var itemsDirty bool

	if diff, _ := m.Data["cards"].(map[string]interface{}); diff == nil {
		if val := m.Data["cards"]; val != nil {
			log = log.Add("Cards", val).Add("Type", reflect.TypeOf(val))
		}
		log.Warn("cards not found")
	} else {
		for k, v := range diff {
			var cardid, change int

			if cardidbuff, err := strconv.ParseInt(k, 10, 0); err != nil {
				log.Add("Error", err.Error()).Warn("cardid key: ", cardid)
				continue
			} else {
				cardid = int(cardidbuff)
			}

			if changebuff, ok := v.(float64); !ok {
				log.Add("ChangeV", v).Add("Type", reflect.TypeOf(v)).Warn("card change value not found")
				continue
			} else {
				change = int(changebuff)
			}

			total := deck.Cards[cardid] + change
			if account.Cards[cardid] < total {
				log.Add("CardID", cardid).Add("Change", change).Add("Total", total).Add("Owned", account.Cards[cardid]).Warn("exceed account limit")
			} else if deck.Cards[cardid]+change < 0 {
				log.Add("CardID", cardid).Add("Change", change).Warn("illegal value")
			} else {
				itemsDirty = true
				deck.Cards[cardid] = total
			}
		}
	}

	if err := accounts_decks.Update(rt.DB, deck); err != nil {
		log.Add("Error", err).Error("update")
		socket.Write(wsout.ErrorJSON("vii", "internal error"))
		return
	}

	if itemsDirty {
		if err := accounts_decks_items.Delete(rt.DB, account.Username, id); err != nil {
			log.Add("Error", err).Error("delete items")
			socket.Write(wsout.ErrorJSON("vii", "internal error"))
		}
		if err := accounts_decks_items.Insert(rt.DB, account.Username, id, deck.Cards); err != nil {
			log.Add("Error", err).Error("insert items")
			socket.Write(wsout.ErrorJSON("vii", "internal error"))
		}
	}

	socket.Write(wsout.MyAccount(account.Data()).EncodeToJSON())
}
