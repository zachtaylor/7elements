package chat

type Receiver interface {
	SendChat(*Message)
}
