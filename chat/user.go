package chat

type User struct {
	Name string
	Send func(*Message)
}

func NewUser(name string, send func(*Message)) *User {
	return &User{
		Name: name,
		Send: send,
	}
}
