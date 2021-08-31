import { Component, OnInit } from '@angular/core'
import { VII } from '../7.service'

@Component({
  selector: 'app-lost',
  templateUrl: './lost.component.html',
  styleUrls: ['./lost.component.css']
})
export class LostComponent implements OnInit {
  constructor(public vii : VII) { }

  ngOnInit() {
    if (this.vii.ready()) this.vii.goto('/')
  }

  restart() {
    this.vii.goto('/')
    setTimeout(() => { window.location.reload() })
  }
}
