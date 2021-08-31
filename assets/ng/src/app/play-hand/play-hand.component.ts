import { Component, HostBinding, HostListener } from '@angular/core';
import { Subscription } from 'rxjs';
import { GlobalData, GameState, GameMenu } from '../api';
import { VII } from '../7.service';
import { SettingsService } from '../settings.service';

@Component({
  selector: 'vii-play-hand',
  templateUrl: './play-hand.component.html',
  styleUrls: ['./play-hand.component.css'],
  // changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PlayHandComponent {
  glob : GlobalData
  private $glob : Subscription

  hand : string[]
  private $hand : Subscription

  expand = false

  state : GameState
  private $state : Subscription

  @HostBinding('style.margin-bottom.px') private get getOffset() {
    return this.expand ? 0 : -319
  }

  @HostListener('mouseenter') onMouseEnter() {
    // console.debug('mouseenter')
  }

  @HostListener('mouseleave') onMouseLeave() {
    // console.debug('mouseleave')
  }

  constructor(public vii : VII, public settings : SettingsService) { }

  ngOnInit() {
    this.$glob = this.vii.global$.subscribe(glob => { this.glob = glob })
    this.$hand = this.vii.gamehand$.subscribe(hand => { this.onGameHand(hand) })
    this.$state = this.vii.gamestate$.subscribe(state => { this.state = state })
  }

  ngOnDestroy() {
    this.$glob.unsubscribe()
    this.$hand.unsubscribe()
    this.$state.unsubscribe()
  }

  onGameHand(hand) {
    console.debug('game hand', hand)
    this.hand = hand
  }
  // private copy(hand) : object {
  //   let obj = {}
  //   Object.keys(hand).forEach(key => {
  //     obj[key] = hand[key]
  //   })
  //   return obj
  // }

  clickToggle() {
    this.expand = !this.expand
  }
  
  clickCard(gcid : string) {
    console.debug('click hand', gcid)
    this.expand = false
    let gamecard = this.vii.card$(gcid).value
    let card = this.glob.cards[gamecard.cardid - 1]
    let overlay = new GameMenu('Hand: ' + card.name, this.vii.gamemenu$.value)
    overlay.card = gamecard
    overlay.stack = this.vii.gamemenu$.value
    this.vii.gamemenu$.next(overlay)
  }

  sendChat(chatinput) {
    console.debug('send chat', chatinput.value)
    this.vii.sendmessage('game', chatinput.value)
    chatinput.value = ''
  }

}
