import { Component } from '@angular/core'
import { MessageService } from './message.service'
import { UserService } from './user.service'
import { BreadcrumbService } from './breadcrumb.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = '7 Elements Online'
  breadcrumb = '/'

  constructor(public breadcrumbService : BreadcrumbService, public messageService : MessageService, public userService : UserService) {
    // messageService.add('wzzaaappp')
  }
}
