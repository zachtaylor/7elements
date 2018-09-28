import { Component, OnInit, Input } from '@angular/core'
import { Deck } from '../api'
import { ConnService } from '../conn.service'

@Component({
  selector: 'app-deck-summary',
  templateUrl: './deck-summary.component.html',
  styleUrls: ['./deck-summary.component.css']
})
export class DeckSummaryComponent implements OnInit {
  @Input() deck : Deck

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
  }

}
