import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { PingData, MyAccount } from './api'
import { HttpClient } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'

@Injectable({
  providedIn: 'root'
})
export class ConnService {
  ping$ = new BehaviorSubject<PingData>(null)
  myaccount$ = new BehaviorSubject<MyAccount>(null)
  notifications$ = new BehaviorSubject<Array<Notification>>(new Array<Notification>())

  constructor(private http : HttpClient, private cookieService : CookieService) {
    this.http.get<PingData>(`/api/ping.json`).subscribe(ping => {
      console.debug('conn.ping', ping)
      this.ping$.next(ping)
    })
    if (cookieService.get(`SessionID`)) {
      this.http.get<MyAccount>(`/api/myaccount.json`).subscribe(account => {
        console.debug('conn.myaccount', account)
        this.myaccount$.next(account)
        if (!account) this.cookieService.delete(`SessionID`)
      });
    }
  }

  isLoggedOut() {
    return !this.myaccount$.value
  }

  getCard(ping : PingData, cardid : number) {
    let cards = ping.cards
    for (let i=0; i<cards.length; i++) {
      if (cards[i].id == cardid) {
        return cards[i]
      }
    }
    return null
  }

  getDeck(ping : PingData, deckid : number) {
    let decks = ping.decks
    for (let i=0; i<decks.length; i++) {
      if (decks[i].id == deckid) {
        return decks[i]
      }
    }
    return null
  }
}
