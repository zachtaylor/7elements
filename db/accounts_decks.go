package db

import (
	"time"

	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/db"
)

func NewAccountDeckService(db *db.DB) vii.AccountDeckService {
	return &AccountDeckService{
		conn:  db,
		cache: make(map[string]vii.AccountDecks),
	}
}

type AccountDeckService struct {
	conn  *db.DB
	cache map[string]vii.AccountDecks
}

func (ads *AccountDeckService) Find(username string) (vii.AccountDecks, error) {
	if ads.cache[username] == nil {
		if data, err := ads.Get(username); err != nil {
			return nil, err
		} else {
			ads.cache[username] = data
		}
	}
	return ads.cache[username], nil
}

func (ads *AccountDeckService) Get(username string) (vii.AccountDecks, error) {
	rows, err := ads.conn.Query(
		"SELECT username, id, name, wins, cover, max(register) FROM accounts_decks WHERE username=? GROUP BY id",
		username,
	)

	if err != nil {
		return nil, err
	}

	decks := make(vii.AccountDecks, 0)

	for rows.Next() {
		deck := vii.NewAccountDeck()
		var registerbuff int64

		err = rows.Scan(&deck.Username, &deck.ID, &deck.Name, &deck.Wins, &deck.CoverID, &registerbuff)
		if err != nil {
			return nil, err
		}

		deck.Register = time.Unix(registerbuff, 0)
		decks = append(decks, deck)
	}
	rows.Close()

	for _, deck := range decks {
		rows, err = ads.conn.Query("SELECT cardid, amount FROM accounts_decks_items WHERE username=? AND id=?",
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

func (ads *AccountDeckService) Forget(username string) {
	delete(ads.cache, username)
}

func (ads *AccountDeckService) Update(deck *vii.AccountDeck) (err error) {
	deck.Wins = 0
	deck.Register = time.Now()
	if err = ads.Delete(deck.Username, deck.ID); err != nil {
	} else {
		err = ads.Insert(deck)
	}
	return
}

func (ads *AccountDeckService) Insert(deck *vii.AccountDeck) error {
	_, err := ads.conn.Exec("INSERT INTO accounts_decks (username, name, id, wins, register, cover) VALUES (?, ?, ?, ?, ?, ?)",
		deck.Username,
		deck.Name,
		deck.ID,
		deck.Wins,
		deck.Register.Unix(),
		deck.CoverID,
	)
	if err != nil {
		return err
	}

	for cardId, amount := range deck.Cards {
		_, err := ads.conn.Exec("INSERT INTO accounts_decks_items(username, id, cardid, amount) VALUES (?, ?, ?, ?)",
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

func (ads *AccountDeckService) UpdateName(username string, id int, name string) error {
	res, err := ads.conn.Exec(
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

func (ads *AccountDeckService) UpdateTallyWinCount(username string, id int) error {
	res, err := ads.conn.Exec(
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

func (ads *AccountDeckService) Delete(username string, deckid int) (err error) {
	_, err = ads.conn.Exec("DELETE FROM accounts_decks WHERE username=? AND id=?",
		username,
		deckid,
	)
	if err == nil {
		_, err = ads.conn.Exec("DELETE FROM accounts_decks_items WHERE username=? AND id=?",
			username,
			deckid,
		)
	}
	return
}
