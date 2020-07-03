package chat

import (
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

type Room struct {
	Service Service
	id      string
	name    string
	users   map[string]*User
	history []*Message
	lock    cast.Mutex
}

// NewRoom creates a chat room
//
// messageBuffer must be > 0
func NewRoom(service Service, id, name string, messageBuffer int) *Room {
	room := &Room{
		Service: service,
		id:      id,
		name:    name,
		users:   make(map[string]*User),
		history: make([]*Message, messageBuffer),
	}
	return room
}

func (r *Room) Name() string {
	return r.name
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
			user.Socket.Send("/chat/game", msg.JSON())
		}
	}
}

func (r *Room) History() []*Message {
	return r.history
}

func (r *Room) AddUser(socket *websocket.T) *User {
	username := ""
	if socket.Session != nil {
		username = socket.Session.Name()
	} else {
		username = "anon"
	}
	r.lock.Lock()
	user := NewUser(username, socket)
	r.users[socket.ID()] = user
	r.lock.Unlock()
	return user
}

func (r *Room) User(key string) (user *User) {
	return r.users[key]
}

func (r *Room) Unsubscribe(key string) {
	r.lock.Lock()
	delete(r.users, key)
	r.lock.Unlock()
	if len(r.users) < 1 {
		r.Destroy()
	}
}

func (r *Room) Destroy() {
	r.Service.Remove(r.id)
}
