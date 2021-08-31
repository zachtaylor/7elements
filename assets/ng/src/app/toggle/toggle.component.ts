import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core'

@Component({
  selector: 'toggle',
  templateUrl: './toggle.component.html',
  styleUrls: ['./toggle.component.css']
})
export class ToggleComponent implements OnInit {
  @Input() check = false
  @Output() checkChange = new EventEmitter<boolean>()

  constructor() {
  }

  ngOnInit() {
  }

  toggle() {
    this.check = !this.check
    this.checkChange.emit(this.check)
  }
}
