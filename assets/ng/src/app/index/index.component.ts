import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.css']
})
export class IndexComponent implements OnInit {

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
  }

}
