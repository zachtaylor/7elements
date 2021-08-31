import { Component, OnInit, Input } from '@angular/core'
import { GameSeat, GameState, GameMenu } from '../api'
import { VII } from '../7.service'
import { Subscription, interval } from 'rxjs'

@Component({
  selector: 'app-game-seat',
  templateUrl: './game-seat.component.html',
  styleUrls: ['./game-seat.component.css']
})
export class GameSeatComponent implements OnInit {
  @Input() username: string

  timer: number
  private $ticker: Subscription

  seat: GameSeat
  private $seat: Subscription

  present: string[]
  private $present: Subscription

  state: GameState
  private $state: Subscription

  objectKeys = Object.keys

  constructor(private vii: VII) { }

  ngOnInit() {
    this.$ticker = interval(1000).subscribe(_ => {
      if (this.timer > 0) this.timer--
    })
    this.$seat = this.vii.seat$(this.username).subscribe(seat => {
      this.seat = seat
    })
    this.$present = this.vii.present$(this.username).subscribe(present => {
      this.present = present
    })
    this.$state = this.vii.gamestate$.subscribe(state => {
      if (!state) { return }
      this.state = state
      this.timer = state.timer
    })
  }

  ngOnDestroy() {
    console.debug('game-seat ngOnDestroy')
    this.$ticker.unsubscribe()
    this.$seat.unsubscribe()
    this.$present.unsubscribe()
    this.$state.unsubscribe()
  }

  clickToken(tokenid: string) {
    let token = this.vii.token$(tokenid).value
    let overlay = new GameMenu('Token: ' + token.name, this.vii.gamemenu$.value)
    overlay.token = token
    this.vii.gamemenu$.next(overlay)
  }

  phaseIs(state: string): boolean {
    return this.state.name == state
  }

  phaseIsMy(state: string): boolean {
    return this.state.seat == this.username && this.phaseIs(state)
  }

  getTitleIcon(): string {
    if (this.state.data && this.state.data.target && this.state.data.target == this.username) {
      return 'target'
    } else if (this.phaseIsMy('sunrise')) {
      return 'sunrise'
    } else if (this.phaseIsMy('main')) {
      return 'sun'
    } else if (this.phaseIsMy('play')) {
      return 'star'
    } else if (this.phaseIsMy('target')) {
      return 'magnify'
    } else if (this.phaseIsMy('trigger')) {
      return 'star'
    } else if (this.phaseIsMy('attack')) {
      return 'attack'
    } else if (this.phaseIsMy('combat')) {
      return 'combat'
    } else if (this.phaseIsMy('sunset')) {
      return 'sunset'
    } else {
    }
  }

  getTitleText(): string {
    if (this.state.data && this.state.data.target && this.state.data.target == this.username) {
      return "Targeted"
    } else if (this.phaseIsMy('sunrise')) {
      return 'Sunrise'
    } else if (this.phaseIsMy('main')) {
      return 'Main'
    } else if (this.phaseIsMy('play')) {
      return 'Play'
    } else if (this.phaseIsMy('target')) {
      return 'Target'
    } else if (this.phaseIsMy('trigger')) {
      return 'Trigger'
    } else if (this.phaseIsMy('attack')) {
      return 'Attack'
    } else if (this.phaseIsMy('combat')) {
      return 'Combat'
    } else if (this.phaseIsMy('sunset')) {
      return 'Sunset'
    }
  }
}
