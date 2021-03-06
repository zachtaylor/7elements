import { Component, OnInit } from '@angular/core'
import { ConnService } from '../conn.service'
import { WebsocketService } from '../websocket.service'

@Component({
  selector: 'app-chats',
  templateUrl: './chats.component.html',
  styleUrls: ['./chats.component.css']
})
export class ChatsComponent implements OnInit {
  channel = ''
  openChats = ['all']

  constructor(public conn: ConnService, public ws : WebsocketService) {
  }

  ngOnInit() {
  }

  clickNewChat() {
    let name = prompt('Please enter new chat name', 'all')
    if (!name) return;
    this.openChats.push(name)
    this.channel = name
    this.ws.send('/chat/join', {
      'channel':name
    })
  }

  clickChannel(name : string) {
    console.debug('chat channel: ', name)
    if (this.channel == name) {
      return
    }
    this.channel = name
    this.ws.send('/chat/join', {
      'channel':name
    })
  }

  sendChat(inputChat) {
    let msg = inputChat.value
    if (!msg) return;
    console.debug('chat send: ', msg)
    this.ws.send('/chat', {
      "channel":this.channel,
      "message":msg
    })
    inputChat.value = ''
  }
}
