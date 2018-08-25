import { Component, OnInit } from '@angular/core';
import { UserService } from '../user.service';
import { GlobalService } from '../global.service';

@Component({
  selector: 'app-myaccount',
  templateUrl: './myaccount.component.html',
  styleUrls: ['./myaccount.component.css']
})
export class MyAccountComponent implements OnInit {

  constructor(public globalService : GlobalService, public userService : UserService) {
  }

  ngOnInit() {
  }

}
