import { Component, OnInit } from '@angular/core'
import { PingService } from '../ping.service'

@Component({
  selector: 'app-intro',
  templateUrl: './intro.component.html',
  styleUrls: ['./intro.component.css']
})
export class IntroComponent implements OnInit {

  constructor(public pingService : PingService) {
  }

  ngOnInit() {
  }

}
