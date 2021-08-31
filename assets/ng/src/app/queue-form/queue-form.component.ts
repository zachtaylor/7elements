import { Component, OnInit } from '@angular/core'
import { QueueSetting, GlobalData, MyAccount } from '../api'
import { Subscription } from 'rxjs'
import { VII } from '../7.service'
import { SettingsService } from '../settings.service'

@Component({
  selector: 'vii-queue-form',
  templateUrl: './queue-form.component.html',
  styleUrls: ['./queue-form.component.css']
})
export class QueueFormComponent implements OnInit {
  queue: QueueSetting
  private $queue: Subscription

  glob: GlobalData
  private $glob: Subscription

  myaccount: MyAccount
  private $myaccount: Subscription

  constructor(
    public vii: VII,
    public settings : SettingsService,
  ) { }

  ngOnInit() {
    this.$queue = this.settings.queue$.subscribe(queue => {
      this.queue = queue
    })
    this.$glob = this.vii.global$.subscribe(glob => {
      this.glob = glob
    })
    this.$myaccount = this.vii.account$.subscribe(myaccount => {
      this.myaccount = myaccount
    })
  }

  ngOnDestroy() {
    this.$queue.unsubscribe()
  }

  // clickPlay is called when on click "Play" for a new game
  clickPlay() {
    this.vii.send('/game/new', {
      deckid: this.queue.id,
      owner: this.queue.owner,
      custom: this.queue.custom,
      pvp: this.queue.pvp,
      hands: this.queue.hands,
      speed: this.queue.speed,
    })
  }

  onChangeHands(handsize: string) {
    this.queue.hands = handsize
    this.onChangeQueue()
  }

  onChangeSpeed(speed: string) {
    this.queue.speed = speed
    this.onChangeQueue()
  }

  onChangeQueue() {
    this.settings.queue$.next(this.queue)
  }
}
