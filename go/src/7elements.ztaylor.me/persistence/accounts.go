package persistence

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"errors"
	"time"
)

func AccountsGet(username string) (*SE.Account, error) {
	row := connection.QueryRow(
		"SELECT username, email, password, coins, language, register, lastlogin FROM accounts WHERE username=?",
		username,
	)

	account := &SE.Account{}
	var registerbuff, lastloginbuff int64

	if err := row.Scan(&account.Username, &account.Email, &account.Password, &account.Coins, &account.Language, &registerbuff, &lastloginbuff); err != nil {
		return nil, err
	} else {
		account.Register = time.Unix(registerbuff, 0)
		account.LastLogin = time.Unix(lastloginbuff, 0)
	}

	log.Add("Username", username).Debug("accounts: get")
	return account, nil
}

func AccountsInsert(username string) error {
	account := SE.Accounts.Cache[username]
	if account == nil {
		return errors.New("accounts: insert 404: " + username)
	}

	_, err := connection.Exec(
		"INSERT INTO accounts (username, email, password, coins, language, register, lastlogin) VALUES (?, ?, ?, ?, ?, ?, ?)",
		account.Username,
		account.Email,
		account.Password,
		account.Coins,
		account.Language,
		account.Register.Unix(),
		account.LastLogin.Unix(),
	)

	if err != nil {
		return err
	}

	log.Add("Username", username).Debug("accounts: insert")
	return nil
}

func AccountsDelete(username string) error {
	_, err := connection.Exec("DELETE FROM accounts WHERE username=?", username)

	if err == nil {
		log.Add("Username", username).Debug("accounts: delete")
	}

	return err
}
