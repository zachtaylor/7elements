import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { Chat, ChatMessage, Game, MyAccount, Notification, PingData } from './api'
import { HttpClient, HttpParams } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'


@Injectable({
  providedIn: 'root'
})
export class ConnService {
  chats = new Map<string, BehaviorSubject<Chat>>()
  games = new Map<string, BehaviorSubject<Game>>()
  myaccount$ = new BehaviorSubject<MyAccount>(null)
  notifications$ = new BehaviorSubject<Array<Notification>>([])
  ping$ = new BehaviorSubject<PingData>(null)

  private ws: WebSocket

  constructor(private http : HttpClient, private cookieService : CookieService) {
    this.http.get<PingData>(`/api/ping.json`).subscribe(ping => {
      console.debug('conn.ping', ping)
      this.ping$.next(ping)
    })
    if (cookieService.get(`SessionID`)) {
      this.http.get<MyAccount>(`/api/myaccount.json`).subscribe(account => {
        console.debug('conn.myaccount', account)
        this.myaccount$.next(account)
        if (!account) this.cookieService.delete(`SessionID`)
      });
    }
    let me = this
    this.ws = new WebSocket(window.location.protocol.replace('http', 'ws')+window.location.host+'/api/websocket')
    this.ws.onopen = function() {
      console.debug('ws opened')
    }
    this.ws.onmessage = function(msg) {
      if (msg.data instanceof Blob) {
        let reader = new FileReader()
        reader.onload = function() {
          let data = me.parseWS(reader.result.toString())
          if (data) me.serveWS(data.uri, data.data)
          else return console.error("websocket.message failed to parse", msg.data)
        }
        reader.readAsText(msg.data)
      } else {
        let data = me.parseWS(msg.data.toString())
        if (data) me.serveWS(data.uri, data.data)
        else return console.error("websocket.message failed to parse", msg.data)
      }
    }
    this.ws.onclose = function(e) {
      console.warn('ws connection lost', e)
    }
  }

  isLoggedOut() {
    return !this.myaccount$.value
  }

  addNotification(level : string, source : string, message : string) {
    let next = this.notifications$.getValue()
    let not = new Notification()
    not.level = level
    not.source = source
    not.message = message
    next.push(not)
    this.notifications$.next(next)
  }

  clearNotifications() {
    this.notifications$.next([])
  }

  getCard(ping : PingData, cardid : number) {
    let cards = ping.cards
    for (let i=0; i<cards.length; i++) {
      if (cards[i].id == cardid) {
        return cards[i]
      }
    }
    return null
  }

  getDeck(ping : PingData, deckid : number) {
    let decks = ping.decks
    for (let i=0; i<decks.length; i++) {
      if (decks[i].id == deckid) {
        return decks[i]
      }
    }
    return null
  }

  getChat(channel : string) : BehaviorSubject<Chat> {
    if (!this.chats[channel]) {
      let chat = new Chat()
      chat.name = channel
      this.chats[channel] = new BehaviorSubject<Chat>(chat)
    }
    return this.chats[channel]
  }

  getGame(key : string) : BehaviorSubject<Game> {
    if (!this.games[key]) {
      this.games[key] = new BehaviorSubject<Game>(null)
    }
    return this.games[key]
  }

  newGame(ai : boolean, usep2p : boolean, deckid : number) {
    let path = '/api/newgame.json'
    if (ai) path += '?ai=true'
    this.http.get(path, {
      params:new HttpParams({
        fromObject:{
          "ai":ai.toString(),
          "deckid":deckid.toString(),
          "usep2p":usep2p.toString(),
        }
      })
    }).subscribe(match => {
      console.log('new game', match)
    })
  }

  sendWS(uri : string, obj : object) {
    this.ws.send(JSON.stringify({
      "uri":uri,
      "data":obj
    }))
  }

  parseWS(msg : string) {
    try {
      return JSON.parse(msg)
    } catch (e) {
      console.error(e)
      return false
    }
  }

  serveWS(uri : string, data : object) {
    if (uri=='/error') {
      this.serveError(data)
    } else if (uri=='/chat') {
      this.serveChat(data)
    } else if (uri=='/chat/join') {
      this.serveChatJoin(data)
    } else {
      console.debug('ws serve: ', uri, data)
    }
  }

  serveError(data : any) {
    this.addNotification('error', data.source, data.message)
  }

  serveChat(data : any) {
    console.debug('serve chat: ', data.channel, data.username, data.message, data.time)
    let chat$ = this.getChat(data.channel)
    let next = chat$.getValue()
    let msg = new ChatMessage()
    msg.channel = data.channel
    msg.username = data.username
    msg.message = data.message
    msg.time = data.time
    next.messages.unshift(msg)
    chat$.next(next)
  }

  serveChatJoin(data : any) {
    console.debug('serve chat join: ', data.channel, data.username, data.messages)
    this.getChat(data.channel).next(data)
  }
}
