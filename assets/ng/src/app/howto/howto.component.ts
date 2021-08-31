import { Component, OnInit } from '@angular/core';
import { VII } from '../7.service';

@Component({
  selector: 'app-howto',
  templateUrl: './howto.component.html',
  styleUrls: ['./howto.component.css']
})
export class HowtoComponent implements OnInit {

  constructor(public vii: VII) { }

  ngOnInit() {
  }

}
