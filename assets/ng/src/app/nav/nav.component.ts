import { Component, OnInit } from '@angular/core'
import { Subscription } from 'rxjs'
import { MyAccount } from '../api'
import { VII } from '../7.service'

@Component({
  selector: 'vii-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.css']
})
export class NavComponent implements OnInit {
  myaccount : MyAccount
  private $myaccount : Subscription

  constructor(private vii : VII) { }

  ngOnInit() {
    this.$myaccount = this.vii.account$.subscribe(myaccount => { this.myaccount = myaccount })
  }

  ngOnDestroy() {
    this.$myaccount.unsubscribe()
  }

  getAccountLinkName() : string {
    if (this.myaccount) return this.myaccount.username
    return 'My Account'
  }
}
