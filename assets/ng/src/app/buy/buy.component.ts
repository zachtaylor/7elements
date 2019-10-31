import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'
import { MyAccount, PingData, GlobalData } from '../api'
import { Subscription } from 'rxjs'

@Component({
  selector: 'app-packs',
  templateUrl: './buy.component.html',
  styleUrls: ['./buy.component.css']
})
export class BuyComponent implements OnInit {
  glob: GlobalData
  myaccount: MyAccount
  private $glob: Subscription
  private $myaccount: Subscription

  constructor(public conn: ConnService) {
  }

  ngOnInit() {
    this.$glob = this.conn.global$.subscribe(glob => {
      this.glob = glob
    })
    this.$myaccount = this.conn.myaccount$.subscribe(myaccount => {
      this.myaccount = myaccount
    })
  }

  ngOnDestroy() {
    this.$glob.unsubscribe()
    this.$myaccount.unsubscribe()
  }

  sendBuyPack(packid: number) {
    this.conn.sendWS('/packs/buy', {
      'packid': packid,
    })
  }
}
