package player

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
)

func (cache *Cache) Signup(username, email, password string) (*T, error) {
	a := &account.T{
		Username: username,
		Email:    email,
		Password: password,
		Skill:    1000,
		Register: cast.Now(),
	}

	player, err := cache.Login(a)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 3; i++ {
		proto := deck.NewPrototype()
		proto.User = username
		a.Decks[proto.ID] = proto
	}

	if err := cache.Settings.Accounts.Insert(a); err != nil {
		return nil, err
	}

	// if err := emails.SendValidationEmail(account); err != nil {
	// 	log.Clone().Add("mail-user", options.String("mail-user")).Add("mail-pass", options.String("mail-pass")).Add("mail-host", options.String("mail-host")).Add("Error", err).Error("/api/signup: send validation email")
	// }

	return player, nil
}
