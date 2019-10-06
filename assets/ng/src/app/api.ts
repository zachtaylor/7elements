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
  cover : string
  cards : Map<number, number>
}

export class Game {
  id : string
  life : number
  hand : Array<GameCard>
  events : Array<String>
  state : GameState
  username : string
  opponent : string
  elements : Map<number, Array<boolean>>
  seats : Map<string, GameSeat>

  constructor() {
    this.elements = new Map<number, Array<boolean>>()
    this.seats = new Map<string, GameSeat>()
  }
}

export class GameState {
  id : string
  name : string
  data : any
  seat : string
  timer : number
  reacts : Map<string, string>
}

export class GameSeat {
  username : string
  elements : Map<number, Array<boolean>>
  active : Map<string, GameCard>
  hand : number
  past : Map<string, GameCard>
  future : number
  life : number
  color : string
}

export class GameCard {
  gcid : string
  cardid : number
  name : string
  text : string
  username : string
  image : string
  awake : boolean
  powers : Array<Power>
  body? : {
    attack : number
    health : number
  }
}

export class MyAccount {
  username : string
  email : string
  sessionlife: string
  coins : number
  cards : Map<number, number>
  decks : Array<Deck>
  gameid : number
}

export class Notification {
  level : string
  source : string
  message : string
}

export class PingData {
  online : number
}

export class GlobalData {
  cards : Array<Card>
  decks : Array<Deck>
  packs : Array<any>
  users : number
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

export class Settings {
  deck : {
    id : number
    account : boolean
  }
  game : {
    hand : boolean
    zoom : GameCard
    target : Target
  }
}

export class Target {
  display : string
  helper : string
  send : any
}