import { Component, OnInit, Input, HostBinding } from '@angular/core'
import { Deck } from '../api'
import { ConnService } from '../conn.service'
import { Subscription } from 'rxjs'

@Component({
  selector: 'app-deck-summary',
  templateUrl: './deck-summary.component.html',
  styleUrls: ['./deck-summary.component.css']
})
export class DeckSummaryComponent implements OnInit {
  @Input() deck : Deck
  @HostBinding('class.active') isActive : boolean = false
  private $settings : Subscription

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
    this.$settings = this.conn.settings$.subscribe(settings => {
      if (!this.deck) {
        return
      }
      this.isActive = (!!this.deck.username == settings.deck.account) && (this.deck.id == settings.deck.id)
    })
  }

  ngOnDestroy() {
    this.$settings.unsubscribe()
  }
}
