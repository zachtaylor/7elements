package persistence

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"errors"
	"time"
)

func AccountsCardsGet(username string) (SE.CardCollection, error) {
	rows, err := connection.Query("SELECT username, card, register, notes FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	collection := SE.CardCollection{}

	for rows.Next() {
		accountcard := &SE.AccountCard{}
		var registerbuff int64

		err = rows.Scan(&accountcard.Username, &accountcard.Card, &registerbuff, &accountcard.Notes)
		if err != nil {
			return nil, err
		}

		accountcard.Register = time.Unix(registerbuff, 0)

		if list := collection[accountcard.Card]; list != nil {
			collection[accountcard.Card] = append(list, accountcard)
		} else {
			collection[accountcard.Card] = []*SE.AccountCard{accountcard}
		}
	}
	rows.Close()

	log.Add("Username", username).Debug("accounts_cards: get")
	return collection, nil
}

func AccountsCardsInsert(username string) error {
	collection := SE.AccountsCards.Cache[username]
	if collection == nil {
		return errors.New("accounts_cards: insert 404: " + username)
	}

	for cardId, list := range collection {
		for _, accountcard := range list {
			_, err := connection.Exec("INSERT INTO accounts_cards(username, card, register, notes) VALUES (?, ?, ?, ?)",
				username,
				cardId,
				accountcard.Register.Unix(),
				accountcard.Notes,
			)

			if err != nil {
				return err
			}
		}
	}

	log.Add("Username", username).Debug("accounts_cards: insert")
	return nil
}

func AccountsCardsDelete(username string) error {
	_, err := connection.Exec("DELETE FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	log.Add("Username", username).Debug("accounts_cards: delete")
	return nil
}
