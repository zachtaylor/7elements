package json

import (
	"7elements.ztaylor.me/log"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Json map[string]interface{}

type Jsoner interface {
	Json() Json
}

func I64toS(i int64) string {
	return strconv.FormatInt(i, 10)
}

func UItoS(ui uint) string {
	return I64toS(int64(ui))
}

func (j Json) Write(w http.ResponseWriter) {
	data, err := json.Marshal(j)
	if err != nil {
		log.Clone().Add("Error", err).Error("json: write")
	}

	w.Write(data)
}

func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}
