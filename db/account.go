package db

import (
	"time"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/db"
)

type AccountService struct {
	conn *db.DB
}

func NewAccountService(db *db.DB) *AccountService {
	return &AccountService{
		conn: db,
	}
}

func (s *AccountService) Get(username string) (*account.T, error) {
	row := s.conn.QueryRow(
		"SELECT username, email, password, skill, coins, register, lastlogin FROM accounts WHERE username=?",
		username,
	)

	a := &account.T{}
	var registerbuff, lastloginbuff int64

	if err := row.Scan(&a.Username, &a.Email, &a.Password, &a.Skill, &a.Coins, &registerbuff, &lastloginbuff); err != nil {
		return nil, err
	} else {
		a.Register = time.Unix(registerbuff, 0)
		a.LastLogin = time.Unix(lastloginbuff, 0)
	}

	return a, nil
}

func (s *AccountService) GetCount() (int, error) {
	row := s.conn.QueryRow("SELECT COUNT(*) FROM accounts")

	var ibuf int

	if err := row.Scan(&ibuf); err != nil {
		return -1, err
	}
	return ibuf, nil
}

func (s *AccountService) Insert(account *account.T) error {
	_, err := s.conn.Exec(
		"INSERT INTO accounts (username, email, password, skill, coins, register, lastlogin) VALUES (?, ?, ?, ?, ?, ?, ?)",
		account.Username,
		account.Email,
		account.Password,
		account.Skill,
		account.Coins,
		account.Register.Unix(),
		account.LastLogin.Unix(),
	)

	return err
}

func (s *AccountService) UpdateCoins(account *account.T) error {
	_, err := s.conn.Exec("UPDATE accounts SET coins=? WHERE username=?", account.Coins, account.Username)
	return err
}

func (s *AccountService) UpdateEmail(account *account.T) error {
	_, err := s.conn.Exec(
		"UPDATE accounts SET email=? WHERE username=?",
		account.Email,
		account.Username,
	)
	return err
}

func (s *AccountService) UpdateLogin(account *account.T) error {
	_, err := s.conn.Exec(
		"UPDATE accounts SET lastlogin=? WHERE username=?",
		account.LastLogin.Unix(),
		account.Username,
	)
	return err
}

func (s *AccountService) UpdatePassword(account *account.T) error {
	_, err := s.conn.Exec(
		"UPDATE accounts SET password=? WHERE username=?",
		account.Password,
		account.Username,
	)
	return err
}

func (s *AccountService) Delete(username string) error {
	_, err := s.conn.Exec("DELETE FROM accounts WHERE username=?", username)
	return err
}

func (s *AccountService) GetCards(username string) (card.Count, error) {
	rows, err := s.conn.Query("SELECT card FROM accounts_cards WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	cards := card.Count{}

	for rows.Next() {
		var buf int
		if err = rows.Scan(&buf); err != nil {
			return nil, err
		}
		cards[buf]++
	}
	rows.Close()

	return cards, nil
}

func (s *AccountService) InsertCard(username string, cardid int) error {
	_, err := s.conn.Exec("INSERT INTO accounts_cards(username, card) VALUES (?, ?)",
		username,
		cardid,
	)
	return err
}

func (s *AccountService) DeleteCards(username string) error {
	_, err := s.conn.Exec("DELETE FROM accounts_cards WHERE username=?",
		username,
	)
	return err
}

func (s *AccountService) GetDecks(user string) (deck.Prototypes, error) {
	return getUserDecks(s.conn, user)
}

// func (s *AccountService) DeleteAndInsertCards(username string) error {
// 	if cardcollection := s.TestCards(username); cardcollection != nil {
// 		if err := s.DeleteCards(username); err != nil {
// 			return err
// 		} else if err := s.InsertCards(username); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// this is a product of refactoring accounts_decks
// func (s *AccountService) TallyWinCount(username string) error {
// 	_, err := s.conn.Exec(
// 		"UPDATE accounts SET skill=skill+1 WHERE AND username=?",
// 		username,
// 	)
// 	return err
// }
