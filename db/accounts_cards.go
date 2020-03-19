package db

import (
	"errors"
	"time"

	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/db"
)

type AccountCardService struct {
	conn  *db.DB
	cache map[string]account.Cards
}

func NewAccountCardService(db *db.DB) account.CardService {
	return &AccountCardService{
		conn:  db,
		cache: make(map[string]account.Cards),
	}
}

func (acs *AccountCardService) Test(username string) account.Cards {
	return acs.cache[username]
}

func (acs *AccountCardService) Forget(username string) {
	delete(acs.cache, username)
}

func (acs *AccountCardService) Find(username string) (account.Cards, error) {
	if acs.cache[username] == nil {
		if stack, err := acs.Get(username); err != nil {
			return nil, err
		} else {
			acs.cache[username] = stack
		}
	}
	return acs.cache[username], nil
}

func (acs *AccountCardService) Get(username string) (account.Cards, error) {
	rows, err := acs.conn.Query("SELECT username, card, register, notes FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	collection := account.Cards{}

	for rows.Next() {
		c := &account.Card{}
		var registerbuff int64

		err = rows.Scan(&c.Username, &c.ProtoID, &registerbuff, &c.Notes)
		if err != nil {
			return nil, err
		}

		c.Register = time.Unix(registerbuff, 0)

		if list := collection[c.ProtoID]; list != nil {
			collection[c.ProtoID] = append(list, c)
		} else {
			collection[c.ProtoID] = []*account.Card{c}
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

func (acs *AccountCardService) InsertCard(c *account.Card) error {
	if acs.cache[c.Username] == nil {
		acs.cache[c.Username] = account.Cards{}
	}
	cards := acs.cache[c.Username]
	if list := cards[c.ProtoID]; list == nil {
		cards[c.ProtoID] = make([]*account.Card, 0)
	}
	cards[c.ProtoID] = append(cards[c.ProtoID], c)
	acs.cache[c.Username] = cards
	_, err := acs.conn.Exec("INSERT INTO accounts_cards(username, card, register, notes) VALUES (?, ?, ?, ?)",
		c.Username,
		c.ProtoID,
		c.Register.Unix(),
		c.Notes,
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
