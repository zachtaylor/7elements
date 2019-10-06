package db

import (
	"errors"
	"time"

	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/db"
)

type AccountCardService struct {
	conn  *db.DB
	cache map[string]vii.AccountCards
}

func NewAccountCardService(db *db.DB) vii.AccountCardService {
	return &AccountCardService{
		conn:  db,
		cache: make(map[string]vii.AccountCards),
	}
}

func (acs *AccountCardService) Test(username string) vii.AccountCards {
	return acs.cache[username]
}

func (acs *AccountCardService) Forget(username string) {
	delete(acs.cache, username)
}

func (acs *AccountCardService) Find(username string) (vii.AccountCards, error) {
	if acs.cache[username] == nil {
		if stack, err := acs.Get(username); err != nil {
			return nil, err
		} else {
			acs.cache[username] = stack
		}
	}
	return acs.cache[username], nil
}

func (acs *AccountCardService) Get(username string) (vii.AccountCards, error) {
	rows, err := acs.conn.Query("SELECT username, card, register, notes FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	collection := vii.AccountCards{}

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

func (acs *AccountCardService) Insert(username string) error {
	stack := acs.Test(username)
	if stack == nil {
		return errors.New("accountscards missing")
	}

	for _, list := range stack {
		for _, accountcard := range list {
			if err := acs.InsertCard(accountcard); err != nil {
				return err
			}
		}
	}

	return nil
}

func (acs *AccountCardService) InsertCard(card *vii.AccountCard) error {
	if acs.cache[card.Username] == nil {
		acs.cache[card.Username] = vii.AccountCards{}
	}
	cards := acs.cache[card.Username]
	if list := cards[card.CardId]; list == nil {
		cards[card.CardId] = make([]*vii.AccountCard, 0)
	}
	cards[card.CardId] = append(cards[card.CardId], card)
	acs.cache[card.Username] = cards
	_, err := acs.conn.Exec("INSERT INTO accounts_cards(username, card, register, notes) VALUES (?, ?, ?, ?)",
		card.Username,
		card.CardId,
		card.Register.Unix(),
		card.Notes,
	)
	return err
}

func (acs *AccountCardService) Delete(username string) error {
	_, err := acs.conn.Exec("DELETE FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	return nil
}

func (acs *AccountCardService) DeleteAndInsert(username string) error {
	if cardcollection := acs.Test(username); cardcollection != nil {
		if err := acs.Delete(username); err != nil {
			return err
		} else if err := acs.Insert(username); err != nil {
			return err
		}
	}
	return nil
}
