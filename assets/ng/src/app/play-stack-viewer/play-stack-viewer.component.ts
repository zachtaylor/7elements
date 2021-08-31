import { Component, OnInit } from '@angular/core';
import { VII } from '../7.service';
import { GlobalData } from '../api';
import { Subscription } from 'rxjs';

@Component({
  selector: 'vii-play-stack-viewer',
  templateUrl: './play-stack-viewer.component.html',
  styleUrls: ['./play-stack-viewer.component.css']
})
export class PlayStackViewerComponent implements OnInit {
  viewid : string

  private $state : Subscription

  constructor(public vii : VII) { }

  ngOnInit() {
    this.$state = this.vii.gamestate$.subscribe(state => { this.viewid = state.id })
  }

  ngOnDestroy() {
    this.$state.unsubscribe()
  }

  clickStackLeft() {
    let state = this.vii.state$(this.viewid).value
    this.viewid = state.stack
  }

  clickStackRight() {
    let state = this.vii.state$(this.viewid).value
    this.viewid = state.parent
  }

}
