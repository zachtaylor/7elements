import { Component, OnInit } from '@angular/core'
import { UserService } from '../user.service'
import { PingService } from '../ping.service'

@Component({
  selector: 'app-myaccount',
  templateUrl: './myaccount.component.html',
  styleUrls: ['./myaccount.component.css']
})
export class MyAccountComponent implements OnInit {

  constructor(public pingService : PingService, public userService : UserService) {
  }

  ngOnInit() {
  }

}
