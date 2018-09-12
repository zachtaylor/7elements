import { Component, OnInit } from '@angular/core'
import { Deck, MyAccount, PingData } from '../api'
import { UserService } from '../user.service'
import { PingService } from '../ping.service'
import { CookieService } from 'ngx-cookie-service'
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-games',
  templateUrl: './games.component.html',
  styleUrls: ['./games.component.css']
})
export class GamesComponent implements OnInit {
  deckID = 0
  deckP2P = true

  constructor(public userService : UserService, public pingService : PingService, private http : HttpClient, private cookieService : CookieService) {
  }

  ngOnInit() {
    this.deckID = parseInt(this.cookieService.get('DeckID'))
    this.deckP2P = parseInt(this.cookieService.get('DeckP2P')) > 0
  }

  getDeck(deckid : number, deckp2p : boolean, user : MyAccount, ping : PingData) {
    if (user) {
      if (deckp2p) {
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
    } else if (user && ping) {
      return ping.decks[0]
    }
    return null
  }

  startVsAiGame() {
    // this.http.get<>('/api/newgame.json').subscribe(things => this.data$.next(things))
    console.log('start vs ai')
  }

  deckLeft(event : any) {
    console.log('deck left', event)
  }

  deckRight(event : any) {
    console.log('deck right', event)
  }

  startVsHumanQueue() {
    console.log('start vs human queue')
  }
}
