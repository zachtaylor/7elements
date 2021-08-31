import { Component, OnInit } from '@angular/core'
import { Deck, GlobalData } from '../api'
import { ActivatedRoute } from '@angular/router'
import { VII } from '../7.service'
import { Subscription } from 'rxjs'

@Component({
  selector: 'app-decks.id',
  templateUrl: './decks.id.component.html',
  styleUrls: ['./decks.id.component.css']
})
export class DecksIdComponent implements OnInit {
  glob: GlobalData
  deck: Deck
  private id: number
  private $id: Subscription
  private $glob: Subscription

  constructor(private route: ActivatedRoute, public vii : VII) { }

  ngOnInit() {
    this.$id = this.route.params.subscribe(params => {
       this.id = +params['id']
    })
    this.$glob = this.vii.global$.subscribe(glob => {
      if (glob) {
        this.glob = glob
        this.deck = glob.decks[this.id]
      }
    })
  }

  ngOnDestroy() {
    this.$id.unsubscribe()
    this.$glob.unsubscribe()
  }

}
