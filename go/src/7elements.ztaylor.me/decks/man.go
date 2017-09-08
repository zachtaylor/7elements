package decks

import (
	"7elements.ztaylor.me/db"
	"errors"
	"strconv"
	"time"
)

var cache = make(map[string]Decks)

func Store(username string, decks Decks) error {
	if cache[username] != nil {
		return errors.New("decks already stored: " + username)
	}
	cache[username] = decks
	return nil
}

func Test(username string) Decks {
	return cache[username]
}

func Forget(username string) {
	delete(cache, username)
}

func Get(username string) (Decks, error) {
	if cache[username] == nil {
		decks, err := Load(username)

		if err == nil {
			cache[username] = decks
		}

		return decks, err
	}
	return cache[username], nil
}

func Load(username string) (Decks, error) {
	rows, err := db.Connection.Query("SELECT username, name, id, wins, register FROM accounts_decks WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	decks := make(Decks)

	for rows.Next() {
		deck := New()
		var registerbuff int64

		err = rows.Scan(&deck.Username, &deck.Name, &deck.Id, &deck.Wins, &registerbuff)
		if err != nil {
			return nil, err
		}

		deck.Register = time.Unix(registerbuff, 0)
		decks[deck.Id] = deck
	}
	rows.Close()

	rows, err = db.Connection.Query("SELECT id, cardid, amount FROM accounts_decks_items WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var deckid, cardid, amount int

		err = rows.Scan(&deckid, &cardid, &amount)
		if err != nil {
			return nil, err
		}

		if deck := decks[deckid]; deck != nil {
			deck.Cards[cardid] = amount
		} else {
			return nil, errors.New("deckid error #" + strconv.FormatInt(int64(cardid), 10))
		}
	}
	rows.Close()

	return decks, nil
}

func Insert(username string, deckid int) error {
	var insertdecks Decks
	decks := Test(username)

	if decks == nil {
		return errors.New("decks missing")
	}

	if deckid == 0 {
		insertdecks = decks
	} else if deck := decks[deckid]; deck != nil {
		insertdecks = Decks{deckid: deck}
	} else {
		return errors.New("decks missing for username:" + username)
	}

	for deckid, deck := range insertdecks {
		_, err := db.Connection.Exec("INSERT INTO accounts_decks (username, name, id, wins, register) VALUES (?, ?, ?, ?, ?)",
			username,
			deck.Name,
			deck.Id,
			deck.Wins,
			deck.Register.Unix(),
		)
		if err != nil {
			return err
		}

		for cardId, amount := range deck.Cards {
			_, err := db.Connection.Exec("INSERT INTO accounts_decks_items(username, id, cardid, amount) VALUES (?, ?, ?, ?)",
				username,
				deckid,
				cardId,
				amount,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Delete(username string, deckid int) error {
	var deletedecks []int
	decks := Test(username)

	if decks == nil {
		return errors.New("accounts_decks: account has no decks: " + username)
	}

	if deckid == 0 {
		for i, _ := range decks {
			deletedecks = append(deletedecks, i)
		}
	} else if deck := decks[deckid]; deck != nil {
		deletedecks = []int{deck.Id}
	} else {
		return errors.New("accounts_decks: delete 404: " + username)
	}

	for _, deckid := range deletedecks {
		_, err := db.Connection.Exec("DELETE FROM accounts_decks WHERE username=? AND id=?",
			username,
			deckid,
		)

		if err != nil {
			return err
		}

		_, err = db.Connection.Exec("DELETE FROM accounts_decks_items WHERE username=? AND id=?",
			username,
			deckid,
		)

		if err != nil {
			return err
		}

	}

	return nil
}
