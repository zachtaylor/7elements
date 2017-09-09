package sessions

import (
	"golang.org/x/net/websocket"
	"ztaylor.me/events"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

type Socket struct {
	SessionId uint
	Username  string
	Done      chan error
	conn      *websocket.Conn
}

func OpenSocket(session *Session, conn *websocket.Conn) {
	socket := &Socket{
		SessionId: session.Id,
		Username:  session.Username,
		Done:      make(chan error),
		conn:      conn,
	}
	startWatch(socket, session)
}

func (socket *Socket) Send(name string, data json.Json) {
	if socket.conn == nil {
		log.Add("Username", socket.Username).Add("Name", name).Warn("socket-send: conn is nil")
		return
	}
	websocket.Message.Send(socket.conn, json.Json{
		"name": name,
		"data": data,
	}.String())
}

func startWatch(socket *Socket, session *Session) {
	events.Fire("WebsocketOpen", socket)
	for socket.conn != nil {
		msgIn, msgErr := socket.receivers()
		select {
		case <-session.Done:
			socket.conn = nil
			close(socket.Done)
		case err := <-msgErr:
			socket.conn = nil
			close(socket.Done)
			if err.Error() != "EOF" {
				log.Add("Error", err).Add("Username", session.Username).Error("socket: receive")
			}
		case msg := <-msgIn:
			if msg == nil {
				socket.conn = nil
				close(socket.Done)
			} else {
				events.Fire("WebsocketMessage", socket, msg.Name, msg.Data)
			}
		}
	}
	events.Fire("WebsocketClose", socket)
}

func (socket *Socket) receivers() (chan *WebsocketMessage, chan error) {
	receiver := make(chan *WebsocketMessage)
	errors := make(chan error)
	go func() {
		defer close(receiver)
		defer close(errors)
		msg := NewWebsocketMessage()
		err := websocket.JSON.Receive(socket.conn, msg)
		if err != nil {
			errors <- err
		} else {
			receiver <- msg
		}
	}()
	return receiver, errors
}
