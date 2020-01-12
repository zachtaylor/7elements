import { Component, OnInit } from '@angular/core'
import { WebsocketService } from '../websocket.service'

@Component({
  selector: 'app-lost',
  templateUrl: './lost.component.html',
  styleUrls: ['./lost.component.css']
})
export class LostComponent implements OnInit {
  constructor(public ws : WebsocketService) { }

  ngOnInit() {
    if (this.ws.ready()) this.ws.redirect('/')
  }

  restart() {
    this.ws.init()
    this.ws.redirect('/')
  }
}
