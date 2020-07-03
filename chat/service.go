package chat

type Service interface {
	// Sockets() *websocket.Cache
	Get(key string) *Room
	Remove(key string)
	New(key string, messageBuffer int) *Room
}
