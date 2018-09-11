import { Component, OnInit } from '@angular/core'
import { UserService } from '../user.service'
import { PingService } from '../ping.service'

@Component({
  selector: 'app-packs',
  templateUrl: './packs.component.html',
  styleUrls: ['./packs.component.css']
})
export class PacksComponent implements OnInit {

  constructor(public pingService : PingService, public userService : UserService) {
  }

  ngOnInit() {
  }

}
