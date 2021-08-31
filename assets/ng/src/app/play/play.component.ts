import { Component, OnInit } from '@angular/core'
import { VII } from '../7.service'
import { PingData, Game, GameCard, MyAccount, Power, GlobalData, GameToken, GameSeat, Card, GameMenu, QueueSetting, GameState } from '../api'
import { Subscription, interval } from 'rxjs'
import { SettingsService } from '../settings.service'

@Component({
  selector: 'app-play',
  templateUrl: './play.component.html',
  styleUrls: ['./play.component.css']
})
export class PlayComponent implements OnInit {

  glob: GlobalData
  private $glob: Subscription

  myaccount: MyAccount
  private $myaccount: Subscription

  state : GameState
  private $state: Subscription

  overlay: GameMenu
  private $overlay: Subscription
  
  timer: number
  private $ticker: Subscription

  constructor(public vii: VII, public settings : SettingsService) { }

  ngOnInit() {
    this.$glob = this.vii.global$.subscribe(glob => {
      this.glob = glob
    })
    this.$myaccount = this.vii.account$.subscribe(myaccount => {
      this.myaccount = myaccount
    })
    this.$state = this.vii.gamestate$.subscribe(state => {
      this.state = state
      if (state) this.timer = state.timer
    })
    this.$overlay = this.vii.gamemenu$.subscribe(overlay => {
      this.overlay = overlay
    })
    this.$ticker = interval(1000).subscribe(_ => {
      if (this.timer > 0) this.timer--
    })
  }

  ngOnDestroy() {
    console.debug('play ngOnDestroy')
    this.$glob.unsubscribe()
    this.$myaccount.unsubscribe()
    this.$state.unsubscribe()
    this.$overlay.unsubscribe()
    this.$ticker.unsubscribe()
  }

  // handMouseEnter is called when the mouse enter the hand div
  handMouseEnter() {
    // console.log('hand mouse enter')
  }

  // handMouseEnter is called when the mouse leaves the hand div
  handMouseLeave() {
    // console.log('hand mouse leave')
  }

  showChoiceViewer() : boolean {
    return !!this.overlay ||
      (this.state && this.state.name=='target' && this.state.seat==this.vii.user())
  }

  showStackViewer() : boolean {
    return this.state && 
      (this.state.name=='play'||this.state.name=='trigger') &&
      (!this.state.reacts[this.vii.user()])
  }

  // template display macros


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
  //     this.settings.game.overlay = new GameMenu("Target...", this.settings.game.overlay)
  //     this.settings.game.overlay.target = this.game.state.data.helper
  //     this.settings.game.overlay.title = this.game.state.data.display
  //     this.settings.game.overlay.choices = choices
  //     this.vii.settings$.next(this.settings)
  //   }
  // }

  powerString(power: Power): string {
    if (power.usesturn) return 'use'
    else if (power.useskill) return 'kill'
    else return "activate"
  }

  clickDefend(game: Game, id: string) {
    this.sendGCID(game, id)
    let overlay = this.vii.gamemenu$.value
    this.vii.gamemenu$.next(overlay.stack)
  }

  // clickHand() {
  //   this.play.hand=!this.play.hand
  //   this.settings.play$.next(this.play)
  // }

  clickTarget(target: any) {
    let overlay = this.vii.gamemenu$.value
    let f = overlay.targetF
    if (f) f(target)
    this.vii.gamemenu$.next(overlay.stack)
  }

  // clickEnd(game: Game) {
  //   this.sendPass(game)
  //   this.games.data$.next(null)
  //   this.router.navigateByUrl('/')
  // }

  // send data

  zoomShowPower(token: GameToken, power: Power): boolean {
    if (power.trigger) return false
    else if (!power.usesturn) return true
    else return token.awake
  }

  sendChoice(game: Game, choice: string) {
    this.vii.send('/game', {
      gameid: game.id,
      uri: game.state.id,
      choice: choice,
    })
  }

  sendGCID(game: Game, id: string) {
    this.vii.send('/game', {
      gameid: game.id,
      uri: game.state.id,
      id: id,
    })
  }
}
