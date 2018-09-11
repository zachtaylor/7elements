export type Card = {
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

export type Deck = {
  id : number
  name : string
  username : string
  level : number
  wins : number
  color : string
  cards : Map<number, number>
}

export type MyAccount = {
  username : string
  email : string
  sessionlife: string
  coins : number
  cards : Map<number, number>
  decks : Array<Deck>
  games : Array<Game>
}

export type PingData = {
  // version : number
  // update : string
  cards : Array<Card>
  decks : Array<Deck>
  online : number
}

export type Power = {
  id : number
  text : string
  costs : Map<number, number>
  trigger : string
  usesturn : boolean
  useskill : boolean
  target : string
}
