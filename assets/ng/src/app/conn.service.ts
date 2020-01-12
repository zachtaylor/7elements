import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { Chat, ChatMessage, Game, MyAccount, Notification, PingData, Settings, GlobalData, Overlay } from './api'
import { Router } from '@angular/router'
import { HttpClient, HttpParams } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'
import { GameService } from './game.service'
import { WebsocketService } from './websocket.service'
import { NotificationService } from './notification.service'

@Injectable({
  providedIn: 'root'
})
export class ConnService {
  chats = new Map<string, BehaviorSubject<Chat>>()
  myaccount$ = new BehaviorSubject<MyAccount>(null)
  ping$ = new BehaviorSubject<PingData>(null)
  global$ = new BehaviorSubject<GlobalData>(null)

  // Runtime settings; TODO load cookies
  settings$ = new BehaviorSubject<Settings>(new Settings())

  constructor(
    private router: Router,
    private http: HttpClient,
    private cookieService: CookieService,
    private ws: WebsocketService,
    private notifications: NotificationService,
  ) {
    let me = this
    this.http.get<GlobalData>(`/api/global.json`).subscribe(glob => {
      console.debug('glob', glob)
      this.global$.next(glob)
    })
    me.ws.route('/myaccount', function(msg) {
      me.serveMyAccount(msg)
    })
    me.ws.route('/ping', function(msg) {
      me.servePing(msg)
    })
    me.ws.route('/chat', function(msg) {
      me.serveChat(msg)
    })
    me.ws.route('/chat/join', function(msg) {
      me.serveChatJoin(msg)
    })
    me.ws.route('/chat/game', function(msg) {
      me.serveChatGame(msg)
    })
  }

  /**
   * isLoggedIn returns whether this.myaccount$ has been set
   * This is useful because templates do not execute when a Subscription returns null
   */
  isLoggedIn(): boolean {
    return (!!this.myaccount$.value) && (!!this.myaccount$.value.username)
  }

  username() {
    if (!this.isLoggedIn()) {
      return ""
    }
    return this.myaccount$.value.username
  }

  card(cardid: number) {
    let cards = this.global$.value.cards
    for (let i = 0; i < cards.length; i++) {
      if (cards[i].id == cardid) {
        return cards[i]
      }
    }
    return null
  }

  deck(deckid: number) {
    if (!this.global$.value) return null
    let decks = this.global$.value.decks
    for (let i = 0; i < decks.length; i++) {
      if (decks[i].id == deckid) {
        return decks[i]
      }
    }
    return null
  }

  chat$(channel: string): BehaviorSubject<Chat> {
    if (!this.chats[channel]) {
      let chat = new Chat()
      chat.name = channel
      this.chats[channel] = new BehaviorSubject<Chat>(chat)
    }
    return this.chats[channel]
  }

  // newGame(ai: boolean, usep2p: boolean, deckid: number) {
  //   let path = '/api/newgame.json'
  //   if (ai) path += '?ai=true'
  //   this.http.get<Game>(path, {
  //     params: new HttpParams({
  //       fromObject: {
  //         "ai": ai.toString(),
  //         "deckid": deckid.toString(),
  //         "usep2p": usep2p.toString(),
  //       }
  //     })
  //   }).subscribe(match => {
  //     console.debug('new match found', match)
  //     this.game.data$.next(match)
  //     this.redirect('/play')
  //   })
  // }


  // Websocket Mux


  serveMyAccount(data: any) {
    if (data.username) {
      this.cookieService.set('SessionID', data.session, 1, undefined, undefined, window.location.protocol == 'https:')
      this.notifications.add('vii', 'account updated', 3000)
      console.debug('conn.serveMyAccount:', data.username, data.session, data)
    } else {
      console.debug('conn.serveMyAccount: unset cookie')
      this.cookieService.delete('SessionID')
      this.ws.redirect('/')
    }
    this.myaccount$.next(data)
  }

  servePing(data: any) {
    if (!data || !data.ping) {
      console.debug('conn.servePing keepalive')
      this.ws.send('/ping', null)
    } else {
      console.debug('conn.servePing:', data)
      this.ping$.next(data)
    }
  }

  // serveAlert(data : any) {
  //   this.addNotification(data.level, data.source, data.text)
  // }

  serveChat(data: any) {
    console.debug('conn.serveChat:', data.channel, data.username, data.message)
    let chat$ = this.chat$(data.channel)
    let next = chat$.getValue()
    let msg = new ChatMessage()
    msg.channel = data.channel
    msg.username = data.username
    msg.message = data.message
    next.messages.unshift(msg)
    chat$.next(next)
    this.notifications.add(data.username, data.message, 10000)
  }

  serveChatJoin(data: any) {
    console.debug('conn.serveChatJoin', data.channel)
    this.chat$(data.channel).next(data)
  }

  serveChatGame(data:any) {
    console.debug('conn.serveChatGame:', data.channel, data.username, data.message)
    this.notifications.add(data.username, data.message, this.settings$.value.game.chat.hideT)
  }
}
