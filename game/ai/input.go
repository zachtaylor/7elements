package ai

import (
	"bytes"
	"encoding/json"
	"time"

	"taylz.io/http/websocket"
)

// Writer is where the ai receives input
type Writer struct {
	AI *AI
}

func (w *Writer) Name() string { return w.AI.Name }

func (w *Writer) Done() <-chan struct{} { return w.AI.done }

// func (w *Writer) WriteMessage(msg *websocket.Message) error {
// 	w.Receive(msg.URI, msg.Data)
// 	return nil
// }
func (w *Writer) Write(data []byte) error {
	msg := &websocket.Message{}
	if err := json.NewDecoder(bytes.NewBufferString(string(data))).Decode(msg); err != nil {
		w.AI.View.Game.Log().Add("Error", err).Warn("failed to parse message")
	} else {
		w.AI.View.Game.Log().Trace("received ", msg.URI)
	}
	w.Receive(msg.URI, msg.Data)
	return nil
}

// Receive data from the game runtime
func (w *Writer) Receive(uri string, data map[string]any) {
	go time.AfterFunc(w.AI.Delay, func() {
		w.receive(uri, data)
	})
}
func (w *Writer) receive(uri string, data map[string]any) {
	if uri == "/game/state" {
		w.AI.GameState(data)
	} else if uri == "/game/choice" {
		w.AI.GameChoice(data)
	} else if uri == "/game" {
	} else if uri == "/game/react" {
	} else if uri == "/game/card" {
	} else if uri == "/game/hand" {
	} else if uri == "/game/seat" {
	} else if uri == "/alert" {
	} else {
		w.AI.View.Game.Log().Warn("uri unknown", uri)
	}
}
