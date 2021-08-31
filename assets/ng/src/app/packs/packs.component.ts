import { Component, OnInit } from '@angular/core';
import { GlobalData, MyAccount } from '../api';
import { Subscription } from 'rxjs';
import { VII } from '../7.service';

@Component({
  selector: 'vii-packs',
  templateUrl: './packs.component.html',
  styleUrls: ['./packs.component.css']
})
export class PacksComponent implements OnInit {
  glob: GlobalData
  private $glob: Subscription

  myaccount: MyAccount
  private $myaccount: Subscription

  constructor(public vii: VII) { }

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

  sendBuyPack(packid: number) {
    this.vii.send('/packs/buy', {
      'packid': packid,
    })
  }
}
