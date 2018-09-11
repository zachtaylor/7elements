import { Component, OnInit, Input } from '@angular/core'
import { Deck } from '../api'
import { PingService } from '../ping.service'

@Component({
  selector: 'app-deck-summary',
  templateUrl: './deck-summary.component.html',
  styleUrls: ['./deck-summary.component.css']
})
export class DeckSummaryComponent implements OnInit {
  @Input() deck : Deck

  constructor(public pingService : PingService) {
  }

  ngOnInit() {
  }

}
