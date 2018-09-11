import { Component, OnInit } from '@angular/core'
import { PingService } from '../ping.service'

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.css']
})
export class IndexComponent implements OnInit {

  constructor(public pingService : PingService) {
  }

  ngOnInit() {
  }

}
