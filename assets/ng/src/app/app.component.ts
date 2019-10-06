import { Component, HostListener } from '@angular/core'
import { BreadcrumbService } from './breadcrumb.service'
import { ConnService } from './conn.service'
import { AdaptiveService } from './adaptive.service'
import { Game } from './api'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = '7 Elements Online'
  game : Game

  constructor(public breadcrumbService : BreadcrumbService, public conn : ConnService, public adapt : AdaptiveService) {
    conn.game$.subscribe(game => {
      this.game = game
    })
  }

  @HostListener('window:resize', ['$event.target.innerWidth'])
  onResize(width : number) {
    this.adapt.onResize(width)
  }
}
