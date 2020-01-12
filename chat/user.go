package chat

type User struct {
	Name string
	Send func(path string, m *Message)
}

func NewUser(name string, send func(path string, m *Message)) *User {
	return &User{
		Name: name,
		Send: send,
	}
}
