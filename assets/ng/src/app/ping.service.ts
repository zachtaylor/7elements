import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { PingData } from './api'
import { HttpClient } from '@angular/common/http'

@Injectable({
  providedIn: 'root'
})
export class PingService {
  data$ = new BehaviorSubject<PingData>(null);

  constructor(private http : HttpClient) {
    this.http.get<PingData>(`/api/ping.json`).subscribe(ping => {
      console.debug('ping.service', ping)
      this.data$.next(ping)
    })
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
