import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core'
import { Overlay, Game, GameToken, Card, OverlayChoice, GameCard, ElementCount, Power } from '../api'
import { Subscription } from 'rxjs'
import { WebsocketService } from '../websocket.service'
import { ConnService } from '../conn.service'
import { GameService } from '../game.service'

@Component({
  selector: 'overlay',
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './overlay.component.html',
  styleUrls: ['./overlay.component.css']
})
export class OverlayComponent implements OnInit {
  game: Game
  overlay: Overlay
  target: string

  private $game: Subscription
  private $overlay: Subscription

  constructor(public ws : WebsocketService, public conn: ConnService, public games: GameService) {

  }

  ngOnInit() {
    let me = this
    this.$game = this.games.data$.subscribe(game => {
      this.game = game
    })
    this.$overlay = this.games.overlay$.subscribe(overlay => {
      me.overlay = overlay
    })
  }
  ngOnDestroy() {
    this.$game.unsubscribe()
    this.$overlay.unsubscribe()
  }

  close() {
    this.games.overlay$.next(this.overlay.stack)
  }

  targettokens(seat: string, target: string): Array<GameToken> {
    let targets = []
    if (target=='being') {
      let tokenids = this.game.seats[seat].present
      for (let j = 0; j < tokenids.length; j++) {
        let token = this.games.token$(tokenids[j]).value
        if (token.body) {
          targets.push(token)
        }
      }
    } else if (target=='being-item') {
      let tokenids = this.game.seats[seat].present
      for (let j = 0; j < tokenids.length; j++) {
        let token = this.games.token$(tokenids[j]).value
        targets.push(token)
      }
    }
    return targets
  }

  targetpastcards(seat: string, target: string): Array<GameCard> {
    console.log('targetpastcards')
    let targets = []
    if (target=='past-being-item') {
      let cardids = this.game.seats[seat].past
      for (let j = 0; j < cardids.length; j++) {
        let card = this.games.card$(cardids[j]).value
        targets.push(card)
      }
    } else if (target=='mypast-being' && seat==this.game.username) {
      let cardids = this.game.seats[seat].past
      console.log('mypast-being found past', cardids)
      for (let j = 0; j < cardids.length; j++) {
        let card$ = this.games.card$(cardids[j])
        let card = card$.value
        console.log('mypast-being card', cardids[j], ' is ' , card)
        if (card && card.body) {
          targets.push(card)
        }
      }
    }
    return targets
  }

  testkarma(card: Card) {
    let karmacount = new ElementCount()
    let keys = Object.keys(this.game.seats[this.game.username].elements)
    for (let i = 0; i < keys.length; i++) {
      let el = keys[i]
      let set = this.game.seats[this.game.username].elements[el]
      for (let j = 0; j < set.length; j++) {
        if (set[j]) {
          if (karmacount.items[el]) {
            karmacount.items[el] += 1
          } else {
            karmacount.items[el] = 1
          }
        }
      }
    }
    return karmacount.test(card.costs)
  }

  testplayable() {
    return this.overlay.card.type=='spell'||this.game.state.name=='main'
  }

  testtarget(arg : any) {
    if (!this.overlay.target) return false
    else if (this.overlay.target=='player-being' && arg.username) return true
    else if (this.overlay.target=='being' && arg.body) return true
    else if (this.overlay.target=='being-item' && (arg.id+"").length > 1) return true
  }

  sendplay() {
    let card = this.overlay.card
    this.close()
    this.games.sendplay(this.game, card)
  }

  sendpower(power: Power) {
    let token = this.overlay.token
    this.close()
    this.games.sendtrigger(this.game, token, power)
  }

  sendtarget(target: string) {
    let f = this.overlay.targetF
    this.close()
    f(target)
  }

  // attacks

  showattack(): boolean {
    return this.overlay.token.awake&&
      this.overlay.token.body&&
      this.game.state.name=='main'&&
      this.game.state.seat==this.game.username
  }

  sendattack(id: string) {
    this.ws.send('/game', {
      gameid: this.game.id,
      uri: 'attack',
      id: id,
    })
    this.close()
  }

  // defends

  showdefend(): boolean {
    return this.overlay.token.awake&&
      this.overlay.token.body&&
      this.game.state.name=='attack'&&
      this.game.state.seat!=this.game.username
  }

  senddefend(id: string) {
    this.ws.send('/game', {
      gameid: this.game.id,
      uri: this.game.state.id,
      id: id,
    })
    this.close()
  }

  // choices

  sendchoice(game: Game, choice: string) {
    this.ws.send('/game', {
      gameid: game.id,
      uri: game.state.id,
      choice: choice,
    })
    this.close()
  }
}
