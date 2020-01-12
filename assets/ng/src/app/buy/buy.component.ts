import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'
import { MyAccount, PingData, GlobalData } from '../api'
import { Subscription } from 'rxjs'
import { WebsocketService } from '../websocket.service'

@Component({
  selector: 'app-buy',
  templateUrl: './buy.component.html',
  styleUrls: ['./buy.component.css']
})
export class BuyComponent implements OnInit {
  glob: GlobalData
  myaccount: MyAccount
  private $glob: Subscription
  private $myaccount: Subscription

  constructor(public conn: ConnService, public ws: WebsocketService) {
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
    this.ws.send('/packs/buy', {
      'packid': packid,
    })
  }
}
