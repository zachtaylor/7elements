package accounts_decks_items

import (
	"strconv"
	"strings"

	"github.com/zachtaylor/7elements/card"
	"taylz.io/db"
)

func Insert(conn *db.DB, username string, deckid int, cards card.Count) (err error) {
	sb := strings.Builder{}
	sb.WriteString("INSERT INTO accounts_decks_items (username, id, cardid, amount) VALUES ")
	first := true
	for k, v := range cards {
		if v < 1 {
			continue
		}
		if first {
			first = false
		} else {
			sb.WriteString(",")
		}
		sb.WriteString("('")
		sb.WriteString(username)
		sb.WriteString("',")
		sb.WriteString(strconv.FormatInt(int64(deckid), 10))
		sb.WriteString(",")
		sb.WriteString(strconv.FormatInt(int64(k), 10))
		sb.WriteString(",")
		sb.WriteString(strconv.FormatInt(int64(v), 10))
		sb.WriteString(")")
	}

	_, err = conn.Exec(sb.String())
	return
}
