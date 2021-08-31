import { Component, OnInit, Input, HostBinding } from '@angular/core'
import { GameToken, Game, GameState } from '../api'
import { Subscription } from 'rxjs'
import { VII } from '../7.service'

@Component({
  selector: 'app-game-token',
  templateUrl: './game-token.component.html',
  styleUrls: ['./game-token.component.css']
})
export class GameTokenComponent implements OnInit {
  @Input() tokenid: string

  state : GameState
  private $state: Subscription

  token : GameToken
  private $token : Subscription

  @HostBinding('class.triggered') get isTriggered(): boolean {
    let data = this.state.data
    return data && this.token && (
      (data.id == this.token.id) ||
      (data.token && data.token.id == this.token.id) ||
      (data.attack && data.attack.id == this.token.id) ||
      (data.block && data.block.id == this.token.id) ||
      (data.target && data.target.id == this.token.id)
    )
  }

  constructor(public vii : VII) {
  }

  ngOnInit() {
    this.$state = this.vii.gamestate$.subscribe(state => {
      this.state = state
    })
    this.$token = this.vii.token$(this.tokenid).subscribe(token => {
      this.token = token
    })
  }

  ngOnDestroy() {
    this.$state.unsubscribe()
    this.$token.unsubscribe()
  }

  getIcon(): string {
    let data = this.state.data
    if (!this.token) {
      return ''
    } else if (!data) {
      if (this.token.awake) {
        return 'awake.svg'
      } else {
        return 'asleep.svg'
      }
    } else if (this.state.name == 'combat' && (
      (data.attack.id == this.token.id) ||
      (data.block && data.block.id == this.token.id)
    )) {
      return 'combat.svg'
    } else if (this.state.name == 'attack' && data.id == this.token.id) {
      return 'attack.svg'
    } else if (data.token && data.token.id == this.token.id) {
      return 'star.svg'
    } else if (data.target == this.token.id) {
      return 'target.svg'
    } else if (this.token.awake) {
      return 'awake.svg'
    } else {
      return 'asleep.svg'
    }
  }
}
