import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core'

@Component({
  selector: 'toggle',
  templateUrl: './toggle.component.html',
  styleUrls: ['./toggle.component.css']
})
export class ToggleComponent implements OnInit {
  @Output() check = new EventEmitter()
  isChecked : boolean

  constructor() {
  }

  ngOnInit() {
  }

  change(isChecked : boolean) {
    this.check.next(isChecked)
  }
}
