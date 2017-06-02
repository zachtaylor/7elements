package persistence

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"errors"
	"time"
)

func AccountsPacksGet(username string) ([]*SE.AccountPack, error) {
	rows, err := connection.Query("SELECT username, art, register FROM accounts_packs WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	accountspacks := make([]*SE.AccountPack, 0)

	for rows.Next() {
		accountpack := &SE.AccountPack{}
		var registerbuff int64

		err = rows.Scan(&accountpack.Username, &accountpack.ArtId, &registerbuff)
		if err != nil {
			return nil, err
		}

		accountpack.Register = time.Unix(registerbuff, 0)
		accountspacks = append(accountspacks, accountpack)
	}
	rows.Close()

	log.Add("Username", username).Debug("accounts_packs: get")
	return accountspacks, nil
}

func AccountsPacksInsert(username string) error {
	accountspacks := SE.AccountsPacks.Cache[username]
	if accountspacks == nil {
		return errors.New("accounts_packs: insert 404: " + username)
	}

	for _, accountpack := range accountspacks {
		_, err := connection.Exec("INSERT INTO accounts_packs(username, art, register) VALUES(?, ?, ?)",
			username,
			accountpack.ArtId,
			accountpack.Register.Unix(),
		)

		if err != nil {
			return err
		}
	}

	log.Add("Username", username).Debug("accounts_packs: insert")
	return nil
}

func AccountsPacksDelete(username string) error {
	_, err := connection.Exec("DELETE FROM accounts_packs WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	log.Add("Username", username).Debug("accounts_packs: delete")
	return nil
}
