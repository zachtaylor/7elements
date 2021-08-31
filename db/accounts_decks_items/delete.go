package accounts_decks_items

import "taylz.io/db"

func Delete(conn *db.DB, username string, deckid int) (err error) {
	_, err = conn.Exec("DELETE FROM accounts_decks_items WHERE username=? AND id=?",
		username,
		deckid,
	)
	return
}
