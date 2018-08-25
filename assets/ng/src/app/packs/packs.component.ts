import { Component, OnInit } from '@angular/core'
import { UserService } from '../user.service'
import { GlobalService } from '../global.service';

@Component({
  selector: 'app-packs',
  templateUrl: './packs.component.html',
  styleUrls: ['./packs.component.css']
})
export class PacksComponent implements OnInit {

  constructor(public globalService : GlobalService, public userService : UserService) {
  }

  ngOnInit() {
  }

}
