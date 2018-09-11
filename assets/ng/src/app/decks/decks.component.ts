import { Component, OnInit } from '@angular/core'
import { PingService } from '../ping.service'
import { UserService } from '../user.service'

@Component({
  selector: 'app-decks',
  templateUrl: './decks.component.html',
  styleUrls: ['./decks.component.css']
})
export class DecksComponent implements OnInit {

  constructor(public pingService : PingService, public userService : UserService ) {

  }

  ngOnInit() {
  }

}
