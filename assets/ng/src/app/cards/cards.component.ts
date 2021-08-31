import { Component, OnInit } from '@angular/core'
import { VII } from '../7.service'
import { Card, MyAccount, GlobalData } from '../api'
import { Subscription } from 'rxjs'

@Component({
  selector: 'app-cards',
  templateUrl: './cards.component.html',
  styleUrls: ['./cards.component.css']
})
export class CardsComponent implements OnInit {
  glob : GlobalData
  private $glob : Subscription

  myaccount : MyAccount
  private $myaccount : Subscription

  constructor(public vii : VII) { }

  ngOnInit() {
    this.$glob = this.vii.global$.subscribe(glob => {
      this.glob = glob
    })
    this.$myaccount = this.vii.account$.subscribe(myaccount => {
      this.myaccount = myaccount
    })
  }

  ngOnDestroy() {
    this.$glob.unsubscribe()
    this.$myaccount.unsubscribe()
  }
  
  count() : number {
    if (!this.glob) return 0
    return this.glob.cards.length
  }
}
