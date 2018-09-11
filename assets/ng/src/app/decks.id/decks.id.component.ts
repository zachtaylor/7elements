import { Component, OnInit, OnDestroy } from '@angular/core'
import { ActivatedRoute } from '@angular/router'
import { PingService } from '../ping.service'

@Component({
  selector: 'app-decks.id',
  templateUrl: './decks.id.component.html',
  styleUrls: ['./decks.id.component.css']
})
export class DecksIdComponent implements OnInit {
  id: number;
  private sub: any;

  constructor(private route: ActivatedRoute, public pingService : PingService) { }

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
       this.id = +params['id']; // (+) converts string 'id' to a number
    });
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
