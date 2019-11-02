import { Component, OnInit } from '@angular/core'
import { ActivatedRoute, Router } from '@angular/router'
import { ConnService } from '../conn.service'
import { PingData, Game, GameCard, MyAccount, Power, Settings, Target, GlobalData } from '../api'
import { Subscription, Observable, interval } from 'rxjs'

@Component({
  selector: 'app-play',
  templateUrl: './play.component.html',
  styleUrls: ['./play.component.css']
})
export class PlayComponent implements OnInit {
  glob: GlobalData
  myaccount: MyAccount
  game: Game
  timer: number
  settings: Settings
  private $ticker: Subscription
  private $glob: Subscription
  private $myaccount: Subscription
  private $game: Subscription
  private $settings: Subscription

  constructor(public conn: ConnService, private router: Router) {
  }

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
    this.$game = this.conn.game$.subscribe(game => {
      if (!game) {
        return
      }
      this.game = game
      this.timer = game.state.timer
      if (game.state.name == 'target') {
        this.updateTarget()
      }
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
    this.conn.sendWS('/game/new', {
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

  // template conditional display macros

  showStartOverlay() {
    let username = this.conn.getUsername()
    return this.game && username && this.game.state &&
      this.game.state.name == "start" && !this.game.state.reacts[username]
  }

  showSunriseOverlay() {
    let username = this.conn.getUsername()
    return this.game && username && this.game.state &&
      this.game.state.name == "sunrise" && this.game.state.seat == username
  }

  showChoiceOverlay() {
    let username = this.conn.getUsername()
    return this.game && this.game.state.name == 'choice' && this.game.state.seat == username
  }

  showEndOverlay() {
    return this.game && this.game.state && this.game.state.name == "end"
  }

  showPlayOverlay() {
    return this.game && this.game.state && this.game.state.name == "play"
  }

  updateTarget() {
    console.debug('updateTarget', this.game.state.data.helper, this.game.state.data.display)
    if (this.game.state.seat == this.game.username) {
      let me = this
      this.settings.game.target.helper = this.game.state.data.helper
      this.settings.game.target.display = this.game.state.data.display
      this.settings.game.target.send = function (val: any) {
        me.sendGCID(me.game, val)
        me.settings.game.target = null
      }
    }
  }

  powerString(power: Power): string {
    if (power.usesturn) return 'use'
    else if (power.useskill) return 'kill'
    else return "activate"
  }

  clickAttack(game: Game, gcid: string) {
    this.conn.sendWS('/game', {
      gameid: game.id,
      uri: 'attack',
      gcid: gcid,
    })
    this.settings.game.zoom = null
    this.conn.settings$.next(this.settings)
  }

  clickDefend(game: Game, gcid: string) {
    this.sendGCID(game, gcid)
    this.settings.game.zoom = null
    this.conn.settings$.next(this.settings)
  }

  clickHand(game: Game, card: GameCard) {
    console.debug('hand click', card)
    this.settings.game.hand = false
    if (card.type == 'spell' && card.powers.length > 0 && card.powers[0].target != 'self') {
      let conn = this.conn
      this.settings.game.target = {
        display: card.powers[0].text,
        helper: card.powers[0].target,
        send: function (val: any) {
          conn.sendWS('/game', {
            'uri': 'play',
            'gameid': game.id,
            'gcid': card.gcid,
            'target': val
          })
        }
      }
    } else {
      this.conn.sendWS('/game', {
        'uri': 'play',
        'gameid': game.id,
        'gcid': card.gcid
      })
    }
    this.conn.settings$.next(this.settings)
  }

  clickPower(game: Game, card: GameCard, power: Power) {
    console.debug('click power', card, power)
    if (power.target != '' && power.target != 'self') {
      let me = this
      this.settings.game.target = new Target()
      this.settings.game.target.helper = power.target
      this.settings.game.target.send = function (target: any) {
        me.sendPower(game, card.gcid, power.id, target)
      }
    } else {
      this.sendPower(game, card.gcid, power.id, '')
    }
    this.settings.game.zoom = null
    this.conn.settings$.next(this.settings)
  }

  clickTarget(game: Game, target: any) {
    this.settings.game.target.send(target)
    this.settings.game.target.helper = ''
    this.settings.game.target.send = false
    this.conn.settings$.next(this.settings)
  }

  clickEnd(game: Game) {
    this.sendPass(game)
    this.conn.game$.next(null)
    this.router.navigateByUrl('/')
  }

  // send data

  sendChat(game: Game, chatinput: any) {
    console.debug('send chat', chatinput)
    this.conn.sendWS('/game', {
      gameid: game.id,
      uri: 'chat',
      text: chatinput.value,
    })
    chatinput.value = ''
  }

  sendPass(game: Game) {
    console.debug('send pass')
    this.conn.sendWS('/game', {
      gameid: game.id,
      uri: 'pass',
      pass: game.state.id,
    })
  }

  zoomShowPower(card: GameCard, power: Power): boolean {
    if (power.trigger) return false
    else if (!power.usesturn) return true
    else return card.awake
  }

  sendChoice(game: Game, choice: string) {
    this.conn.sendWS('/game', {
      gameid: game.id,
      uri: game.state.id,
      choice: choice,
    })
  }

  sendSunriseElement(game: Game, i: number) {
    this.conn.sendWS('/game', {
      gameid: game.id,
      uri: game.state.id,
      elementid: i,
    })
  }

  sendPower(game: Game, gcid: string, powerid: number, target: any) {
    this.conn.sendWS('/game', {
      gameid: game.id,
      uri: 'trigger',
      gcid: gcid,
      powerid: powerid,
      target: target
    })
  }

  sendGCID(game: Game, gcid: string) {
    this.conn.sendWS('/game', {
      gameid: game.id,
      uri: game.state.id,
      gcid: gcid,
    })
  }
}
