package persistence

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"errors"
	"time"
)

func AccountsDecksGet(username string) (map[int]*SE.AccountDeck, error) {
	rows, err := connection.Query("SELECT name, id, wins, register FROM accounts_decks WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	decks := make(map[int]*SE.AccountDeck)

	for rows.Next() {
		deck := SE.NewAccountDeck()
		var registerbuff int64

		err = rows.Scan(&deck.Name, &deck.Id, &deck.Wins, &registerbuff)
		if err != nil {
			return nil, err
		}

		deck.Register = time.Unix(registerbuff, 0)
		decks[deck.Id] = deck
	}
	rows.Close()

	rows, err = connection.Query("SELECT id, cardid, amount FROM accounts_decks_items WHERE username=?",
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
			return nil, errors.New("accounts_decks: found item for missing deck")
		}
	}
	rows.Close()

	log.Add("Username", username).Debug("accounts_decks: get")

	return decks, nil
}

func AccountsDecksInsert(username string, deckid int) error {
	var insertdecks map[int]*SE.AccountDeck
	accountsdecks := SE.AccountsDecks.Cache[username]

	if accountsdecks == nil {
		return errors.New("accounts_decks: account has no decks: " + username)
	}

	if deckid == 0 {
		insertdecks = SE.AccountsDecks.Cache[username]
	} else if deck := accountsdecks[deckid]; deck != nil {
		insertdecks = map[int]*SE.AccountDeck{deck.Id: deck}
	} else {
		return errors.New("accounts_decks: insert 404: " + username)
	}

	for deckid, deck := range insertdecks {
		_, err := connection.Exec("INSERT INTO accounts_decks (username, name, id, wins, register) VALUES (?, ?, ?, ?, ?)",
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
			_, err := connection.Exec("INSERT INTO accounts_decks_items(username, id, cardid, amount) VALUES (?, ?, ?, ?)",
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

	log.Add("Username", username).Debug("accounts_decks: insert")

	return nil
}

func AccountsDecksDelete(username string, deckid int) error {
	var deletedecks []int
	accountsdecks := SE.AccountsDecks.Cache[username]

	if accountsdecks == nil {
		return errors.New("accounts_decks: account has no decks: " + username)
	}

	if deckid == 0 {
		for i, _ := range accountsdecks {
			deletedecks = append(deletedecks, i)
		}
	} else if deck := accountsdecks[deckid]; deck != nil {
		deletedecks = []int{deck.Id}
	} else {
		return errors.New("accounts_decks: delete 404: " + username)
	}

	for _, deckid := range deletedecks {
		_, err := connection.Exec("DELETE FROM accounts_decks WHERE username=? AND id=?",
			username,
			deckid,
		)

		if err != nil {
			return err
		}

		_, err = connection.Exec("DELETE FROM accounts_decks_items WHERE username=? AND id=?",
			username,
			deckid,
		)

		if err != nil {
			return err
		}

	}

	log.Add("Username", username).Debug("accounts_decks: delete")
	return nil
}
