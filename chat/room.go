package chat

import (
	"sync"

	"ztaylor.me/http/websocket"
)

type Room struct {
	lock     sync.Mutex
	service  Service
	key      string
	users    map[string]*User
	userkeys map[string]string
	history  []*Message
}

// NewRoom creates a chat room
//
// messageBuffer must be > 0
func NewRoom(service Service, key string, messageBuffer int) *Room {
	room := &Room{
		service:  service,
		key:      key,
		users:    make(map[string]*User),
		userkeys: make(map[string]string),
		history:  make([]*Message, messageBuffer),
	}
	return room
}

func (r *Room) Name() string {
	return r.key
}

func (r *Room) AddMessage(msg *Message) {
	r.lock.Lock()
	r.addMessage(msg)
	r.lock.Unlock()
}
func (r *Room) addMessage(msg *Message) {
	for i := len(r.history) - 1; i > 0; i-- {
		r.history[i] = r.history[i-1]
	}
	r.history[0] = msg
	for _, user := range r.users {
		if msg != nil {
			user.Send("/chat/game", msg)
		}
	}
}

func (r *Room) History() []*Message {
	return r.history
}

func (r *Room) AddUser(socket *websocket.T) *User {
	username := ""
	r.lock.Lock()
	for i := 1; username == "" || r.users[username] != nil; i++ {
		username = "anon#" + socket.ID[:i]
	}
	user := NewUser(username, func(path string, msg *Message) {
		socket.Message(path, msg.JSON())
	})
	r.users[username] = user
	r.userkeys[socket.ID] = username
	r.lock.Unlock()
	return user
}

func (r *Room) User(key string) (user *User) {
	return r.users[r.userkeys[key]]
}

func (r *Room) Unsubscribe(username string) {
	r.lock.Lock()
	delete(r.users, username)
	keys := make([]string, 0)
	for key, _username := range r.userkeys {
		if username == _username { // this should only proc once
			keys = append(keys, key) // I've written this as a loop
		} // for the sake of completeness in the data type
	}
	for _, key := range keys { // probably len(keys) == 1
		delete(r.userkeys, key)
	}
	r.lock.Unlock()
	if len(r.users) < 1 {
		r.Destroy()
	}
}

func (r *Room) Destroy() {
	r.service.Remove(r.key)
}
