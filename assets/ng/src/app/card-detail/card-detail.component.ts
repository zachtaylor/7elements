import { Component, OnInit, Input, HostBinding } from '@angular/core'
import { Card } from '../api'

@Component({
  selector: 'app-card-detail',
  templateUrl: './card-detail.component.html',
  styleUrls: ['./card-detail.component.css']
})
export class CardDetailComponent implements OnInit {
  @Input() card : Card
  
  // change background color
  @HostBinding('class.white') private get isWhite() {
    return this.card && this.card.costs[1] > 0
  }
  @HostBinding('class.red') private get isRed() {
      return this.card && this.card.costs[2] > 0
  }
  @HostBinding('class.yellow') private get isYellow() {
      return this.card && this.card.costs[3] > 0
  }
  @HostBinding('class.green') private get isGreen() {
      return this.card && this.card.costs[4] > 0
  }
  @HostBinding('class.blue') private get isBlue() {
      return this.card && this.card.costs[5] > 0
  }
  @HostBinding('class.violet') private get isViolet() {
      return this.card && this.card.costs[6] > 0
  }
  @HostBinding('class.black') private get isBlack() {
      return this.card && this.card.costs[7] > 0
  }

  constructor() {
  }

  ngOnInit() {
  }

}
