import { Component, OnInit } from '@angular/core'
import { Deck, MyAccount } from '../api'
import { Subscription, Observable, BehaviorSubject } from 'rxjs'
import { ActivatedRoute } from '@angular/router'
import { ConnService } from '../conn.service'
import { FormGroup, FormControl } from '@angular/forms'

@Component({
  selector: 'app-mydecks.id',
  templateUrl: './mydecks.id.component.html',
  styleUrls: ['./mydecks.id.component.css']
})
export class MyDecksIdComponent implements OnInit {
  deck: Deck
  myaccount: MyAccount
  diff = new Map<number, number>()
  private id: number
  private $id: Subscription
  private $myaccount: Subscription

  form = new FormGroup({
    name: new FormControl(''),
  })

  constructor(private route: ActivatedRoute, public conn: ConnService) {
  }

  ngOnInit() {
    let me = this
    this.$id = this.route.params.subscribe(params => {
      me.id = +params['id']
    })

    this.$myaccount = this.conn.myaccount$.subscribe(myaccount => {
      this.myaccount = myaccount
      if (!myaccount) return;
      me.deck = myaccount.decks[me.id]
    })
  }

  ngOnDestroy() {
    this.$id.unsubscribe()
    this.$myaccount.unsubscribe()
  }

  countLive(diff: Map<number, number>): number {
    return this.countDeck()+this.countDiff(diff)
  }

  countDeck(): number {
    let num = 0
    let me = this
    let keys = Object.keys(this.deck.cards)
    for (let i = 0; i < keys.length; i++) {
      let count = me.count(+keys[i])
      if (count) num += count
    }
    return num
  }

  
  countDiff(diff: Map<number, number>): number {
    let num = 0
    let keys = Object.keys(diff)
    for (let i = 0; i < keys.length; i++) {
      num += diff[keys[i]]
    }
    return num
  }

  count(id: number): number {
    return this.getDeck(id) + this.getDiff(id)
  }

  getDeck(id: number): number {
    if (this.deck[id]) {
      return this.deck[id]
    }
    this.deck[id] = 0
    return 0
  }

  getDiff(id: number): number {
    if (this.diff[id]) {
      return this.diff[id]
    }
    this.diff[id] = 0
    return 0
  }

  getView(): Map<number, number> {
    let me = this
    let view = new Map<number, number>()
    this.conn.global$.value.cards.forEach(card => {
      let count = me.count(card.id)
      if (count > 0) view[card.id] = count
    })
    return view
  }

  clickUpdate() {

  }

  clickUp(id: number) {
    if (this.count(id) + 1 > this.myaccount.cards[id]) return;
    this.diff[id]++
  }
  clickDown(id: number) {
    if (this.count(id) - 1 < 0) return;
    this.diff[id]--
  }

  onSubmit() {
    console.warn('submit')
  }
}
