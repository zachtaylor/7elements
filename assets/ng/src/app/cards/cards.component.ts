import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'

@Component({
  selector: 'app-cards',
  templateUrl: './cards.component.html',
  styleUrls: ['./cards.component.css']
})
export class CardsComponent implements OnInit {

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
  }

}
