package accountscards

import (
	"elemen7s.com/db"
	"errors"
	"time"
)

var cache = make(map[string]Stack)

func Test(username string) Stack {
	return cache[username]
}

func Forget(username string) {
	delete(cache, username)
}

func Get(username string) (Stack, error) {
	if cache[username] == nil {
		if stack, err := Load(username); err != nil {
			return nil, err
		} else {
			cache[username] = stack
		}
	}
	return cache[username], nil
}

func Load(username string) (Stack, error) {
	rows, err := db.Connection.Query("SELECT username, card, register, notes FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	collection := Stack{}

	for rows.Next() {
		accountcard := &AccountCard{}
		var registerbuff int64

		err = rows.Scan(&accountcard.Username, &accountcard.CardId, &registerbuff, &accountcard.Notes)
		if err != nil {
			return nil, err
		}

		accountcard.Register = time.Unix(registerbuff, 0)

		if list := collection[accountcard.CardId]; list != nil {
			collection[accountcard.CardId] = append(list, accountcard)
		} else {
			collection[accountcard.CardId] = []*AccountCard{accountcard}
		}
	}
	rows.Close()

	return collection, nil
}

func Insert(username string) error {
	stack := Test(username)
	if stack == nil {
		return errors.New("accountscards missing")
	}

	for _, list := range stack {
		for _, accountcard := range list {
			if err := InsertCard(accountcard); err != nil {
				return err
			}
		}
	}

	return nil
}

func InsertCard(card *AccountCard) error {
	_, err := db.Connection.Exec("INSERT INTO accounts_cards(username, card, register, notes) VALUES (?, ?, ?, ?)",
		card.Username,
		card.CardId,
		card.Register.Unix(),
		card.Notes,
	)
	return err
}

func Delete(username string) error {
	_, err := db.Connection.Exec("DELETE FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAndInsert(username string) error {
	if cardcollection := Test(username); cardcollection != nil {
		if err := Delete(username); err != nil {
			return err
		} else if err := Insert(username); err != nil {
			return err
		}
	}
	return nil
}
