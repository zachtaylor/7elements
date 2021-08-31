import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { GlobalData, PingData, Message, MyAccount, GameState, GameMenu, GameCard, GameToken, GameSeat, GameMenuChoice, Power } from './api'
import { HttpClient } from '@angular/common/http'
import { Router, Event, NavigationEnd } from '@angular/router'
import { CookieService } from 'ngx-cookie-service'

@Injectable({
  providedIn: 'root'
})
export class VII {
  path = "/"

  glob : GlobalData
  global$ = new BehaviorSubject<GlobalData>(null)

  ping$ = new BehaviorSubject<PingData>(null)

  messages$ = new BehaviorSubject<Message[]>([])

  account$ = new BehaviorSubject<MyAccount>(null)

  gamestate$ = new BehaviorSubject<GameState>(null)

  gamehand$ = new BehaviorSubject([])

  gamemenu$ = new BehaviorSubject<GameMenu>(null)

  constructor(private http : HttpClient, private router: Router, private cookies : CookieService) {
    this.http.get<GlobalData>(`/api/global.json`).subscribe(glob => { this.serveGlobal(glob) })
    this.wsRoutes()
    this.router.events.subscribe(e => { this.onRouter(e)})
    setTimeout(() => { this.init() })
  }
  private init() {
    this.ws = new WebSocket(window.location.protocol.replace('http', 'ws') + window.location.host + '/api/websocket')
    this.ws.onopen = () => { this.wsOpen() }
    this.ws.onmessage = msg => { this.wsMessage(msg) }
    this.ws.onclose = err => { this.wsClose(err) }
  }
  private onRouter(e : Event) {
    if (e instanceof NavigationEnd) {
      this.path = window.location.pathname
      document.getElementsByTagName('content-wrapper')[0].scrollTop = 0
    }
  }

  private ws: WebSocket
  private routes = new Map<string, any>()
  private wsRoutes() {
    this.wsRoute('/ping', data => { this.servePing(data) })
    this.wsRoute('/redirect', data => { this.serveRedirect(data) })
    this.wsRoute('/chat/message', data => { this.serveChatMessage(data) })
    this.wsRoute('/chat/join', data => { this.serveChatJoin(data) })
    this.wsRoute('/myaccount', data => { this.serveMyAccount(data) })
    this.wsRoute('/myaccount/game', data => { this.serveMyAccountGame(data) })
    this.wsRoute('/game', data => { this.serveGame(data) })
    this.wsRoute('/game/hand', data => { this.serveGameHand(data) })
    this.wsRoute('/game/card', data => { this.serveGameCard(data) })
    this.wsRoute('/game/token', data => { this.serveGameToken(data) })
    this.wsRoute('/game/present', data => { this.serveGamePresent(data) })
    this.wsRoute('/game/state', data => { this.serveGameState(data) })
    this.wsRoute('/game/seat', data => { this.serveGameSeat(data) })
    this.wsRoute('/game/react', data => { this.serveGameReact(data) })
    this.wsRoute('/error', data => { this.serveError(data) })
  }
  private wsRoute(uri: string, f: any) {
    this.routes.set(uri, f)
  }
  private wsOpen () {
    console.debug('ws opened')
  }
  private wsMessage(msg : any) {
    if (msg.data instanceof Blob) {
      let reader = new FileReader()
      reader.onload = () => {
        let data = this.wsParse(reader.result.toString())
        if (data) this.wsServe(data)
        else console.error("websocket.message", "blob parse failed", msg.data)
      }
      reader.readAsText(msg.data)
    } else {
      let data = this.wsParse(msg.data.toString())
      if (data) this.wsServe(data)
      else console.error("websocket.message", "raw parse failed", msg.data)
    }
  }
  private wsParse(msg) : any {
    try {
      return JSON.parse(msg)
    } catch (e) {
      console.error('ws msg parse', msg, e)
      return false
    }
  }
  private wsServe(msg) {
    let uri = msg.uri
    if (!uri) {
      console.warn('received no uri', msg)
      return
    }
    let data = msg.data
    // console.debug('received', uri)
    let route = this.routes.get(uri)
    if (!route) console.warn('received unknown uri', uri, data)
    else route(data)
  }
  private wsClose(e) {
    console.warn('ws connection lost', e)
    this.goto('/lost')
  }

  private serveGlobal(data) {
    this.glob = data
    this.global$.next(this.glob)
  }
  private servePing(data) {
    if (!data || !data.ping) {
      console.debug('ping keep alive')
      this.send('/ping', null)
    } else {
      console.debug('ping data', data)
      this.ping$.next(data)
    }
  }
  private serveRedirect(data) {
    console.debug('redirect', data.location)
    this.goto(data.location)
  }
  private serveChatMessage(data) {
    console.debug('chat message', data.username, data.channel, data.message)
    let msgs = this.chat$(data.channel).value.slice()
    let msg = new Message(data.username, data.channel, data.message, data.time)
    msgs.unshift(msg)
    this.chat$(data.channel).next(msgs)
    this.message(msg.source, msg.channel, msg.message, msg.time, 10000)
  }
  private serveChatJoin(data) {
    console.debug('chat join', data.channel)
    this.chat$(data.channel).next(data)
  }
  private serveMyAccount(data) {
    if (data && data.username) {
      this.cookies.set('SessionID', data.session, 1, undefined, undefined, window.location.protocol == 'https:', 'Strict')
      this.message('vii', 'account', 'updated', '', 3000)
      console.debug('myaccount:', data.username, data.session, data)
    } else {
      console.debug('myaccount: unset')
      this.cookies.delete('SessionID')
      this.goto('/')
    }
    this.account$.next(data)
  }
  private serveMyAccountGame(data) {
    console.debug('/myaccount/game', data.game)
    let account = this.account$.value
    account.game = data.game
    this.account$.next(account)
  }
  private serveGame(data) {
    if (!data) {
      this.serveMyAccountGame({game:''})
      return
    }
    console.debug('/game', data.username)
    if (this.path != '/play') this.goto('/play')
    if (data.seats && data.seats.length) {
      for (let i = 0; i < data.seats.length; i++) {
        let name = data.seats[i]
        this.seat$(name).next(new GameSeat(name))
      }
    } else {
      console.warn('seats missing', data)
    }
  }
  private serveGameHand(data) {
    console.debug('serve hand', data.cards)
    this.gamehand$.next(data.cards)
  }
  private serveGameCard(data) {
    console.debug('/game/card', data.id, data.name)
    this.card$(data.id).next(data)
  }
  private serveGameToken(data) {
    console.debug('/game/token', data.id, data.name, data.user)
    this.token$(data.id).next(data)
  }
  private serveGamePresent(data) {
    console.debug('/game/present', data.username, data.present)
    if (!data) { return }
    let present$ = this.present$(data.username)
    if (data.present && data.present.length) { present$.next(data.present) }
    else { console.warn('present missing', data) }
  }
  private serveGameSeat(data: any) {
    console.debug('/game/seat', data.username)
    let seat$ = this.seat$(data.username)
    if (!seat$.value) { return }
    let seat = seat$.value
    if (seat.life !== data.life) { seat.life = data.life }
    if (seat.hand !== data.hand) { seat.hand = data.hand }
    if (seat.future !== data.future) { seat.future = data.future }
    if (seat.past !== data.past) { seat.past = data.past }
    for (let i = 1; i < 8; i++) {
      // console.debug('/game/seat/elements', i, data.elements[i])
      let elLights = data.elements[i]
      if ((!elLights) || (!elLights.length)) { continue }
      seat.elements[i] = elLights
    }
    seat$.next(seat)
    this.present$(seat.username).next(data.present)
  }
  private serveGameState(data) {
    console.debug('/game/state', data.id, data.name, data.seat)
    this.gamestate$.next(data)
    this.state$(data.id).next(data)
    if (data.stack) {
      let child$ = this.state$(data.stack)
      let child = child$.value
      if (!child) {
        console.warn('stack missing child', data.stack)
        data.stack = null
      } else {
        child.parent = data.id
        child$.next(child)
      }
    }
    let user = this.user()
    if (data.name == 'target' && data.seat == user) {
      this.overlayTarget(data.data)
    } else if (data.name == 'start' && !data.reacts[user]) {
      this.overlayStart()
    } else if (data.name == 'sunrise' && !data.reacts[user]) {
      this.overlayNewElement()
    } else if (data.name == 'choice' && data.seat == user && !data.reacts[user]) {
      console.log(data.data)
      this.overlay(data.data.choice, null, null, data.data.options)
    } else {
      this.gamemenu$.next(null)
    }
  }
  private serveGameReact(data: any) {
    console.debug('/game/react', data.stateid, data.username, data.react)
    let state$ = this.state$(data.stateid)
    let state = state$.value
    if (state.reacts[data.username]) return;
    state.reacts[data.username] = data.react
    state$.next(state)
  }
  private serveError(data) {
    console.log('/error', data.source, data.error)
    this.message('vii', data.source, data.error, '', 7000)
  }

  ready(): boolean {
    return this.ws && (this.ws.readyState == WebSocket.OPEN || this.ws.readyState == WebSocket.CONNECTING)
  }

  user() : string {
    if (this.account$.value) return this.account$.value.username
    return ''
  }

  private gamepresents = new Map<string, BehaviorSubject<string[]>>()
  present$(name : string) : BehaviorSubject<string[]> {
    if (!this.gamepresents.get(name)) this.gamepresents.set(name, new BehaviorSubject<string[]>([]))
    return this.gamepresents.get(name)
  }

  private chats = new Map<string, BehaviorSubject<Message[]>>()
  chat$(channel: string): BehaviorSubject<Message[]> {
    if (!this.chats.get(channel)) this.chats.set(channel, new BehaviorSubject<Message[]>([]))
    return this.chats.get(channel)
  }

  private cards = new Map<string, BehaviorSubject<GameCard>>()
  card$(id: string) : BehaviorSubject<GameCard> {
    if (!this.cards.get(id)) this.cards.set(id, new BehaviorSubject<GameCard>(null))
    return this.cards.get(id)
  }

  private tokens = new Map<string, BehaviorSubject<GameToken>>()
  token$(id: string) : BehaviorSubject<GameToken> {
    if (!this.tokens.get(id)) this.tokens.set(id, new BehaviorSubject<GameToken>(null))
    return this.tokens.get(id)
  }

  private seats = new Map<string, BehaviorSubject<GameSeat>>()
  gameready() : boolean { return this.seats.size > 1 }
  players() : IterableIterator<string> { return this.seats.keys() }
  opponents() : string[] {
    let me = this.user()
    let opponents = []
    this.seats.forEach(function(seat$, name) {
      if (name != me) opponents.push(name)
    })

    return opponents
  }
  seat$(id : string) : BehaviorSubject<GameSeat> {
    if (!this.seats.get(id)) this.seats.set(id, new BehaviorSubject<GameSeat>(null))
    return this.seats.get(id)
  }

  private states = new Map<string, BehaviorSubject<GameState>>()
  state$(id : string) : BehaviorSubject<GameState> {
    if (!this.states.get(id)) this.states.set(id, new BehaviorSubject<GameState>(null))
    return this.states.get(id)
  }

  private overlay(msg: string, card?: GameCard, token?: GameToken, choices?: Array<GameMenuChoice>) {
    console.debug('overlay', msg, choices)
    let menu = new GameMenu(msg, this.gamemenu$.value)
    if (card) menu.card = card
    if (token) menu.token = token
    if (choices) menu.choices = choices
    this.gamemenu$.next(menu)
  }
  private overlayStart() {
    let overlay = new GameMenu('Starting Hand', null)
    overlay.important = true
    overlay.choices.push({
      choice: 'keep',
      display: '<label>Keep</label>'
    })
    overlay.choices.push({
      choice: 'mulligan',
      display: '<label>Mulligan</label>'
    })
    this.gamemenu$.next(overlay)
  }
  private overlayNewElement() {
    let overlay = new GameMenu('Create a new Element', null)
    overlay.important = true
    for (let i = 1; i < 8; i++) {
      overlay.choices.push({
        choice: ''+i,
        display: '<img src="/img/icon/element-'+i+'.png">',
      })
    }
    this.gamemenu$.next(overlay)
  }
  private overlayTarget(data : any) {
    console.debug('overlay target', data.text, data.target)
    let me = this
    let overlay = new GameMenu(data.text, null)
    overlay.important = true
    overlay.target = data.helper
    overlay.targetF = val => {
      this.send('/game', {
        'uri': me.gamestate$.value.id,
        'target': val,
      })
    }
    this.gamemenu$.next(overlay)
  }

  message(source : string, channel : string, message : string, time : string, timeout : number) {
    let msgs = this.messages$.value.slice(), msg = new Message(source, channel, message, time)
    msgs.push(msg)
    this.messages$.next(msgs)
    setTimeout(() => {
      msgs = this.messages$.value.slice() // reset
      let i = msgs.indexOf(msg)
      if (i >= 0) {
        msgs.splice(i, 1)
        this.messages$.next(msgs)
      }
    }, timeout)
  }

  goto(location: string) {
    console.debug('goto', location)
    this.router.navigateByUrl(location)
  }

  send(uri: string, data: object) {
    console.debug('send', uri, data)
    this.ws.send(JSON.stringify({
      'uri': uri,
      'data': data
    }))
  }
  sendmessage(channel : string, message : string) {
    console.warn('send message', channel, message)
  }
  sendpass() {
    console.debug('send pass')
    let state = this.gamestate$.value
    state.reacts[this.user()] = "pass"
    this.send('/game', {
      uri: 'pass',
      pass: state.id,
    })
    this.gamestate$.next(state)
  }
  sendplay(gc: GameCard, pay : object) {
    let paystr = this.paystring(pay)
    let card = this.global$.value.cards[gc.cardid-1]
    if (card.type == 'spell' && card.powers.length > 0 && card.powers[0].target != 'self') {
      let menu = new GameMenu(card.powers[0].text, this.gamemenu$.value)
      menu.target = card.powers[0].target
      menu.targetF = val => {
        this.send('/game', {
          'uri': 'play',
          'pay': paystr,
          'id': gc.id,
          'target': val
        })
      }
      this.gamemenu$.next(menu)
    } else {
      this.send('/game', {
        'uri': 'play',
        'pay': paystr,
        'id': gc.id
      })
    }
  }
  sendtrigger(token: GameToken, power: Power, pay? : object) {
    let paystr = this.paystring(pay)
    if (power.target != '' && power.target != 'self') {
      let menu = new GameMenu(token.name + ' : Target ' + power.target, this.gamemenu$.value)
      menu.token = token
      menu.target = power.target
      let me = this
      menu.targetF = function (val: any) {
        me.send('/game', {
          uri: 'trigger',
          id: token.id,
          pay: paystr,
          powerid: power.id,
          target: val,
        })
      }
      this.gamemenu$.next(menu)
    } else {
      this.send('/game', {
        uri: 'trigger',
        id: token.id,
        pay: paystr,
        powerid: power.id,
      })
    }
  }
  private paystring(pay : object) : string {
    let me = this, str = ""
    Object.keys(pay).forEach(el => {
      let element = +el
      let count = pay[element]
      let letter = me.elementletter(element)
      for (let i=0; i<count; i++) {
        str += letter
      }
    })
    return str
  }
  private elementletter(i : number) : string {
    if (i == 1) return 'W'
    else if (i == 2) return 'R'
    else if (i == 3) return 'Y'
    else if (i == 4) return 'G'
    else if (i == 5) return 'B'
    else if (i == 6) return 'V'
    else if (i == 7) return 'A'
    else return 'X'
  }

  activeElements() : object {
    let seat$ = this.seat$(this.user()), seat = seat$.value, active = {}
    Object.keys(seat.elements).forEach(function(element) {
      for (let i=0; i < seat.elements[element].length; i++) {
        if (seat.elements[element][i]) {
          active[element] = active[element] ? active[element]+1 : 1;
        }
      }
    })
    return active
  }
  targetTokens(seat : GameSeat, target: string): Array<GameToken> {
    let targets = []
    if (target=='being') {
      let tokenids = this.present$(seat.username).value
      for (let j = 0; j < tokenids.length; j++) {
        let token = this.token$(tokenids[j]).value
        if (token.body) targets.push(token)
      }
    } else if (target=='being-item') {
      let tokenids = this.present$(seat.username).value
      for (let j = 0; j < tokenids.length; j++) {
        let token = this.token$(tokenids[j]).value
        targets.push(token)
      }
    } else if (target=='my-being') {
      if (seat.username != this.user()) return []
      let tokenids = this.present$(seat.username).value
      for (let j = 0; j < tokenids.length; j++) {
        let token = this.token$(tokenids[j]).value
        if (token.body) targets.push(token)
      }
    } else if (target=='player-other-being') {
      let tokenids = this.present$(seat.username).value
      for (let j = 0; j < tokenids.length; j++) {
        if (tokenids[j] == this.gamemenu$.value.token.id) continue
        let token = this.token$(tokenids[j]).value
        if (token.body) targets.push(token)
      }
    }
    console.log('target tokens', seat.username, target, targets)
    return targets
  }
  targetPastCards(seat : GameSeat, target: string): Array<GameCard> {
    let targets = []
    if (target=='past-being-item') {
      let cardids = seat.past
      for (let j = 0; j < cardids.length; j++) {
        let gamecard = this.card$(cardids[j]).value
        let card = this.global$.value.cards[gamecard.cardid - 1]
        if (card.type !== 'spell') {
          targets.push(gamecard)
        }
      }
    } else if (target=='mypast-being') {
      if (seat.username != this.user()) return []
      let cardids = seat.past
      console.debug('available past cardids', cardids)
      for (let j = 0; j < cardids.length; j++) {
        let gamecard = this.card$(cardids[j]).value
        let card = this.glob.cards[gamecard.cardid - 1]
        if (card.body) {
          targets.push(gamecard)
        }
      }
    }
    console.log('target past cards', seat.username, target, targets)
    return targets
  }
}
