package json

import (
	"encoding/json"
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

func ItoS(i int) string {
	return I64toS(int64(i))
}

func (j Json) Write(w http.ResponseWriter) {
	w.Write([]byte(j.String()))
}

func (j Json) String() string {
	data, err := json.Marshal(j)
	if err != nil {
		return "jsonerror: " + err.Error()
	}
	return string(data)
}

func NewDecoder(r *http.Request) *json.Decoder {
	return json.NewDecoder(r.Body)
}
