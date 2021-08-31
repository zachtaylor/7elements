package chat

import "sync"

type History struct {
	messages []*Message
	mu       sync.Mutex
}

func NewHistory(buf int) *History {
	return &History{
		messages: make([]*Message, buf),
	}
}

func (h *History) Add(m *Message) {
	h.mu.Lock()
	for i := len(h.messages) - 1; i > 0; i-- {
		h.messages[i] = h.messages[i-1]
	}
	h.messages[0] = m
	h.mu.Unlock()
}

func (h *History) Data() (buf []*Message) {
	buf = make([]*Message, len(h.messages))
	h.mu.Lock()
	copy(buf, h.messages)
	h.mu.Unlock()
	return buf
}
