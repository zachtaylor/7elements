import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'

@Component({
  selector: 'app-myaccount',
  templateUrl: './myaccount.component.html',
  styleUrls: ['./myaccount.component.css']
})
export class MyAccountComponent implements OnInit {

  constructor(public conn : ConnService) {
  }

  ngOnInit() {
  }

}
