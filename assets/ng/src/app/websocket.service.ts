import { Injectable } from '@angular/core'
import { Router } from '@angular/router'

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {

  private ws: WebSocket
  private routes = new Array<{
    uri:string
    f:any
  }>()

  constructor(private router: Router) {
    this.init()
    let me = this
    this.route('/redirect', function(val) {
      me.redirect(val.location)
    })
  }

  public ready(): boolean {
    return this.ws && (this.ws.readyState == WebSocket.OPEN || this.ws.readyState == WebSocket.CONNECTING)
  }

  public closed(): boolean {
    return this.ws && (this.ws.readyState == WebSocket.CLOSED || this.ws.readyState == WebSocket.CLOSING)
  }

  send(uri: string, data: object) {
    console.debug('ws out:', uri, data)
    this.ws.send(JSON.stringify({
      'uri': uri,
      'data': data
    }))
  }

  // websocket init

  public init() {
    if (this.ready()) return;
    this.ws = new WebSocket(window.location.protocol.replace('http', 'ws') + window.location.host + '/api/websocket')
    this.ws.onopen = this._ws_onopen(this)
    this.ws.onmessage = this._ws_onmessage(this)
    this.ws.onclose = this._ws_onclose(this)
  }

  // f(data:object) "no return value"
  public route(uri: string, f: any) {
    this.routes.push({
      uri:uri,
      f:f,
    })
  }

  _ws_onopen(ws : WebsocketService) {
    return function() {
      console.debug('ws opened')
    }
  }

  _ws_onmessage(ws : WebsocketService) {
    return function(msg) {
      if (msg.data instanceof Blob) {
        let reader = new FileReader()
        reader.onload = function () {
          let data = ws.parse(reader.result.toString())
          if (data) ws.serve(data)
          else return console.error("websocket.message", "blob parse failed", msg.data)
        }
        reader.readAsText(msg.data)
      } else {
        let data = ws.parse(msg.data.toString())
        if (data) ws.serve(data)
        else return console.error("websocket.message", "raw parse failed", msg.data)
      }
    }
  }

  _ws_onclose(ws : WebsocketService) {
    return function(e) {
      console.warn('ws connection lost', e)
      ws.redirect('/lost')
    }
  }

  // incoming util

  parse(msg: string) {
    try {
      return JSON.parse(msg)
    } catch (e) {
      console.error(e)
      console.warn('failed to parse ws message', msg)
      return false
    }
  }

  // Websocket mux

  serve(msg: any) {
    let uri = msg.uri
    if (!uri) {
      console.warn('ws in uri?', uri, msg)
      return
    }
    let data = msg.data
    console.debug('ws in', msg, data)
    for (let i = 0; i < this.routes.length; i++) {
      let route = this.routes[i]
      if (route.uri == uri) {
        route.f(data)
        return
      }
    }
    console.log('ws in?', uri, data)
  }

  // Websocket routes

  // url gets router.url
  url(): string {
    return this.router.url
  }

  redirect(location: string) {
    // let location = data.location
    console.debug('redirect', location)
    this.router.navigate([location])
  }
}
