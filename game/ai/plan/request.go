package plan

import "taylz.io/http/websocket"

type RequestFunc = func(string, websocket.MsgData)
