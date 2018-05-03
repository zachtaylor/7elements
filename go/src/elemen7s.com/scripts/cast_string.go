package scripts

import (
	"strconv"
	"ztaylor.me/log"
)

func CastString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case float64:
		return strconv.FormatFloat(v, 'E', -1, 32)
	default:
		log.Add("Val", val).Error(CloningPoolID + ": failed to cast value")
		return ""
	}
}
