package chat

import "sync"

type Room struct {
	lock    sync.Mutex
	service Service
	key     string
	Users   map[string]*User
	history []*Message
}

// NewRoom creates a chat room
//
// messageBuffer must be > 0
func NewRoom(service Service, key string, messageBuffer int) *Room {
	room := &Room{
		service: service,
		key:     key,
		Users:   make(map[string]*User),
		history: make([]*Message, messageBuffer),
	}
	return room
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
	for _, user := range r.Users {
		if msg != nil {
			user.Send(msg)
		}
	}
}

func (r *Room) History() []*Message {
	return r.history
}

func (r *Room) AddUser(name string, send func(*Message)) {
	r.User(NewUser(name, send))
}

func (r *Room) User(user *User) {
	r.lock.Lock()
	r.Users[user.Name] = user
	r.lock.Unlock()
}

func (r *Room) Unsubscribe(username string) {
	r.lock.Lock()
	delete(r.Users, username)
	r.lock.Unlock()
	if len(r.Users) < 1 {
		r.Destroy()
	}
}

func (r *Room) Destroy() {
	r.service.Remove(r.key)
}
