package accounts_decks

import (
	"strconv"
	"strings"

	"github.com/zachtaylor/7elements/deck"
	"taylz.io/db"
)

func InsertAll(conn *db.DB, decks deck.Prototypes) (err error) {
	statement := strings.Builder{}
	statement.WriteString("INSERT INTO accounts_decks (username, id, name, cover) VALUES ")
	first := true
	for _, proto := range decks {
		if !first {
			statement.WriteString(", ")
		} else {
			first = false
		}
		statement.WriteString("('")
		statement.WriteString(proto.User)
		statement.WriteString("',")
		statement.WriteString(strconv.FormatInt(int64(proto.ID), 10))
		statement.WriteString(",'")
		statement.WriteString(proto.Name)
		statement.WriteString("',")
		statement.WriteString(strconv.FormatInt(int64(proto.Cover), 10))
		statement.WriteByte(')')
	}

	// if _, err = conn.Exec(str); err != nil {
	// 	return
	// }

	// statement.Reset()
	// statement.WriteString("INSERT INTO accounts_decks_items (username, id, cardid, amount) VALUES ")
	// first = true

	// for _, proto := range decks {
	// 	for k, v := range proto.Cards {
	// 		if !first {
	// 			statement.WriteString(", ")
	// 		} else {
	// 			first = false
	// 		}

	// 		statement.WriteString("('")
	// 		statement.WriteString(proto.User)
	// 		statement.WriteString("',")
	// 		statement.WriteString(types.StringInt(proto.ID))
	// 		statement.WriteString(",")
	// 		statement.WriteString(types.StringInt(k))
	// 		statement.WriteString(",")
	// 		statement.WriteString(types.StringInt(v))
	// 		statement.WriteByte(')')
	// 	}
	// }

	// str = statement.String()

	// fmt.Println(str)

	_, err = conn.Exec(statement.String())
	return
}
