import { Component, OnInit } from '@angular/core'
import { MyAccount, PingData } from '../api'
import { ConnService } from '../conn.service'
import { CookieService } from 'ngx-cookie-service'
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-games',
  templateUrl: './games.component.html',
  styleUrls: ['./games.component.css']
})
export class GamesComponent implements OnInit {
  deckID = 1
  deckP2P = false

  constructor(public conn : ConnService, private http : HttpClient, private cookieService : CookieService) {
  }

  ngOnInit() {
    let deckID = parseInt(this.cookieService.get('DeckID'))
    if (deckID) {
      this.deckID = deckID
    }
    this.deckP2P = parseInt(this.cookieService.get('DeckP2P')) > 0
  }

  getDeck(deckid : number, deckp2p : boolean, user : MyAccount, ping : PingData) {
    if (!ping) {
      return null
    } else if (!user) {
      return null
    } else if (deckp2p) {
      for (var i=0; i<user.decks.length; i++) {
        var deck = user.decks[i]
        if (deck && deck.id == deckid) {
          return deck
        }
      }
    } else {
      for (var i=0; i<ping.decks.length; i++) {
        var deck = ping.decks[i]
        if (deck && deck.id == deckid) {
          return deck
        }
      }
    }
    return null
  }

  changeDeckP2P() {
    this.deckP2P = !this.deckP2P
  }

  startVsAiGame() {
    // this.http.get<>('/api/newgame.json').subscribe(things => this.data$.next(things))
    console.log('start vs ai')
  }

  deckLeft() {
    if (this.deckID > 1) {
      this.deckID--;
    }
  }

  deckRight() {
    if (this.deckID < 3) {
      this.deckID++;
    }
  }

  startVsHumanQueue() {
    console.log('start vs human queue')
  }
}
