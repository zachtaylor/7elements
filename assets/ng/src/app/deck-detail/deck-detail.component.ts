import { Component, OnInit, Input } from '@angular/core'
import { Deck } from '../api'
import { ConnService } from '../conn.service'

@Component({
  selector: 'app-deck-detail',
  templateUrl: './deck-detail.component.html',
  styleUrls: ['./deck-detail.component.css']
})
export class DeckDetailComponent implements OnInit {
  @Input() deck : Deck

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
  }

}
