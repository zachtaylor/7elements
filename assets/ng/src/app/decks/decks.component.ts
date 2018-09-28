import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'

@Component({
  selector: 'app-decks',
  templateUrl: './decks.component.html',
  styleUrls: ['./decks.component.css']
})
export class DecksComponent implements OnInit {

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
  }

}
