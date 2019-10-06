package chat

type Service interface {
	Get(key string) *Room
	Remove(key string)
	New(key string, messageBuffer int) *Room
}
