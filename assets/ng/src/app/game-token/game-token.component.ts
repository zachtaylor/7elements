import { Component, OnInit, Input, HostBinding } from '@angular/core'
import { GameCard, Game } from '../api'

@Component({
  selector: 'app-game-token',
  templateUrl: './game-token.component.html',
  styleUrls: ['./game-token.component.css']
})
export class GameTokenComponent implements OnInit {
  @Input() card : GameCard
  @Input() game : Game

  @HostBinding('class.triggered') get isTriggered() : boolean {
    let dat = this.game.state.data
    return dat && (dat.gcid == this.card.gcid ||
      (dat.card && dat.card.gcid == this.card.gcid) ||
      (dat.attack && dat.attack.gcid == this.card.gcid) ||
      (dat.block && dat.block.gcid == this.card.gcid) ||
      (dat.target && dat.target.gcid == this.card.gcid))
  }

  constructor() {
  }

  ngOnInit() {
  }

  getIcon() : string {
    let dat = this.game.state.data
    if (!dat) {
      if (this.card.awake) {
        return 'awake.svg'
      } else {
        return 'asleep.svg'
      }
    } else if (this.game.state.name == 'combat' && (
      (dat.attack.gcid == this.card.gcid) ||
      (dat.block && dat.block.gcid == this.card.gcid)
    )) {
      return 'combat.svg'
    } else if (this.game.state.name == 'attack' && dat.gcid == this.card.gcid) {
      return 'attack.svg'
    } else if (dat.card && dat.card.gcid == this.card.gcid) {
      return 'star.svg'
    } else if (dat.target == this.card.gcid) {
      return 'target.svg'
    } else if (this.card.awake) {
      return 'awake.svg'
    } else {
      return 'asleep.svg'
    }
  }
}
