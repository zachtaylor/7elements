package runtime

import "taylz.io/http/websocket"

func (rt *T) OnSocket(id string, oldWS, newWS *websocket.T) {
	go rt.onSocket(id, oldWS, newWS)
}

func (rt *T) onSocket(id string, oldWS, newWS *websocket.T) {
	log := rt.Logger.Add("Socket", id)
	if oldWS == nil && newWS != nil {
		if len(newWS.SessionID()) > 0 {
			log = log.Add("SessionID", newWS.SessionID())
			go rt.hydrate(newWS)
		}
		log.Trace("open")
	} else if oldWS != nil && newWS == nil {
		if len(oldWS.SessionID()) > 0 {
			log = log.Add("SessionID", oldWS.SessionID())
		}
		log.Trace("close")
	} else {
		log.Add("Old", oldWS).Add("New", newWS).Warn("weird")
	}
	rt.Ping()
}
