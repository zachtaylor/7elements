import { Component } from '@angular/core'
import { BreadcrumbService } from './breadcrumb.service'
import { ConnService } from './conn.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = '7 Elements Online'
  breadcrumb = '/'

  constructor(public breadcrumbService : BreadcrumbService, public conn : ConnService) {
  }
}
