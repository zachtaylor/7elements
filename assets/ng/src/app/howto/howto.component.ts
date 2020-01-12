import { Component, OnInit } from '@angular/core';
import { ConnService } from '../conn.service';

@Component({
  selector: 'app-howto',
  templateUrl: './howto.component.html',
  styleUrls: ['./howto.component.css']
})
export class HowtoComponent implements OnInit {

  constructor(public conn: ConnService) { }

  ngOnInit() {
  }

}
