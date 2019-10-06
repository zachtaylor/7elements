package vii

import "ztaylor.me/cast"

// JSONWriter is a Writer that accepts JSON
type JSONWriter interface {
	WriteJSON(cast.JSON)
}
