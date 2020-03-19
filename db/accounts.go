package db

import (
	"time"

	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/db"
)

type AccountService struct {
	conn  *db.DB
	cache map[string]*account.T
}

func NewAccountService(db *db.DB) account.Service {
	return &AccountService{
		conn:  db,
		cache: make(map[string]*account.T),
	}
}

func (as *AccountService) Test(username string) *account.T {
	return as.cache[username]
}

func (as *AccountService) Cache(a *account.T) {
	as.cache[a.Username] = a
}

func (as *AccountService) Forget(username string) {
	delete(as.cache, username)
}

func (as *AccountService) Find(username string) (*account.T, error) {
	if account := as.Test(username); account != nil {
		return account, nil
	}

	if account, err := as.Get(username); account == nil {
		return nil, err
	} else {
		as.cache[username] = account
		return account, nil
	}
}

func (as *AccountService) Get(username string) (*account.T, error) {
	row := as.conn.QueryRow(
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

func (as *AccountService) GetCount() (int, error) {
	row := as.conn.QueryRow("SELECT COUNT(*) FROM accounts")

	var ibuf int

	if err := row.Scan(&ibuf); err != nil {
		return -1, err
	}
	return ibuf, nil
}

func (as *AccountService) Insert(account *account.T) error {
	_, err := as.conn.Exec(
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

func (as AccountService) UpdateCoins(account *account.T) error {
	_, err := as.conn.Exec("UPDATE accounts SET coins=? WHERE username=?", account.Coins, account.Username)
	return err
}

func (as AccountService) UpdateEmail(account *account.T) error {
	_, err := as.conn.Exec(
		"UPDATE accounts SET email=? WHERE username=?",
		account.Email,
		account.Username,
	)
	return err
}

func (as AccountService) UpdateLogin(account *account.T) error {
	_, err := as.conn.Exec(
		"UPDATE accounts SET lastlogin=? WHERE username=?",
		account.LastLogin.Unix(),
		account.Username,
	)
	return err
}

func (as AccountService) UpdatePassword(account *account.T) error {
	_, err := as.conn.Exec(
		"UPDATE accounts SET password=? WHERE username=?",
		account.Password,
		account.Username,
	)
	return err
}

func (as AccountService) Delete(username string) error {
	_, err := as.conn.Exec("DELETE FROM accounts WHERE username=?", username)
	return err
}
