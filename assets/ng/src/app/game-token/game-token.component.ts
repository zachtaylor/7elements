import { Component, OnInit, Input, HostBinding } from '@angular/core'
import { GameToken, Game } from '../api'
import { Subscription } from 'rxjs'
import { GameService } from '../game.service'

@Component({
  selector: 'app-game-token',
  templateUrl: './game-token.component.html',
  styleUrls: ['./game-token.component.css']
})
export class GameTokenComponent implements OnInit {
  @Input() tokenid: string
  token : GameToken
  $token : Subscription
  @Input() game: Game

  @HostBinding('class.triggered') get isTriggered(): boolean {
    let data = this.game.state.data
    return data && this.token && (data.id == this.token.id ||
      (data.token && data.token.id == this.token.id) ||
      (data.attack && data.attack.id == this.token.id) ||
      (data.block && data.block.id == this.token.id) ||
      (data.target && data.target.id == this.token.id))
  }

  constructor(public games : GameService) {
  }

  ngOnInit() {
    let me = this
    this.$token = this.games.token$(this.tokenid).subscribe(token => {
      me.token = token
    })
  }

  ngOnDestroy() {
    this.$token.unsubscribe()
  }

  getIcon(): string {
    let data = this.game.state.data
    if (!this.token) {
      return ''
    } else if (!data) {
      if (this.token.awake) {
        return 'awake.svg'
      } else {
        return 'asleep.svg'
      }
    } else if (this.game.state.name == 'combat' && (
      (data.attack.id == this.token.id) ||
      (data.block && data.block.id == this.token.id)
    )) {
      return 'combat.svg'
    } else if (this.game.state.name == 'attack' && data.id == this.token.id) {
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
