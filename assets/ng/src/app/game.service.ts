import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { Game, GameCard, GameToken, Overlay, Card, Power, OverlayChoice } from './api'
import { WebsocketService } from './websocket.service'
import { NotificationService } from './notification.service'

@Injectable({
  providedIn: 'root'
})
export class GameService {
  private cards = new Map<string, BehaviorSubject<GameCard>>()
  private tokens = new Map<string, BehaviorSubject<GameToken>>()
  hand$ = new BehaviorSubject<Array<GameCard>>([])
  overlay$ = new BehaviorSubject<Overlay>(null)
  data$ = new BehaviorSubject<Game>(null)

  constructor(private ws: WebsocketService, private notifications: NotificationService) {
    let me = this
    this.ws.route('/game', msg => {
      me.serveGame(msg)
    })
    this.ws.route('/game/state', msg => {
      me.serveState(msg)
    })
    this.ws.route('/game/seat', msg => {
      me.serveSeat(msg)
    })
    this.ws.route('/game/card', msg => {
      me.serveCard(msg)
    })
    this.ws.route('/game/token', msg => {
      me.serveToken(msg)
    })
    this.ws.route('/game/hand', msg => {
      me.serveHand(msg)
    })
    this.ws.route('/game/error', msg => {
      me.serveGameError(msg)
    })
    this.ws.route('/game/choice', msg => {
      me.serveGameChoice(msg)
    })
    this.ws.route('/game/react', msg => {
      me.serveGameReact(msg)
    })
  }

  card$(id: string) : BehaviorSubject<GameCard> {
    if (!this.cards[id]) {
      this.cards[id] = new BehaviorSubject<GameCard>(null)
    }
    return this.cards[id]
  }

  token$(id: string) : BehaviorSubject<GameToken> {
    if (!this.tokens[id]) {
      this.tokens[id] = new BehaviorSubject<GameToken>(null)
    }
    return this.tokens[id]
  }

  addOverlay(msg: string, card?: GameCard, token?: GameToken, choices?: Array<OverlayChoice>) {
    let overlay = new Overlay(msg, this.overlay$.value)
    if (choices) {
      overlay.choices = choices
    }
    this.overlay$.next(overlay)
  }


  // mux


  serveGame(data: any) {
    if (data) {
      console.debug('conn.serveGame', data.id, data)
      if (this.ws.url() != '/play') {
        this.ws.redirect('/play')
      }
    }
    this.data$.next(data)
  }

  serveState(data: any) {
    console.debug('conn.serveState', data.seat, data.name, data.data)
    let game = this.data$.value
    if (game) game.state = data
    this.data$.next(game)
    this.overlay$.next(null)
  }

  serveSeat(data: any) {
    console.debug('conn.serveSeat', data.username, data.elements)
    let game = this.data$.value
    let seat = game.seats[data.username]
    Object.keys(data).forEach(key => {
      seat[key] = data[key]
    })
    game.seats[data.username] = seat
    this.data$.next(game)
  }

  serveCard(data: any) {
    console.debug('conn.serveCard', data.id, data.username)
    this.card$(data.id).next(data)
  }

  serveToken(data: any) {
    let game = this.data$.value
    let owner = data.username
    if (data.body && data.body.health < 1) {
      console.debug('conn.serveToken remove', data.id)
      let index = game.seats[owner].indexOf(data.id)
      if (index > -1) {
        game.seats[owner].splice(index, 1)
      }
    } else if (game.seats[owner].present.indexOf(data.id) < -1) {
      console.debug('conn.serveToken add', data.id, owner, game.seats[owner].present)
      game.seats[owner].present.push(data.id)
    } else {
      console.debug('conn.serveToken update', data.id)
    }
    this.token$(data.id).next(data)
    this.data$.next(game)
  }

  serveHand(data: any) {
    console.debug('conn.serveGameHand', data.cards)
    let game = this.data$.value
    if (!game) return;
    game.hand = data.cards
    this.data$.next(game)
  }

  serveGameChoice(data: any) {
    console.debug('conn.serveGameChoice', data.prompt, data)
    this.addOverlay(data.prompt, null, null, data.choices)
  }

  serveGameReact(data: any) {
    console.debug('conn.serveGameReact', data.username, data.react)
    let game = this.data$.value
    if (!game) return;
    game.state.timer = data.timer
    game.state.reacts[data.username] = data.react
    this.data$.next(game)
  }

  serveGameSpawn(data: any) {
    let game = this.data$.value
    console.debug('conn.serveGameSpawn', data.username, data.card.name, game.seats[data.username].present)
    game.seats[data.username].present[data.card.id] = data.card
    this.data$.next(game)
  }
  
  serveGameError(data: any) {
    console.log('conn.serveGameError:', data.source, data.error)
    this.notifications.add(data.source, data.message, 7000)
  }
  

  // game send macros


  // sendplay send a play command
  sendplay(game: Game, card: GameCard) {
    if (card.type == 'spell' && card.powers.length > 0 && card.powers[0].target != 'self') {
      let me = this
      let overlay = new Overlay(card.powers[0].text, this.overlay$.value)
      overlay.target = card.powers[0].target
      overlay.targetF = function (val: any) {
        me.ws.send('/game', {
          'uri': 'play',
          'gameid': game.id,
          'id': card.id,
          'target': val
        })
      }
      this.overlay$.next(overlay)
    } else {
      this.ws.send('/game', {
        'uri': 'play',
        'gameid': game.id,
        'id': card.id
      })
    }
  }

  // sendtrigger sends a trigger command
  sendtrigger(game: Game, token: GameToken, power: Power) {
    if (power.target != '' && power.target != 'self') {
      let me = this
      let overlay = new Overlay(power.text, this.overlay$.value)
      overlay.token = token
      overlay.target = power.target
      overlay.targetF = function (val: any) {
        me._sendtrigger(game.id, power.id, token.id, val)
      }
      this.overlay$.next(overlay)
    } else {
      this._sendtrigger(game.id, power.id, token.id, token.id)
    }
  }
  _sendtrigger(gameid: string, powerid: number, tokenid: string, target: any) {
    this.ws.send('/game', {
      uri: 'trigger',
      gameid: gameid,
      powerid: powerid,
      id: tokenid,
      target: target,
    })
  }
}
