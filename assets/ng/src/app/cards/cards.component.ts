import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'
import { Card } from '../api'
import { Subscription } from 'rxjs'

@Component({
  selector: 'app-cards',
  templateUrl: './cards.component.html',
  styleUrls: ['./cards.component.css']
})
export class CardsComponent implements OnInit {
  cards : Array<Card>
  private $glob : Subscription

  constructor(public conn : ConnService) {
    
  }

  ngOnInit() {
    let me = this
    this.$glob = this.conn.global$.subscribe(glob => {
      if (!glob) return;
      me.cards = Array<Card>()
      for (let i=0; i < glob.cards.length; i++) {
        me.cards.push(glob.cards[i])
      }
      me.cards.sort((a : Card, b : Card) => {
        return a.id - b.id
      })
    })
  }

  ngOnDestroy() {
    this.$glob.unsubscribe()
  }
}
