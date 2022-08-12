package apiws

import (
	"reflect"
	"strconv"

	"github.com/zachtaylor/7elements/db/accounts_decks"
	"github.com/zachtaylor/7elements/db/accounts_decks_items"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func UpdateDeck(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		user := server.Users().GetWebsocket(socket)
		if user == nil {
			log.Warn("no session")
			socket.Write(websocket.MessageText, out.Error("vii", "no user"))
			return
		}
		log = log.Add("Session", user.Session())

		account := server.Accounts().Get(user.Session().Name())
		if account == nil {
			log.Error("no account")
			socket.Write(websocket.MessageText, out.Error("vii", "no account"))
			return
		}

		var id int
		if idbuff, _ := m.Data["id"].(float64); idbuff < 1 {
			log.Add("id", m.Data["id"]).Add("type", reflect.TypeOf(m.Data["id"])).Warn("id missing")
			socket.Write(websocket.MessageText, out.Error("vii", "no deck id"))
			return
		} else {
			id = int(idbuff)
		}
		log = log.Add("DeckID", id)

		deck := account.Decks[id]
		if deck == nil {
			log.Warn("invalid id")
			socket.Write(websocket.MessageText, out.Error("vii", "invalid deck id"))
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

		if diff, _ := m.Data["cards"].(map[string]any); diff == nil {
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

		if err := accounts_decks.Update(server.DB(), deck); err != nil {
			log.Add("Error", err).Error("update")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}

		if itemsDirty {
			if err := accounts_decks_items.Delete(server.DB(), account.Username, id); err != nil {
				log.Add("Error", err).Error("delete items")
				socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			}
			if err := accounts_decks_items.Insert(server.DB(), account.Username, id, deck.Cards); err != nil {
				log.Add("Error", err).Error("insert items")
				socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			}
		}

		socket.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))
	})
}
