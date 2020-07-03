package chat

import (
	"math/rand"

	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/keygen"
)

type MemService struct {
	// sockets *websocket.Cache
	keygen keygen.Keygener
	cache  map[string]*Room
	lock   cast.Mutex
}

// _memService type checks MemService
func _memService(m *MemService) Service {
	return m
}

// func NewMemService(sockets *websocket.Cache) *MemService {
func NewMemService() *MemService {
	return &MemService{
		// sockets: sockets,
		keygen: &keygen.Settings{
			KeySize: 24,
			CharSet: charset.AlphaCapitalNumeric,
			Rand:    rand.New(rand.NewSource(cast.Now().UnixNano())),
		},
		cache: make(map[string]*Room),
	}
}

// func (service *MemService) Sockets() *websocket.Cache {
// 	return service.sockets
// }

// Get a chat room
func (service *MemService) Get(key string) *Room {
	return service.cache[key]
}

// Remove a chat room
func (service *MemService) Remove(key string) {
	service.lock.Lock()
	delete(service.cache, key)
	service.lock.Unlock()
}

// New creates a Chat, or nil if the key is already in use
func (service *MemService) New(name string, messageBuffer int) *Room {
	service.lock.Lock()
	id := ""
	for id == "" || service.cache[id] != nil {
		id = service.keygen.Keygen()
	}
	room := NewRoom(service, id, name, 12)
	service.cache[id] = room
	service.lock.Unlock()
	return room
}
