import { Component, OnInit, Input } from '@angular/core'
import { Deck, PingData } from '../api'
import { PingService } from '../ping.service'

@Component({
  selector: 'app-deck-detail',
  templateUrl: './deck-detail.component.html',
  styleUrls: ['./deck-detail.component.css']
})
export class DeckDetailComponent implements OnInit {
  @Input() deck : Deck

  constructor(public pingService : PingService) {
  }

  ngOnInit() {
  }

}
