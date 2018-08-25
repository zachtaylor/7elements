import { Component, OnInit } from '@angular/core';
import { GlobalService } from '../global.service';
import { UserService } from '../user.service';

@Component({
  selector: 'app-decks',
  templateUrl: './decks.component.html',
  styleUrls: ['./decks.component.css']
})
export class DecksComponent implements OnInit {

  constructor(public globalService : GlobalService, public userService : UserService ) { }

  ngOnInit() {
  }

}
