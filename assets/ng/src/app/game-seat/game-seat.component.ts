import { Component, OnInit, Input } from '@angular/core'
import { Game, GameCard } from '../api'
import { ConnService } from '../conn.service'
import { Subscription, interval } from 'rxjs'

@Component({
  selector: 'app-game-seat',
  templateUrl: './game-seat.component.html',
  styleUrls: ['./game-seat.component.css']
})
export class GameSeatComponent implements OnInit {
  game: Game
  @Input() username: string
  timer: number
  private $ticker: Subscription
  private $game: Subscription

  objectKeys = Object.keys

  constructor(private conn: ConnService) { }

  ngOnInit() {
    this.$ticker = interval(1000).subscribe(_ => {
      if (this.timer > 0) this.timer--
    })
    this.$game = this.conn.game$.subscribe(game => {
      if (!game) {
        return
      }
      this.game = game
      this.timer = game.state.timer
    })
  }

  ngOnDestroy() {
    console.debug('game-seat ngOnDestroy')
    this.$ticker.unsubscribe()
    this.$game.unsubscribe()
  }

  clickCard(card: GameCard) {
    let settings = this.conn.settings$.value
    settings.game.zoom = card
    this.conn.settings$.next(settings)
  }

  phaseIs(state: string): boolean {
    return this.game.state.name == state
  }

  phaseIsMy(state: string): boolean {
    return this.game.state.seat == this.username && this.phaseIs(state)
  }

  getIcon(): string {
    if (this.game.state.data && this.game.state.data.target && this.game.state.data.target == this.game.username) {
      return 'target'
    } else if (this.phaseIsMy('sunrise')) {
      return 'sunrise'
    } else if (this.phaseIsMy('main')) {
      return 'sun'
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
    } else if (this.game.state.reacts[this.username]) {
      return 'asleep'
    } else {
      return 'awake'
    }
  }
}
