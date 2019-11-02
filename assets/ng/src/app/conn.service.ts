import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { Chat, ChatMessage, Game, MyAccount, Notification, PingData, Settings, GlobalData } from './api'
import { Router } from '@angular/router'
import { HttpClient, HttpParams } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'

@Injectable({
  providedIn: 'root'
})
export class ConnService {
  chats = new Map<string, BehaviorSubject<Chat>>()
  myaccount$ = new BehaviorSubject<MyAccount>(null)
  game$ = new BehaviorSubject<Game>(null)
  global$ = new BehaviorSubject<GlobalData>(null)
  ping$ = new BehaviorSubject<PingData>(null)
  notifications$ = new BehaviorSubject<Array<Notification>>([])

  // Runtime settings; TODO load cookies
  settings$ = new BehaviorSubject<Settings>({
    deck: {
      id: 1,
      account: false,
    },
    game: {
      hand: false,
      zoom: null,
      target: {
        display: '',
        helper: '',
        send: false
      }
    }
  })

  private ws: WebSocket

  constructor(private router: Router, private http: HttpClient, private cookieService: CookieService) {
    this.http.get<GlobalData>(`/api/global.json`).subscribe(glob => {
      console.debug('glob', glob)
      this.global$.next(glob)
    })
    let me = this
    this.ws = new WebSocket(window.location.protocol.replace('http', 'ws') + window.location.host + '/api/websocket')
    this.ws.onopen = function () {
      console.debug('ws opened')
    }
    this.ws.onmessage = function (msg) {
      if (msg.data instanceof Blob) {
        let reader = new FileReader()
        reader.onload = function () {
          let data = me.parseWS(reader.result.toString())
          if (data) me.serveWS(data)
          else return console.error("websocket.message failed to parse", msg.data)
        }
        reader.readAsText(msg.data)
      } else {
        let data = me.parseWS(msg.data.toString())
        if (data) me.serveWS(data)
        else return console.error("websocket.message failed to parse", msg.data)
      }
    }
    this.ws.onclose = function (e) {
      console.warn('ws connection lost', e)
      me.myaccount$.next(null)
      me.game$.next(null)
    }
  }

  /**
   * isLoggedIn returns whether this.myaccount$ has been set
   * This is useful because templates do not execute when a Subscription returns null
   */
  isLoggedIn(): boolean {
    return (!!this.myaccount$.value) && (!!this.myaccount$.value.username)
  }

  /**
   * hasGame returns whether Game is set 
   * This is useful because templates do not execute when a Subscription returns null
   */
  hasGame(): boolean {
    return !!this.game$.value
  }

  /**
   * getGame returns Game
   * @hidden
   */
  getGame(): Game {
    return this.game$.value
  }

  getUsername() {
    if (!this.isLoggedIn()) {
      return ""
    }
    return this.myaccount$.value.username
  }

  // addNotification(level : string, source : string, message : string) {
  //   let next = this.notifications$.getValue()
  //   let not = new Notification()
  //   not.level = level
  //   not.source = source
  //   not.message = message
  //   next.push(not)
  //   this.notifications$.next(next)
  //   console.debug('added notification', not)
  // }

  clearNotifications() {
    this.notifications$.next([])
  }

  getCard(cardid: number) {
    let cards = this.global$.value.cards
    for (let i = 0; i < cards.length; i++) {
      if (cards[i].id == cardid) {
        return cards[i]
      }
    }
    return null
  }

  getDeck(deckid: number) {
    if (!this.global$.value) return null
    let decks = this.global$.value.decks
    for (let i = 0; i < decks.length; i++) {
      if (decks[i].id == deckid) {
        return decks[i]
      }
    }
    return null
  }

  getChat(channel: string): BehaviorSubject<Chat> {
    if (!this.chats[channel]) {
      let chat = new Chat()
      chat.name = channel
      this.chats[channel] = new BehaviorSubject<Chat>(chat)
    }
    return this.chats[channel]
  }

  newGame(ai: boolean, usep2p: boolean, deckid: number) {
    let path = '/api/newgame.json'
    if (ai) path += '?ai=true'
    this.http.get<Game>(path, {
      params: new HttpParams({
        fromObject: {
          "ai": ai.toString(),
          "deckid": deckid.toString(),
          "usep2p": usep2p.toString(),
        }
      })
    }).subscribe(match => {
      console.debug('new match found', match)
      this.game$.next(match)
      this.router.navigateByUrl('/play')
    })
  }

  sendWS(uri: string, data: object) {
    console.debug('ws out:', uri, data)
    this.ws.send(JSON.stringify({
      'uri': uri,
      'data': data
    }))
  }

  parseWS(msg: string) {
    try {
      return JSON.parse(msg)
    } catch (e) {
      console.error(e)
      console.warn('failed to parse ws message', msg)
      return false
    }
  }

  // Websocket Mux

  serveWS(msg: any) {
    let uri = msg.uri
    let data = msg.data
    if (!uri) {
      console.warn('ws in uri?', msg)
      return
    } else if (uri == '/redirect') {
      this.serveRedirect(data)
    } else if (uri == '/data/myaccount') {
      this.serveDataMyAccount(data)
    } else if (uri == '/data/ping') {
      this.servePingData(data)
    } else if (uri == '/chat') {
      this.serveChat(data)
    } else if (uri == '/chat/join') {
      this.serveChatJoin(data)
    } else if (uri == '/game') {
      this.serveGame(data)
    } else if (uri == '/game/state') {
      this.serveGameState(data)
    } else if (uri == '/game/seat') {
      this.serveGameSeat(data)
    } else if (uri == '/game/card') {
      this.serveGameCard(data)
    } else if (uri == '/game/hand') {
      this.serveGameHand(data)
    } else if (uri == '/game/choice') {
      this.serveGameChoice(data)
    } else if (uri == '/game/error') {
      this.serveGameError(data)
    } else if (uri == '/game/react') {
      this.serveGameReact(data)
    } else if (uri == '/game/spawn') {
      this.serveGameSpawn(data)
    } else {
      console.log('ws in?', uri, data)
    }
  }

  serveRedirect(data: any) {
    console.debug('conn.serveRedirect:', data.location)
    this.router.navigate([data.location])
  }

  serveDataMyAccount(data: any) {
    if (data.username) {
      // let date = new Date();
      // date.setDate(date.getDate() + 1);
      this.cookieService.set('SessionID', data.session, 1, undefined, undefined, window.location.protocol == 'https:')
      console.debug('conn.serveDataMyAccount:', data.username, data.session, data.cards)
    } else {
      console.debug('unset cookie')
      this.cookieService.delete('SessionID')
      this.router.navigateByUrl('/')
    }
    this.myaccount$.next(data)
  }

  servePingData(data: any) {
    console.debug('conn.servePingData:', data)
    this.ping$.next(data)
  }

  // serveAlert(data : any) {
  //   this.addNotification(data.level, data.source, data.text)
  // }

  serveGameError(data: any) {
    console.log('conn.serveGameError:', data.source, data.error)
    let nots = this.notifications$.value
    let not = new Notification()
    not.source = 'vii'
    not.message = ''
    if (data.source) {
      not.message = data.source + ': '
    }
    not.message += data.message
    nots.push(not)
    this.notifications$.next(nots)
    let me = this.notifications$
    window.setTimeout(() => {
      nots = me.value
      let i = nots.indexOf(not, 0)
      if (i >= 0) {
        nots.splice(i, 1)
      }
      me.next(nots)
    }, 7000)
  }

  serveChat(data: any) {
    console.debug('conn.serveChat:', data.channel, data.username, data.message)
    let chat$ = this.getChat(data.channel)
    let next = chat$.getValue()
    let msg = new ChatMessage()
    msg.channel = data.channel
    msg.username = data.username
    msg.message = data.message
    next.messages.unshift(msg)
    chat$.next(next)
    let nots = this.notifications$.value
    let not = new Notification()
    not.source = data.username
    not.message = data.message
    nots.push(not)
    this.notifications$.next(nots)
    let me = this.notifications$
    window.setTimeout(() => {
      nots = me.value
      let i = nots.indexOf(not, 0)
      if (i >= 0) {
        nots.splice(i, 1)
      }
      me.next(nots)
    }, 7000)
  }

  serveChatJoin(data: any) {
    console.debug('conn.serveChatJoin', data.channel)
    this.getChat(data.channel).next(data)
  }

  serveGame(data: any) {
    if (data) {
      console.debug('conn.serveGame', data.id, data)
      if (this.router.url != '/play') {
        this.router.navigateByUrl('/play')
      }
    }
    this.game$.next(data)
  }

  serveGameState(data: any) {
    console.debug('conn.serveGameState', data.seat, data.name, data.data)
    let game = this.getGame()
    game.state = data
    this.game$.next(game)
  }

  serveGameSeat(data: any) {
    console.debug('conn.serveGameSeat', data.username, data.elements)
    let game = this.getGame()
    let seat = game.seats[data.username]
    Object.keys(data).forEach(key => {
      seat[key] = data[key]
    })
    game.seats[data.username] = seat
    this.game$.next(game)
  }

  serveGameCard(data: any) {
    let game = this.getGame()
    let owner = data.username
    if (data.body && data.body.health < 1) {
      console.debug('conn.serveGameCard remove', data.gcid)
      delete game.seats[owner].active[data.gcid]
    } else if (!game.seats[owner].active[data.gcid]) {
      console.debug('conn.serveGameCard missing', owner, game.seats[owner].active, data.gcid)
    } else {
      console.debug('conn.serveGameCard update', data.gcid)
      game.seats[owner].active[data.gcid] = data
    }
    this.game$.next(game)
  }

  serveGameHand(data: any) {
    console.debug('conn.serveGameHand', data.cards)
    let game = this.getGame()
    game.hand = data.cards
    this.game$.next(game)
  }

  serveGameChoice(data: any) {
    console.debug('conn.serveGameChoice')
    let game = this.getGame()
    game.state.data = data
    this.game$.next(game)
  }

  serveGameReact(data: any) {
    console.debug('conn.serveGameReact', data.username, data.react)
    let game = this.getGame()
    if (!game) return;
    game.state.timer = data.timer
    game.state.reacts[data.username] = data.react
    this.game$.next(game)
  }

  serveGameSpawn(data: any) {
    let game = this.getGame()
    console.debug('conn.serveGameSpawn', data.username, data.card.name, game.seats[data.username].active)
    game.seats[data.username].active[data.card.gcid] = data.card
    this.game$.next(game)
  }
}
