package accounts_decks

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/zachtaylor/7elements/deck"
	"taylz.io/db"
)

func Update(conn *db.DB, deck *deck.Prototype) (err error) {
	res, e := conn.Exec("UPDATE accounts_decks SET name=?, cover=?, register=? WHERE username=? AND id=?",
		deck.Name,
		deck.Cover,
		time.Now().Unix(),
		deck.User,
		deck.ID,
	)
	if e != nil {
		err = e
	} else if change, e := res.RowsAffected(); e != nil {
		err = e
	} else if change != 1 {
		err = errors.New("rows affected: " + strconv.FormatInt(change, 10))
	}
	return
}

func InsertCards(conn *db.DB, username string, deckid int, diff map[int]int) (err error) {
	sb := strings.Builder{}
	deckidstr := strconv.FormatInt(int64(deckid), 10)
	sb.WriteString("INSERT INTO accounts_decks_items (username, id, cardid, amount) VALUES ")
	first := true
	for k, v := range diff {
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
		sb.WriteString(deckidstr)
		sb.WriteString(",")
		sb.WriteString(strconv.FormatInt(int64(k), 10))
		sb.WriteString(",")
		sb.WriteString(strconv.FormatInt(int64(v), 10))
		sb.WriteString(")")
	}

	_, err = conn.Exec(sb.String())
	return
}
