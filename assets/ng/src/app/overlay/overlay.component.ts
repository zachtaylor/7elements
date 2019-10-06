import { Component, OnInit, Input } from '@angular/core'

@Component({
  selector: 'overlay',
  templateUrl: './overlay.component.html',
  styleUrls: ['./overlay.component.css']
})
export class OverlayComponent implements OnInit {
  @Input() title = ''

  constructor() { }

  ngOnInit() {

  }
}
