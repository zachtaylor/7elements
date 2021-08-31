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

func (i *Input) Message(uri string, data map[string]interface{}) {
	i.Receive(uri, data)
}

func (i *Input) Write(data []byte) {
	msg := &websocket.Message{}
	if err := json.NewDecoder(bytes.NewBufferString(string(data))).Decode(msg); err != nil {
		i.AI.Game.Log().Add("Error", err).Warn("failed to parse message")
	} else {
		i.AI.Game.Log().Trace("received ", msg.URI)
	}
	i.Receive(msg.URI, msg.Data)
}

// Receive data from the game runtime
func (i *Input) Receive(uri string, data map[string]interface{}) {
	go time.AfterFunc(i.AI.Delay, func() {
		i.receive(uri, data)
	})
}
func (i *Input) receive(uri string, data map[string]interface{}) {
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
		i.AI.Game.Log().With(map[string]interface{}{
			"URI":      uri,
			"GameId":   i.AI.Game.ID(),
			"Username": i.AI.Seat.Username,
		}).Warn("uri unknown")
	}
}
