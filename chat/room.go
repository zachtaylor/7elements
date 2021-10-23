package chat

import (
	"sync"

	"github.com/zachtaylor/7elements/wsout"
)

type Room struct {
	id    string
	name  string
	users map[string]bool
	mu    sync.Mutex
	hist  *History
	man   *Manager
}

// NewRoom creates a chat room
//
// messageBuffer must be > 0
func NewRoom(man *Manager, id, name string, messageBuffer int) *Room {
	room := &Room{
		id:   id,
		name: name,
		hist: NewHistory(messageBuffer),
		man:  man,
	}
	return room
}

func (r *Room) ID() string { return r.id }

func (r *Room) Name() string { return r.name }

func (r *Room) Add(username, message string) {
	go r.AddSync(NewMessage(username, message))
}

func (r *Room) AddSync(msg *Message) {
	r.mu.Lock()
	r.hist.Add(msg)
	i, keys := 0, make([]string, len(r.users))
	for k := range r.users {
		keys[i] = k
		i++
	}
	r.mu.Unlock()
	data := wsout.Chat(msg.Data()).EncodeToJSON()
	for _, username := range keys {
		if user, _, _ := r.man.Users.Get(username); user != nil {
			user.Write(data)
		}
	}
}

func (r *Room) AddUser(username string) {
	r.mu.Lock()
	r.users[username] = true
	r.mu.Unlock()
}

func (r *Room) RemoveUser(username string) {
	r.mu.Lock()
	delete(r.users, username)
	r.mu.Unlock()
	if len(r.users) < 1 {
		go r.Destroy()
	}
}

func (r *Room) History() (data [][]byte) {
	for _, msg := range r.hist.Data() {
		data = append(data, wsout.Chat(msg.Data()).EncodeToJSON())
	}
	return
}

func (r *Room) Destroy() {
	r.man.cache.Remove(r.id)
}
