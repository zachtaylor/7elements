package player

import (
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/cast"
)

type Cache struct {
	Settings CacheSettings
	cache    map[string]*T
	lock     cast.Mutex
}

// Login returns a player for the account
func (cache *Cache) Login(account *account.T) (player *T, error error) {
	log := cache.Settings.Logger.New()
	log.Trace()

	if player = cache.Get(account.Username); player != nil {
		log.Warn("player exists")
		return
	}

	session := cache.Settings.Sessions.Get(account.SessionID)
	if session != nil {
		log.Warn("session exists")
	} else if session = cache.Settings.Sessions.Find(account.Username); session != nil {
		log.Error("session found")
		account.SessionID = session.ID()
	} else {
		log.Info("session grant")
		session = cache.Settings.Sessions.Grant(account.Username)
		account.SessionID = session.ID()
	}

	player = &T{
		Settings: NewSettings(cache.Settings.Sockets),
		Session:  session,
		Account:  account,
		conns:    make(map[string]bool),
	}
	go cache.waitPlayer(player)
	cache.Set(player)
	return
}

// Size returns the number of connected sockets
func (cache *Cache) Size() int {
	return len(cache.cache)
}

// Get returns a player connected by the given key
func (cache *Cache) Get(name string) *T {
	return cache.cache[name]
}

// Set adds a player connected by the username
func (cache *Cache) Set(player *T) {
	cache.lock.Lock()
	cache.cache[player.Account.Username] = player
	cache.lock.Unlock()
}

func (cache *Cache) Delete(name string) {
	cache.lock.Lock()
	delete(cache.cache, name)
	cache.lock.Unlock()
}
