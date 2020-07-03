import { Component, OnInit, Input } from '@angular/core'
import { Game, GameCard, GameToken, Overlay } from '../api'
import { GameService } from '../game.service'
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

  constructor(private games: GameService) { }

  ngOnInit() {
    this.$ticker = interval(1000).subscribe(_ => {
      if (this.timer > 0) this.timer--
    })
    this.$game = this.games.data$.subscribe(game => {
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

  clickToken(tokenid: string) {
    let token = this.games.token$(tokenid).value
    let overlay = new Overlay('Token: ' + token.name, this.games.overlay$.value)
    overlay.token = token
    this.games.overlay$.next(overlay)
  }

  phaseIs(state: string): boolean {
    return this.game.state.name == state
  }

  phaseIsMy(state: string): boolean {
    return this.game.state.seat == this.username && this.phaseIs(state)
  }

  getTitleIcon(): string {
    if (this.game.state.data && this.game.state.data.target && this.game.state.data.target == this.username) {
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
    if (this.game.state.data && this.game.state.data.target && this.game.state.data.target == this.username) {
      return "Targeted"
    } else if (this.phaseIsMy('sunrise')) {
      return 'Sunrise'
    } else if (this.phaseIsMy('main')) {
      return 'Main'
    } else if (this.phaseIsMy('play')) {
      return 'Play: ' + this.game.state.data.card.name
    } else if (this.phaseIsMy('target')) {
      return 'Target: '
    } else if (this.phaseIsMy('trigger')) {
      return 'Trigger: ' + this.game.state.data.token.name
    } else if (this.phaseIsMy('attack')) {
      return 'Attack'
    } else if (this.phaseIsMy('combat')) {
      return 'Combat'
    } else if (this.phaseIsMy('sunset')) {
      return 'Sunset'
    } else {
    }
  }
}
