import { Component, OnInit } from '@angular/core'
import { ActivatedRoute, Router } from '@angular/router'
import { ConnService } from '../conn.service'
import { PingData, Game, GameCard, MyAccount, Power, Settings, GlobalData, GameToken, GameSeat, Card, Overlay } from '../api'
import { Subscription, Observable, interval } from 'rxjs'
import { GameService } from '../game.service'
import { WebsocketService } from '../websocket.service'

@Component({
  selector: 'app-play',
  templateUrl: './play.component.html',
  styleUrls: ['./play.component.css']
})
export class PlayComponent implements OnInit {
  glob: GlobalData
  myaccount: MyAccount
  timer: number
  game: Game
  overlay: Overlay
  settings: Settings
  private $ticker: Subscription
  private $glob: Subscription
  private $myaccount: Subscription
  private $game: Subscription
  private $overlay: Subscription
  private $settings: Subscription

  constructor(
    public conn: ConnService,
    public ws: WebsocketService,
    public games : GameService,
    private router: Router,
  ) { }

  ngOnInit() {
    console.debug('play ngOnInit')
    this.$ticker = interval(1000).subscribe(_ => {
      if (this.timer > 0) this.timer--
    })
    this.$glob = this.conn.global$.subscribe(glob => {
      this.glob = glob
    })
    this.$myaccount = this.conn.myaccount$.subscribe(myaccount => {
      this.myaccount = myaccount
    })
    this.$game = this.games.data$.subscribe(game => {
      this.game = game
      if (game) this.timer = game.state.timer
    })
    this.$overlay = this.games.overlay$.subscribe(overlay => {
      this.overlay = overlay
    })
    this.$settings = this.conn.settings$.subscribe(settings => {
      this.settings = settings
    })
  }

  ngOnDestroy() {
    console.debug('play ngOnDestroy')
    this.$ticker.unsubscribe()
    this.$glob.unsubscribe()
    this.$myaccount.unsubscribe()
    this.$game.unsubscribe()
    this.$overlay.unsubscribe()
    this.$settings.unsubscribe()
  }

  /**
   * setActiveDeck changes the settings, not affecting game datas
   * @param account use p2p deck
   * @param deckid
   */
  setActiveDeck(account: boolean, deckid: number) {
    let settings = this.conn.settings$.value
    settings.deck.account = account
    settings.deck.id = deckid
    this.conn.settings$.next(settings)
  }

  /**
   * startVsAiGame sends WS request to make a new game vs AI
   */
  startVsAiGame() {
    this.ws.send('/game/new', {
      ai: true,
      account: !!this.conn.settings$.value.deck.account,
      deckid: this.conn.settings$.value.deck.id
    })
  }

  /**
   * startVsHumanQueue sends WS request to make a new game vs another player
   */
  startVsHumanQueue() {
    // this.conn.sendWS('/game/new', {
    //   ai: false,
    //   account: !!this.conn.settings$.value.deck.account,
    //   deckid: this.conn.settings$.value.deck.id
    // })
  }

  // template display macros

  getopponent(username: string) : string {
    if (!this.game) return ""
    let opponent = ""
    for(let key of Object.keys(this.game.seats)) {
      if (key != username) opponent = key
    }
    return opponent
  }

  // updateTarget() {
  //   console.debug('updateTarget', this.game.state.data.helper, this.game.state.data.display)
  //   if (this.game.state.seat == this.myaccount.username) {
  //     let me = this
  //     let choices = []
  //     for (let i=0; i < this.game.state.data.choices.length; i++) {
  //       choices.push({
  //         display:this.game.state.data.display,
  //         f:function(val) {
  //           me.sendGCID(me.game, val)
  //           me.settings.game.overlay.clear()
  //           me.conn.settings$.next(me.settings)
  //         },
  //       })
  //     }
  //     this.settings.game.overlay = new Overlay("Target...", this.settings.game.overlay)
  //     this.settings.game.overlay.target = this.game.state.data.helper
  //     this.settings.game.overlay.title = this.game.state.data.display
  //     this.settings.game.overlay.choices = choices
  //     this.conn.settings$.next(this.settings)
  //   }
  // }

  powerString(power: Power): string {
    if (power.usesturn) return 'use'
    else if (power.useskill) return 'kill'
    else return "activate"
  }

  clickDefend(game: Game, id: string) {
    this.sendGCID(game, id)
    let overlay = this.games.overlay$.value
    this.games.overlay$.next(overlay.stack)
  }

  clickHand() {
    this.settings.game.hand.open=!this.settings.game.hand.open
    this.conn.settings$.next(this.settings)
  }

  clickHandCard(card: GameCard) {
    console.debug('click hand', card)
    this.settings.game.hand.open = false
    if (this.settings.game.hand.quickcast) {
      this.games.sendplay(this.game, card)
    } else {
      let overlay = new Overlay('Hand: ' + card.name, this.games.overlay$.value)
      overlay.card = card
      overlay.stack = this.overlay
      this.games.overlay$.next(overlay)
    }

    this.conn.settings$.next(this.settings)
  }

  clickTarget(target: any) {
    let overlay = this.games.overlay$.value
    let f = overlay.targetF
    if (f) f(target)
    this.games.overlay$.next(overlay.stack)
  }

  // clickEnd(game: Game) {
  //   this.sendPass(game)
  //   this.games.data$.next(null)
  //   this.router.navigateByUrl('/')
  // }

  // send data

  sendChat(game: Game, chatinput: any) {
    console.debug('send chat', chatinput)
    this.ws.send('/game', {
      gameid: game.id,
      uri: 'chat',
      text: chatinput.value,
    })
    chatinput.value = ''
  }

  sendPass(game: Game) {
    console.debug('send pass')
    this.ws.send('/game', {
      gameid: game.id,
      uri: 'pass',
      pass: game.state.id,
    })
  }

  zoomShowPower(token: GameToken, power: Power): boolean {
    if (power.trigger) return false
    else if (!power.usesturn) return true
    else return token.awake
  }

  sendChoice(game: Game, choice: string) {
    this.ws.send('/game', {
      gameid: game.id,
      uri: game.state.id,
      choice: choice,
    })
  }

  sendGCID(game: Game, id: string) {
    this.ws.send('/game', {
      gameid: game.id,
      uri: game.state.id,
      id: id,
    })
  }
}
