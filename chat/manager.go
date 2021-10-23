package chat

import "taylz.io/http/user"

type Manager struct {
	Settings
	cache *Cache
}

func NewManager(settings Settings) (manager *Manager) {
	manager = &Manager{Settings: settings, cache: NewCache()}
	settings.Users.Observe(manager.onUser)
	return
}

func (m *Manager) onUser(id string, oldUser, newUser *user.T) {
	if oldUser != nil && newUser == nil {
		go m.RemoveUser(id)
	}
}

func (m *Manager) RemoveUser(name string) {
	for _, id := range m.cache.Keys() {
		if room := m.cache.Get(id); room.users[name] {
			go room.RemoveUser(name)
		}
	}
}

func (m *Manager) TryMessage(roomNo, username, message string) (err error) {
	if room := m.cache.Get(roomNo); room == nil {
		err = ErrNoRoom
	} else if !room.users[username] {
		err = ErrUserNotInRoom
	} else {
		room.AddSync(NewMessage(username, message))
	}
	return
}

// New creates a Chat, or nil if the key is already in use
func (m *Manager) New(name string, messageBuffer int) (room *Room) {
	m.cache.Sync(func(get CacheGetter, set CacheSetter) {
		id := ""
		for ok := true; ok; ok = (get(id) != nil) {
			id = m.Keygen()
		}
		room = NewRoom(m, id, name, 16)
		set(id, room)
	})
	return
}

// Get returns a chat room
func (m *Manager) Get(key string) *Room { return m.cache.Get(key) }

// Remove a chat room
func (m *Manager) Remove(key string) { m.cache.Remove(key) }
