import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core'
import { GlobalData, GameState, GameMenu, Card, ElementCount, Power, GameSeat, KarmaPlanner } from '../api'
import { Subscription } from 'rxjs'
import { VII } from '../7.service'

@Component({
  selector: 'vii-play-choice-viewer',
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './play-choice-viewer.component.html',
  styleUrls: ['./play-choice-viewer.component.css']
})
export class PlayChoiceViewerComponent implements OnInit {
  seat: GameSeat
  private $seat : Subscription

  state: GameState
  private $state: Subscription

  overlay: GameMenu
  private $overlay: Subscription

  target: string

  constructor(public vii : VII) { }

  ngOnInit() {
    this.$seat = this.vii.seat$(this.vii.user()).subscribe(seat => { this.seat = seat })
    this.$state = this.vii.gamestate$.subscribe(state => { this.state = state })
    this.$overlay = this.vii.gamemenu$.subscribe(overlay => { this.overlay = overlay })
  }

  ngOnDestroy() {
    this.$seat.unsubscribe()
    this.$state.unsubscribe()
    this.$overlay.unsubscribe()
  }

  close() {
    this.vii.gamemenu$.next(this.overlay.stack)
  }

  testkarma(card: Card) : boolean {
    let karmacount = new ElementCount()
    let seat$ = this.vii.seat$(this.vii.user())
    let seat = seat$.value
    let keys = Object.keys(seat.elements)
    for (let i = 0; i < keys.length; i++) {
      let el = +keys[i]
      let set = seat.elements[el]
      for (let j = 0; j < set.length; j++) {
        if (set[j]) {
          if (karmacount.items[el]) {
            karmacount.items.set(el, karmacount.items.get(el) + 1)
          } else {
            karmacount.items.set(el, 1)
          }
        }
      }
    }
    return karmacount.test(card.costs)
  }

  testplayspeed() {
    let card = this.vii.global$.value.cards[this.overlay.card.cardid-1]
    return card.type=='spell'||(this.state.name=='main'&&this.state.seat==this.vii.user())
  }

  isPlayerTargetOption() {
    if (!this.overlay.target) return false
    else if (this.overlay.target=='player') return true
    else if (this.overlay.target=='player-being') return true
    else if (this.overlay.target=='player-other-being') return true
  }

  testtarget(arg : any) {
    console.debug('test target', this.overlay.target, arg)
    if (!this.overlay.target) return false
    else if (this.overlay.target=='player' && arg.username) return true
    else if (this.overlay.target=='player-being' && (arg.username || arg.body)) return true
    else if (this.overlay.target=='being' && arg.body) return true
    else if (this.overlay.target=='being-item' && arg.id) return true
    else if (this.overlay.target=='my-being' && arg.body && arg.owner==this.vii.user()) return true
    else if (this.overlay.target=='player-other-being' && arg.username) return true
  }

  getpaychoices(costs : any) : Array<any> {
    if (!costs[0]) {
      return [costs]
    }

    let active = {}
    Object.keys(this.seat.elements).forEach(key => {
      let element = +key, stack = this.seat.elements[element]
      active[element] = 0
      stack.forEach(ok => {
        if (ok) { active[element]++ }
      })
    })

    console.debug("pay choices active", active)
    let plan = {}, free = {}
    Object.keys(active).forEach(el => {
      let element = +el, count = active[element], cost = 0
      if (costs[element] > 0) cost = costs[element]
      if (cost > 0) plan[element] = cost
      free[element] = count - cost
    })

    console.debug("pay choice initial plan", plan, "free", free)

    let stage = new Array<KarmaPlanner>()
    stage.push(new KarmaPlanner(plan, free))
    for (let i=0; i < costs[0]; i++) {
      console.debug("play choice begin multiply")
      let newstage = new Array<KarmaPlanner>()
      stage.forEach(kp => {
        console.debug("play choice begin multiply use", kp)
        kp.multiply().forEach(newkp => {
          console.debug("play choice multiplied", kp, newkp)
          newstage.push(newkp)
        })
      })
      stage = newstage
    }

    let choices = []
    stage.forEach(kp => {
      choices.push(kp.plan)
    })

    console.log('finished pay choices', choices)

    return choices
  }

  sendplay(pay : object) {
    let card = this.overlay.card
    this.close()
    this.vii.sendplay(card, pay)
  }

  sendpower(power: Power) {
    let token = this.overlay.token
    this.close()
    this.vii.sendtrigger(token, power, power.costs)
  }

  sendtarget(target: string) {
    let f = this.overlay.targetF
    this.close()
    f(target)
  }

  // attacks

  showattack(): boolean {
    return this.overlay.token.awake&&
      this.overlay.token.body&&
      this.overlay.token.user==this.vii.user()&&
      this.state.name=='main'&&
      this.state.seat==this.vii.user()
  }

  sendattack(id: string) {
    this.vii.send('/game', {
      uri: 'attack',
      id: id,
    })
    this.close()
  }

  // defends

  showdefend(): boolean {
    return this.overlay.token.awake&&
      this.overlay.token.body&&
      this.overlay.token.user==this.vii.user()&&
      this.state.name=='attack'&&
      this.state.seat!=this.vii.user()
  }

  senddefend(id: string) {
    this.vii.send('/game', {
      uri: this.state.id,
      id: id,
    })
    this.close()
  }

  // choices

  sendchoice(choice: string) {
    this.vii.send('/game', {
      uri: this.state.id,
      choice: choice,
    })
    this.close()
  }
}
