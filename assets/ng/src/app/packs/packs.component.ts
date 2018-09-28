import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'

@Component({
  selector: 'app-packs',
  templateUrl: './packs.component.html',
  styleUrls: ['./packs.component.css']
})
export class PacksComponent implements OnInit {

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
  }

}
