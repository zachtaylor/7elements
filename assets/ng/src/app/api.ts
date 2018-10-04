export class Card {
  id : number
  name : string
  text : string
  image : string
  type : string
  costs : Map<number, number>
  body? : {
    attack : number
    health : number
  }
  powers : Map<number, Power>
}

export class Chat {
  name : string
  username : string
  messages : Array<ChatMessage>

  constructor() {
    this.messages = new Array<ChatMessage>();
  }
}

export class ChatMessage {
  channel : string
  username : string
  message : string
  time : string
}

export class Deck {
  id : number
  name : string
  username : string
  level : number
  wins : number
  color : string
  cards : Map<number, number>
}

export class Game {
  id : string
  hand : Array<Card>
  elements : Array<number>
  opponents : Array<string>
}

export class MyAccount {
  username : string
  email : string
  sessionlife: string
  coins : number
  cards : Map<number, number>
  decks : Array<Deck>
  games : Array<Game>
}

export class PingData {
  cards : Array<Card>
  decks : Array<Deck>
  online : number
}

export class Power {
  id : number
  text : string
  costs : Map<number, number>
  trigger : string
  usesturn : boolean
  useskill : boolean
  target : string
}
