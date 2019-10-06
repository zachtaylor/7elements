package chat

type MemService map[string]*Room

// _memService type checks MemService
func _memService(m MemService) Service {
	return m
}

// Get a chat room
func (service MemService) Get(key string) *Room {
	return service[key]
}

// Remove a chat room
func (service MemService) Remove(key string) {
	delete(service, key)
}

// New creates a Chat, or nil if the key is already in use
func (service MemService) New(key string, messageBuffer int) *Room {
	if service[key] != nil {
		return nil
	}
	room := NewRoom(service, key, 12)
	service[key] = room
	return room
}
