import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { PingData } from './api'

@Injectable({
  providedIn: 'root'
})
export class GlobalService {
  data : PingData

  constructor(private http: HttpClient) {
    this.http.get<PingData>(`/api/ping.json`).subscribe(ping => {
      console.debug('global.service', ping);
      this.data = ping
    })
  }

  getCard(id : number) {
    if (!this.data) return null;
    for (var i = 0; i < this.data.cards.length; i++) {
      var card = this.data.cards[i];
      if (card.id == id) return card;
    }
    return null;
  }

  getDeck(id : number) {
    if (!this.data) return null;
    for (var i = 0; i < this.data.decks.length; i++) {
      var deck = this.data.decks[i];
      if (deck.id == id) return deck;
    }
    return null;
  }
}
