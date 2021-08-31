import { Component, OnDestroy, OnInit } from '@angular/core'
import { Subscription } from 'rxjs'
import { VII } from '../7.service'
import { MyAccount } from '../api'

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.css']
})
export class IndexComponent implements OnInit, OnDestroy {

  myaccount: MyAccount
  private $myaccount: Subscription

  constructor(public vii : VII) { }

  ngOnInit() {
    this.$myaccount = this.vii.account$.subscribe(myaccount => {
      this.myaccount = myaccount
    })
  }

  ngOnDestroy() {
    this.$myaccount.unsubscribe()
  }
}
