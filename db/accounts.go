package db

import (
	"time"

	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/db"
)

type AccountService struct {
	conn  *db.DB
	cache map[string]*vii.Account
}

func NewAccountService(db *db.DB) vii.AccountService {
	return &AccountService{
		conn:  db,
		cache: make(map[string]*vii.Account),
	}
}

func (as *AccountService) Test(username string) *vii.Account {
	return as.cache[username]
}

func (as *AccountService) Cache(a *vii.Account) {
	as.cache[a.Username] = a
}

func (as *AccountService) Forget(username string) {
	delete(as.cache, username)
}

func (as *AccountService) Find(username string) (*vii.Account, error) {
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

func (as *AccountService) Get(username string) (*vii.Account, error) {
	row := as.conn.QueryRow(
		"SELECT username, email, password, skill, coins, register, lastlogin FROM accounts WHERE username=?",
		username,
	)

	account := vii.NewAccount()
	var registerbuff, lastloginbuff int64

	if err := row.Scan(&account.Username, &account.Email, &account.Password, &account.Skill, &account.Coins, &registerbuff, &lastloginbuff); err != nil {
		return nil, err
	} else {
		account.Register = time.Unix(registerbuff, 0)
		account.LastLogin = time.Unix(lastloginbuff, 0)
	}

	return account, nil
}

func (as *AccountService) GetCount() (int, error) {
	row := as.conn.QueryRow("SELECT COUNT(*) FROM accounts")

	var ibuf int

	if err := row.Scan(&ibuf); err != nil {
		return -1, err
	}
	return ibuf, nil
}

func (as *AccountService) Insert(account *vii.Account) error {
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

func (as AccountService) UpdateCoins(account *vii.Account) error {
	_, err := as.conn.Exec("UPDATE accounts SET coins=? WHERE username=?", account.Coins, account.Username)
	return err
}

func (as AccountService) UpdateEmail(account *vii.Account) error {
	_, err := as.conn.Exec(
		"UPDATE accounts SET email=? WHERE username=?",
		account.Email,
		account.Username,
	)
	return err
}

func (as AccountService) UpdateLogin(account *vii.Account) error {
	_, err := as.conn.Exec(
		"UPDATE accounts SET lastlogin=? WHERE username=?",
		account.LastLogin.Unix(),
		account.Username,
	)
	return err
}

func (as AccountService) UpdatePassword(account *vii.Account) error {
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
