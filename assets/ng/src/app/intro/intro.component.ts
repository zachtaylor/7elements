import { Component, OnInit } from '@angular/core'
import { VII } from '../7.service'

@Component({
  selector: 'app-intro',
  templateUrl: './intro.component.html',
  styleUrls: ['./intro.component.css']
})
export class IntroComponent implements OnInit {

  constructor(public vii : VII) {
  }

  ngOnInit() {
  }

}
