package ai

import (
	"bytes"
	"encoding/json"
	"time"

	"taylz.io/http/websocket"
)

// Input is where the ai receives input
type Input struct {
	AI *AI
}

func (i *Input) Name() string { return i.AI.Name }

func (i *Input) Done() <-chan struct{} { return i.AI.done }

func (i *Input) WriteMessage(msg *websocket.Message) error {
	i.Receive(msg.URI, msg.Data)
	return nil
}
func (i *Input) WriteMessageData(data []byte) error {
	msg := &websocket.Message{}
	if err := json.NewDecoder(bytes.NewBufferString(string(data))).Decode(msg); err != nil {
		i.AI.Game.Log().Add("Error", err).Warn("failed to parse message")
	} else {
		i.AI.Game.Log().Trace("received ", msg.URI)
	}
	i.Receive(msg.URI, msg.Data)
	return nil
}

// Receive data from the game runtime
func (i *Input) Receive(uri string, data map[string]any) {
	go time.AfterFunc(i.AI.Delay, func() {
		i.receive(uri, data)
	})
}
func (i *Input) receive(uri string, data map[string]any) {
	if uri == "/game/state" {
		i.AI.GameState(data)
	} else if uri == "/game/choice" {
		i.AI.GameChoice(data)
	} else if uri == "/game" {
	} else if uri == "/game/react" {
	} else if uri == "/game/card" {
	} else if uri == "/game/hand" {
	} else if uri == "/game/seat" {
	} else if uri == "/alert" {
	} else {
		i.AI.Game.Log().With(map[string]any{
			"URI":      uri,
			"GameId":   i.AI.Game.ID(),
			"Username": i.AI.Seat.Username,
		}).Warn("uri unknown")
	}
}
