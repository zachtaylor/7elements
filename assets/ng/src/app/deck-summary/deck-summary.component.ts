import { Component, OnInit, Input, HostBinding } from '@angular/core'
import { Deck } from '../api'
import { Subscription } from 'rxjs'
import { SettingsService } from '../settings.service'

@Component({
  selector: 'app-deck-summary',
  templateUrl: './deck-summary.component.html',
  styleUrls: ['./deck-summary.component.css']
})
export class DeckSummaryComponent implements OnInit {
  @Input() deck : Deck
  @HostBinding('class.active') isActive : boolean = false
  private $queue : Subscription

  constructor(public settings : SettingsService) { }

  ngOnInit() {
    this.$queue = this.settings.queue$.subscribe(queue => {
      if (!this.deck) {
        return
      }
      this.isActive = (this.deck.username == queue.owner) && (this.deck.id == queue.id)
    })
  }

  ngOnDestroy() {
    this.$queue.unsubscribe()
  }
}
