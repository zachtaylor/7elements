import { Component, HostListener } from '@angular/core'
import { BreadcrumbService } from './breadcrumb.service'
import { ConnService } from './conn.service'
import { AdaptiveService } from './adaptive.service'
import { NotificationService } from './notification.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = '7 Elements Online'

  constructor(
    public breadcrumbService : BreadcrumbService,
    public adapt : AdaptiveService,
    public conn : ConnService,
    public notifications : NotificationService,
    ) {
  }

  @HostListener('window:resize', ['$event.target.innerWidth'])
  onResize(width : number) {
    this.adapt.onResize(width)
  }
}
