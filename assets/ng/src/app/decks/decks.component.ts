import { Component, OnInit } from '@angular/core';
import { QueueSetting, MyAccount, GlobalData } from '../api';
import { Subscription } from 'rxjs';
import { VII } from '../7.service';
import { SettingsService } from '../settings.service';

@Component({
  selector: 'vii-decks',
  templateUrl: './decks.component.html',
  styleUrls: ['./decks.component.css']
})
export class DecksComponent implements OnInit {
  queue : QueueSetting
  glob : GlobalData
  myaccount : MyAccount
  private $queue: Subscription
  private $glob: Subscription
  private $myaccount: Subscription

  constructor(
    public vii : VII,
    public settings: SettingsService
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
    this.$glob.unsubscribe()
    this.$myaccount.unsubscribe()
  }

  setActiveDeck(owner: string, deckid: number) {
    this.queue.owner = owner
    this.queue.id = deckid
    this.settings.queue$.next(this.queue)
  }

}
