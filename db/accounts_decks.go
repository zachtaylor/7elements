package db

import (
	"time"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/log"
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
	rows, err := conn.Query(
		"SELECT username, id, name, wins, color, max(register) FROM accounts_decks WHERE username=? GROUP BY id",
		username,
	)

	if err != nil {
		return nil, err
	}

	decks := make(vii.AccountDecks, 0)

	for rows.Next() {
		deck := vii.NewAccountDeck()
		var registerbuff int64

		err = rows.Scan(&deck.Username, &deck.ID, &deck.Name, &deck.Wins, &deck.Color, &registerbuff)
		if err != nil {
			return nil, err
		}

		deck.Register = time.Unix(registerbuff, 0)
		decks = append(decks, deck)
	}
	rows.Close()

	for _, deck := range decks {
		rows, err = conn.Query("SELECT cardid, amount FROM accounts_decks_items WHERE username=? AND id=?",
			username,
			deck.ID,
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
	deck.Wins = 0
	deck.Register = time.Now()

	if err := service.Delete(deck.Username, deck.ID); err != nil {
		log.Error("cannot delete username, deckid")
		return err
	} else if err := service.Insert(deck); err != nil {
		log.Error("cannot insert deck")
		return err
	}

	return nil
}

func (_ AccountDeckService) Insert(deck *vii.AccountDeck) error {
	_, err := conn.Exec("INSERT INTO accounts_decks (username, name, id, wins, register, color) VALUES (?, ?, ?, ?, ?, ?)",
		deck.Username,
		deck.Name,
		deck.ID,
		deck.Wins,
		deck.Register.Unix(),
		deck.Color,
	)
	if err != nil {
		return err
	}

	for cardId, amount := range deck.Cards {
		_, err := conn.Exec("INSERT INTO accounts_decks_items(username, id, cardid, amount) VALUES (?, ?, ?, ?)",
			deck.Username,
			deck.ID,
			cardId,
			amount,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (_ AccountDeckService) UpdateName(username string, id int, name string) error {
	res, err := conn.Exec(
		"UPDATE accounts_decks SET name=? WHERE username=? AND id=?;",
		name,
		username,
		id,
	)

	if err != nil {
		return err
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected != 1 {
		return ErrUpdateFailed
	}

	return nil
}

func (_ AccountDeckService) UpdateTallyWinCount(username string, id int) error {
	res, err := conn.Exec(
		"UPDATE accounts SET wins=wins+1 WHERE username=? AND id=?",
		username,
		id,
	)

	if err != nil {
		return err
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected != 1 {
		return ErrUpdateFailed
	}

	return nil
}

func (service AccountDeckService) Delete(username string, deckid int) error {
	_, err := conn.Exec("DELETE FROM accounts_decks WHERE username=? AND id=?",
		username,
		deckid,
	)
	if err != nil {
		log.Add("Error", err).Error("cannot delete accounts_decks")
		return err
	}
	_, err = conn.Exec("DELETE FROM accounts_decks_items WHERE username=? AND id=?",
		username,
		deckid,
	)
	return err
}
