export class ElementCount {
  items: Map<number, number>
  constructor(items? : Map<number, number>) {
    if (items) {
      this.items = items
    } else {
      this.items = new Map<number, number>()
    }
  }
  length() : number {
    let length = 0, keys = Object.keys(this.items)
    for (let i = 0; i < keys.length; i++) {
      length += this.items[keys[i]]
    }
    return length
  }
  test(costs: Map<number, number>) : boolean {
    let keys = Object.keys(costs)
    for (let i=0; i < keys.length; i++) {
      let el = keys[i]
      if (el=='0') {

      } else if (this.items[el] < costs[el]) {
        return false
      }
    }
    // guarantees costs where el==0
    let countcosts = new ElementCount(costs)
    return this.length() >= countcosts.length()
  }
}

export class Card {
  id: number
  name: string
  text: string
  image: string
  type: string
  costs: Map<number, number>
  body?: {
    attack: number
    health: number
  }
  powers: Array<Power>
  constructor(data: any) {
    if (!data) return
    if (data.id) this.id = data.id
    if (data.name) this.name = data.name
    if (data.text) this.text = data.text
    if (data.image) this.image = data.image
    if (data.type) this.type = data.type
    if (data.costs) this.costs = data.costs
    if (data.body) this.body = data.body
    if (data.powers) this.powers = data.powers
  }
  // enoughkarma(karma: Map<number, number>) : boolean {
  //   let karmacount = new ElementCount(karma)
  //   return karmacount.test(this.costs)
  // }
}

export class Chat {
  name: string
  username: string
  messages: Array<ChatMessage>

  constructor() {
    this.messages = new Array<ChatMessage>();
  }
}

export class ChatMessage {
  channel: string
  username: string
  message: string
  time: string
}

export class Deck {
  id: number
  name: string
  username: string
  level: number
  wins: number
  cover: string
  cards: Map<number, number>
}

export class Game {
  id: string
  username: string
  state: GameState
  hand: Map<string, GameCard>
  seats: Map<string, GameSeat>

  constructor() {
    this.seats = new Map<string, GameSeat>()
  }
}

export class GameState {
  id: string
  name: string
  data: any
  seat: string
  timer: number
  reacts: Map<string, string>
}

export class GameSeat {
  username: string
  elements: Map<number, Array<boolean>>
  present: Array<string>
  hand: number
  past: Array<string>
  future: number
  life: number
  color: string
}

export class GameCard {
  id: string
  cardid: number
  name: string
  text: string
  username: string
  image: string
  powers: Array<Power>
  body?: {
    attack: number
    health: number
  }
  type: string
}

export class GameToken {
  id: string
  awake: boolean
  cardid: number
  name: string
  username: string
  text: string
  powers: Array<Power>
  body?: {
    attack: number
    health: number
  }
  type: string
}

export class MyAccount {
  username: string
  email: string
  sessionlife: string
  coins: number
  cards: Map<number, number>
  decks: Array<Deck>
  gameid: number
}

export class Notification {
  level: string
  source: string
  message: string
}

export class PingData {
  ping: number
  online: number
}

export class GlobalData {
  cards: Array<Card>
  decks: Array<Deck>
  packs: Array<any>
  users: number
}

export class Power {
  id: number
  text: string
  costs: Map<number, number>
  trigger: string
  usesturn: boolean
  useskill: boolean
  target: string
}

export class Settings {
  deck: {
    id: number
    account: boolean
  }
  game: {
    hand: {
      open: boolean
      openChangeOnUpdate: boolean
      quickcast: boolean
    }
    chat: {
      off: boolean
      hideT: number
    }
  }
  constructor() {
    this.deck = {
      id:1,
      account:false,
    }
    this.game = {
      hand: {
        open:false,
        openChangeOnUpdate:true,
        quickcast:false,
      },
      chat: {
        off:false,
        hideT:7000,
      }
    }
  }
}

export class Overlay {
  title: string // text
  target: string  // targeting enum
  targetF: any // this func procs for 'target' selection
  token: GameToken
  card: GameCard
  choices: Array<OverlayChoice> // offered choices
  stack: Overlay
  important: boolean // hides the close button

  constructor(title: string, stack: Overlay) {
    this.title=title
    this.choices = []
    this.stack = stack
  }

  show() : boolean {
    if (this.title) {
      return true
    } else if (this.target) {
      return true
    } else if (this.token != null) {
      return true
    } else if (this.card != null) {
      return true
    } 
    return false
  }
}

export class OverlayChoice {
  display: string
  f: any
}
