package accounts

import (
	"elemen7s.com/db"
	"errors"
	"time"
	"ztaylor.me/events"
)

var cache = make(map[string]*Account)

func Store(a *Account) error {
	if a2 := Test(a.Username); a2 != nil && a != a2 {
		return errors.New("account already stored")
	}
	cache[a.Username] = a
	return nil
}

func Forget(username string) {
	delete(cache, username)
}

func Test(username string) *Account {
	return cache[username]
}

func Get(username string) (*Account, error) {
	if account := Test(username); account != nil {
		return account, nil
	}

	if account, err := Load(username); account == nil {
		return nil, err
	} else {
		cache[username] = account
		return account, nil
	}
}

func Load(username string) (*Account, error) {
	row := db.Connection.QueryRow(
		"SELECT username, email, password, skill, coins, packs, language, register, lastlogin FROM accounts WHERE username=?",
		username,
	)

	account := &Account{}
	var registerbuff, lastloginbuff int64

	if err := row.Scan(&account.Username, &account.Email, &account.Password, &account.Skill, &account.Coins, &account.Packs, &account.Language, &registerbuff, &lastloginbuff); err != nil {
		return nil, err
	} else {
		account.Register = time.Unix(registerbuff, 0)
		account.LastLogin = time.Unix(lastloginbuff, 0)
	}

	events.Fire("accounts.Load", account)
	return account, nil
}

func Insert(username string) error {
	account := Test(username)
	if account == nil {
		return errors.New("account missing")
	}

	_, err := db.Connection.Exec(
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

func UpdatePackCount(username string, packCount int) error {
	if account := Test(username); account == nil {
		return errors.New("accounts.UpdatePackCount: account missing")
	} else {
		account.Packs = packCount
		_, err := db.Connection.Exec("UPDATE accounts SET packs=? WHERE username=?", packCount, username)
		return err
	}
}

func UpdateAddCoins(username string, count int) error {
	if account := Test(username); account == nil {
		return errors.New("accounts.UpdateAddCoins: account missing")
	} else {
		account.Coins += count
		_, err := db.Connection.Exec("UPDATE accounts SET coins=? WHERE username=?", account.Coins, username)
		return err
	}
}

func Delete(username string) error {
	_, err := db.Connection.Exec("DELETE FROM accounts WHERE username=?", username)
	return err
}
