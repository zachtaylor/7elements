import { Component, OnInit } from '@angular/core'
import { Deck, GlobalData, MyAccount } from '../api'
import { Subscription, Observable, BehaviorSubject } from 'rxjs'
import { ActivatedRoute } from '@angular/router'
import { VII } from '../7.service'
import { FormGroup, FormControl } from '@angular/forms'

@Component({
  selector: 'app-mydecks.id',
  templateUrl: './mydecks.id.component.html',
  styleUrls: ['./mydecks.id.component.css']
})
export class MyDecksIdComponent implements OnInit {
  deck: Deck
  dname = ''
  dcover = 0
  diff = new Map<number, number>()

  private id: number
  private $id: Subscription

  glob : GlobalData
  private $glob : Subscription

  myaccount: MyAccount
  private $myaccount: Subscription

  form = new FormGroup({
    name: new FormControl(''),
  })

  private $input: Subscription

  constructor(
    public vii: VII,
    private route: ActivatedRoute,
  ) { }

  ngOnInit() {
    let me = this
    this.$id = this.route.params.subscribe(params => {
      me.id = +params['id']
    })

    this.$myaccount = this.vii.account$.subscribe(myaccount => {
      console.debug('myaccount')
      me.myaccount = myaccount
      if (!myaccount) return;
      me.deck = myaccount.decks[me.id]
    })

    this.$glob = this.vii.global$.subscribe(glob => { this.glob = glob })

    this.$input = this.form.valueChanges.subscribe(() => {
      let val = this.form.get('name').value
      if (val == this.deck.name) {
        this.dname = ''
      } else {
        this.dname = val
      }
    })
  }

  ngOnDestroy() {
    this.$id.unsubscribe()
    this.$myaccount.unsubscribe()
    this.$glob.unsubscribe()
    this.$input.unsubscribe()
  }

  cover() : number {
    if (this.dcover) return this.dcover
    else return this.deck.cover
  }

  countLive(): number {
    return this.countDeck()+this.countDiff()
  }

  countDeck(): number {
    let num = 0
    Object.keys(this.deck.cards).forEach(v => {
      num += this.deck.cards[v]
    })
    return num
  }

  countDiff(): number {
    let num = 0
    let keys = Object.keys(this.diff)
    for (let i = 0; i < keys.length; i++) {
      num += this.diff[keys[i]]
    }
    return num
  }

  countDiffString(): string {
    let diff = this.countDiff()
    if (diff < 0) return ''+diff;
    else return '+'+diff;
  }

  count(id: number): number {
    return this.getDeck(id) + this.getDiff(id)
  }

  getDeck(id: number): number {
    if (this.deck.cards[id]) {
      return this.deck.cards[id]
    }
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
    this.vii.global$.value.cards.forEach(card => {
      let count = me.count(card.id)
      if (count > 0) view[card.id] = count
    })
    return view
  }

  clickUp(id: number) {
    if (this.count(id) + 1 > this.myaccount.cards[id]) return;
    this.diff[id]++
  }
  clickDown(id: number) {
    if (this.count(id) - 1 < 0) return;
    this.diff[id]--
  }

  clickCover(cardid : number) {
    if (this.deck.cover == cardid) {
      this.dcover = 0
    } else {
      this.dcover = +cardid
    }
  }

  clickSave() {
    let obj = {
      id: this.deck.id,
    }
    if (this.dname != this.deck.name) { obj["name"] = this.dname }
    if (this.countDiff() != 0) { obj["cards"] = this.diff }
    if (this.dcover > 0) { obj["cover"] = this.dcover }
    this.dname = ''
    this.diff = new Map<number, number>()
    this.dcover = 0
    this.vii.send('/deck', obj)
  }
}
