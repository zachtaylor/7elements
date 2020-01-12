import { Component, OnInit } from '@angular/core'
import { Deck } from '../api'
import { ActivatedRoute } from '@angular/router'
import { ConnService } from '../conn.service'
import { Subscription } from 'rxjs'

@Component({
  selector: 'app-decks.id',
  templateUrl: './decks.id.component.html',
  styleUrls: ['./decks.id.component.css']
})
export class DecksIdComponent implements OnInit {
  deck: Deck
  private id: number
  private $id: Subscription
  private $glob: Subscription

  constructor(private route: ActivatedRoute, public conn : ConnService) { }

  ngOnInit() {
    let me = this

    this.$id = this.route.params.subscribe(params => {
       this.id = +params['id']; // (+) converts string 'id' to a number
    })
    this.$glob = this.conn.global$.subscribe(glob => {
      if (glob) me.deck = me.conn.deck(me.id)
    })
  }

  ngOnDestroy() {
    this.$id.unsubscribe()
  }

}
