package dbservice

import (
	"elemen7s.com"
	"elemen7s.com/db"
	"errors"
	"time"
)

func init() {
	vii.AccountCardService = make(AccountCardService)
}

type AccountCardService map[string]vii.AccountsCards

func (service AccountCardService) Test(username string) vii.AccountsCards {
	return service[username]
}

func (service AccountCardService) Forget(username string) {
	delete(service, username)
}

func (service AccountCardService) Get(username string) (vii.AccountsCards, error) {
	if service[username] == nil {
		if stack, err := service.Load(username); err != nil {
			return nil, err
		} else {
			service[username] = stack
		}
	}
	return service[username], nil
}

func (service AccountCardService) Load(username string) (vii.AccountsCards, error) {
	rows, err := db.Connection.Query("SELECT username, card, register, notes FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	collection := vii.AccountsCards{}

	for rows.Next() {
		accountcard := &vii.AccountCard{}
		var registerbuff int64

		err = rows.Scan(&accountcard.Username, &accountcard.CardId, &registerbuff, &accountcard.Notes)
		if err != nil {
			return nil, err
		}

		accountcard.Register = time.Unix(registerbuff, 0)

		if list := collection[accountcard.CardId]; list != nil {
			collection[accountcard.CardId] = append(list, accountcard)
		} else {
			collection[accountcard.CardId] = []*vii.AccountCard{accountcard}
		}
	}
	rows.Close()

	return collection, nil
}

func (service AccountCardService) Insert(username string) error {
	stack := service.Test(username)
	if stack == nil {
		return errors.New("accountscards missing")
	}

	for _, list := range stack {
		for _, accountcard := range list {
			if err := service.InsertCard(accountcard); err != nil {
				return err
			}
		}
	}

	return nil
}

func (service AccountCardService) InsertCard(card *vii.AccountCard) error {
	_, err := db.Connection.Exec("INSERT INTO accounts_cards(username, card, register, notes) VALUES (?, ?, ?, ?)",
		card.Username,
		card.CardId,
		card.Register.Unix(),
		card.Notes,
	)
	return err
}

func (service AccountCardService) Delete(username string) error {
	_, err := db.Connection.Exec("DELETE FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	return nil
}

func (service AccountCardService) DeleteAndInsert(username string) error {
	if cardcollection := service.Test(username); cardcollection != nil {
		if err := service.Delete(username); err != nil {
			return err
		} else if err := service.Insert(username); err != nil {
			return err
		}
	}
	return nil
}
