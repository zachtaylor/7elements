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
  color : string
  cards : Map<number, number>
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
