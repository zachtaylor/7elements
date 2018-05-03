package dbservice

import (
	"elemen7s.com"
	"elemen7s.com/db"
	"time"
)

func init() {
	vii.AccountDeckService = make(AccountDeckService)
}

type AccountDeckService map[string]vii.AccountDecks

func (service AccountDeckService) Get(username string) (vii.AccountDecks, error) {
	if service[username] == nil {
		if data, err := service.Load(username); err != nil {
			return nil, err
		} else {
			service[username] = data
		}
	}
	return service[username], nil
}

func (service AccountDeckService) Load(username string) (vii.AccountDecks, error) {
	rows, err := db.Connection.Query(
		"SELECT username, id, version, name, wins, color, max(register) FROM accounts_decks WHERE username=? GROUP BY id",
		username,
	)

	if err != nil {
		return nil, err
	}

	decks := make(vii.AccountDecks, 0)

	for rows.Next() {
		deck := vii.NewAccountDeck()
		var registerbuff int64

		err = rows.Scan(&deck.Username, &deck.Id, &deck.Version, &deck.Name, &deck.Wins, &deck.Color, &registerbuff)
		if err != nil {
			return nil, err
		}

		deck.Register = time.Unix(registerbuff, 0)
		decks = append(decks, deck)
	}
	rows.Close()

	for _, deck := range decks {
		rows, err = db.Connection.Query("SELECT cardid, amount FROM accounts_decks_items WHERE username=? AND id=? AND version=?",
			username,
			deck.Id,
			deck.Version,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var cardid, amount int

			err = rows.Scan(&cardid, &amount)
			if err != nil {
				return nil, err
			}

			deck.Cards[cardid] = amount
		}
		rows.Close()
	}

	return decks, nil
}

func (service AccountDeckService) Forget(username string) {
	delete(service, username)
}

func (service AccountDeckService) Update(deck *vii.AccountDeck) error {
	for !service.checkUniqueDeckIdVersion(deck.Username, deck.Id, deck.Version) {
		deck.Version = vii.NewKey()
	}
	deck.Wins = 0
	deck.Register = time.Now()

	if err := service.Insert(deck); err != nil {
		return err
	}

	return nil
}

func (_ AccountDeckService) Insert(deck *vii.AccountDeck) error {
	_, err := db.Connection.Exec("INSERT INTO accounts_decks (username, name, id, version, wins, register, color) VALUES (?, ?, ?, ?, ?, ?, ?)",
		deck.Username,
		deck.Name,
		deck.Id,
		deck.Version,
		deck.Wins,
		deck.Register.Unix(),
		deck.Color,
	)
	if err != nil {
		return err
	}

	for cardId, amount := range deck.Cards {
		_, err := db.Connection.Exec("INSERT INTO accounts_decks_items(username, id, version, cardid, amount) VALUES (?, ?, ?, ?, ?)",
			deck.Username,
			deck.Id,
			deck.Version,
			cardId,
			amount,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (_ AccountDeckService) checkUniqueDeckIdVersion(username string, deckid int, version string) bool {
	row := db.Connection.QueryRow(
		"SELECT register FROM accounts_decks WHERE username=? AND deckid=? AND version=?",
		username,
		deckid,
		version,
	)

	var buff int64
	if err := row.Scan(&buff); err != nil {
		return true
	}
	return buff > 0
}

func (_ AccountDeckService) UpdateName(username string, id int, name string) error {
	res, err := db.Connection.Exec(
		"UPDATE accounts_decks SET name=? WHERE username=? AND id=? AND version = (SELECT max(version) FROM accounts_decks WHERE id=?)",
		name,
		username,
		id,
		id,
	)

	if err != nil {
		return err
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected != 1 {
		return ERRupdate_failed
	}

	return nil
}

func (_ AccountDeckService) UpdateTallyWinCount(username string, id int, version string) error {
	res, err := db.Connection.Exec(
		"UPDATE accounts SET wins=wins+1 WHERE username=? AND id=? AND version=?",
		username,
		id,
		version,
	)

	if err != nil {
		return err
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected != 1 {
		return ERRupdate_failed
	}

	return nil
}

// func (_ AccountDeckService) Delete(username string, deckid int) error {
// 	var deletedecks []int
// 	decks := Test(username)

// 	if decks == nil {
// 		return errors.New("accounts_decks: account has no decks: " + username)
// 	}

// 	if deckid == 0 {
// 		for i, _ := range decks {
// 			deletedecks = append(deletedecks, i)
// 		}
// 	} else if deck := decks[deckid]; deck != nil {
// 		deletedecks = []int{deck.Id}
// 	} else {
// 		return errors.New("accounts_decks: delete 404: " + username)
// 	}

// 	for _, deckid := range deletedecks {
// 		_, err := db.Connection.Exec("DELETE FROM accounts_decks WHERE username=? AND id=?",
// 			username,
// 			deckid,
// 		)

// 		if err != nil {
// 			return err
// 		}

// 		_, err = db.Connection.Exec("DELETE FROM accounts_decks_items WHERE username=? AND id=?",
// 			username,
// 			deckid,
// 		)

// 		if err != nil {
// 			return err
// 		}

// 	}

// 	return nil
// }
