package db

import (
	"time"

	"github.com/zachtaylor/7tcg"
)

type AccountService map[string]*vii.Account

func init() {
	vii.AccountService = make(AccountService)
}

func (cache AccountService) Test(username string) *vii.Account {
	return cache[username]
}

func (cache AccountService) Cache(a *vii.Account) {
	cache[a.Username] = a
}

func (cache AccountService) Forget(username string) {
	delete(cache, username)
}

func (cache AccountService) Get(username string) (*vii.Account, error) {
	if account := cache.Test(username); account != nil {
		return account, nil
	}

	if account, err := cache.Load(username); account == nil {
		return nil, err
	} else {
		cache[username] = account
		return account, nil
	}
}

func (cache AccountService) Load(username string) (*vii.Account, error) {
	row := conn.QueryRow(
		"SELECT username, email, password, skill, coins, packs, language, register, lastlogin FROM accounts WHERE username=?",
		username,
	)

	account := vii.NewAccount()
	var registerbuff, lastloginbuff int64

	if err := row.Scan(&account.Username, &account.Email, &account.Password, &account.Skill, &account.Coins, &account.Packs, &account.Language, &registerbuff, &lastloginbuff); err != nil {
		return nil, err
	} else {
		account.Register = time.Unix(registerbuff, 0)
		account.LastLogin = time.Unix(lastloginbuff, 0)
	}

	return account, nil
}

func (cache AccountService) Insert(account *vii.Account) error {
	_, err := conn.Exec(
		"INSERT INTO accounts (username, email, password, skill, coins, packs, language, register, lastlogin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		account.Username,
		account.Email,
		account.Password,
		account.Skill,
		account.Coins,
		account.Packs,
		account.Language,
		account.Register.Unix(),
		account.LastLogin.Unix(),
	)

	return err
}

func (_ AccountService) UpdatePacks(account *vii.Account) error {
	_, err := conn.Exec("UPDATE accounts SET packs=? WHERE username=?", account.Packs, account.Username)
	return err
}

func (_ AccountService) UpdateCoins(account *vii.Account) error {
	_, err := conn.Exec("UPDATE accounts SET coins=? WHERE username=?", account.Coins, account.Username)
	return err
}

func (_ AccountService) UpdateLogin(account *vii.Account) error {
	_, err := conn.Exec(
		"UPDATE accounts SET lastlogin=? WHERE username=?",
		account.LastLogin.Unix(),
		account.Username,
	)
	return err
}

func (_ AccountService) UpdatePassword(account *vii.Account) error {
	_, err := conn.Exec(
		"UPDATE accounts SET password=? WHERE username=?",
		account.Password,
		account.Username,
	)
	return err
}

func (_ AccountService) Delete(username string) error {
	_, err := conn.Exec("DELETE FROM accounts WHERE username=?", username)
	return err
}
