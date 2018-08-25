import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class MessageService {
  messages: string[] = [];
 
  add(message : string) {
    this.messages.push(message);
  }
  clear() {
    this.messages = [];
  }
  remove(i : number) {
    console.warn("message service remove", i, this.messages[i])
  }
}
