package accounts

import (
	"strings"
	"time"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/db/accounts_decks"
	"ztaylor.me/cast"
	"ztaylor.me/db"
)

func Get(conn *db.DB, username string) (*account.T, error) {
	row := conn.QueryRow(
		"SELECT username, email, password, skill, coins, register, lastlogin FROM accounts WHERE username=?",
		username,
	)

	a := &account.T{}
	var registerbuff, lastloginbuff int64

	if err := row.Scan(
		&a.Username, &a.Email, &a.Password, &a.Skill, &a.Coins, &registerbuff, &lastloginbuff,
	); err == nil {
		a.Register = time.Unix(registerbuff, 0)
		a.LastLogin = time.Unix(lastloginbuff, 0)
	} else {
		return nil, err
	}

	if cards, err := getCards(conn, username); err == nil {
		a.Cards = cards
	} else {
		return nil, err
	}

	if decks, err := accounts_decks.Get(conn, username); err == nil {
		a.Decks = decks
	} else {
		return nil, err
	}

	return a, nil
}

func getCards(conn *db.DB, username string) (card.Count, error) {
	rows, err := conn.Query("SELECT card FROM accounts_cards WHERE username=?",
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

func Count(conn *db.DB) (int, error) {
	row := conn.QueryRow("SELECT COUNT(*) FROM accounts")

	var ibuf int

	if err := row.Scan(&ibuf); err != nil {
		return -1, err
	}
	return ibuf, nil
}

func Insert(conn *db.DB, account *account.T) error {
	_, err := conn.Exec(
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

func InsertCards(conn *db.DB, account *account.T, cardids []int) error {
	statement := strings.Builder{}
	statement.WriteString("INSERT INTO accounts_cards(username, card) VALUES ")
	for i, cardid := range cardids {
		if i > 0 {
			statement.WriteString(", ")
		}
		statement.WriteString("('")
		statement.WriteString(account.Username)
		statement.WriteString("',")
		statement.WriteString(cast.StringI(cardid))
		statement.WriteByte(')')
	}
	if _, err := conn.Exec(statement.String()); err != nil {
		return err
	}
	return nil
}

func UpdateCoins(conn *db.DB, account *account.T) error {
	_, err := conn.Exec("UPDATE accounts SET coins=? WHERE username=?", account.Coins, account.Username)
	return err
}

func UpdateEmail(conn *db.DB, account *account.T) error {
	_, err := conn.Exec(
		"UPDATE accounts SET email=? WHERE username=?",
		account.Email,
		account.Username,
	)
	return err
}

func UpdateLogin(conn *db.DB, account *account.T) error {
	_, err := conn.Exec(
		"UPDATE accounts SET lastlogin=? WHERE username=?",
		account.LastLogin.Unix(),
		account.Username,
	)
	return err
}

func UpdatePassword(conn *db.DB, account *account.T) error {
	_, err := conn.Exec(
		"UPDATE accounts SET password=? WHERE username=?",
		account.Password,
		account.Username,
	)
	return err
}

func Delete(conn *db.DB, username string) error {
	_, err := conn.Exec("DELETE FROM accounts WHERE username=?", username)
	return err
}

// func  InsertCard(username string, cardid int) error {
// 	_, err := conn.Exec("INSERT INTO accounts_cards(username, card) VALUES (?, ?)",
// 		username,
// 		cardid,
// 	)
// 	return err
// }

func DeleteCards(conn *db.DB, username string) error {
	_, err := conn.Exec("DELETE FROM accounts_cards WHERE username=?",
		username,
	)
	return err
}

// func  DeleteAndInsertCards(username string) error {
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
// func  TallyWinCount(username string) error {
// 	_, err := conn.Exec(
// 		"UPDATE accounts SET skill=skill+1 WHERE AND username=?",
// 		username,
// 	)
// 	return err
// }
