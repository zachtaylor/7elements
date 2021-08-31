import { Component, HostListener } from '@angular/core'
import { AdaptiveService } from './adaptive.service'
import { CookieService } from 'ngx-cookie-service'
import { VII } from './7.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = '7 Elements Online'

  constructor(public adapt : AdaptiveService, public cookies : CookieService, private vii: VII) { }

  @HostListener('window:resize', ['$event.target.innerWidth'])
  onResize(width : number) { this.adapt.onResize(width) }

  ngOnInit() { }

  ngOnDestroy() { }

  clickAllowCookies() {
    this.cookies.set('allow', 'true', 90, undefined, undefined, window.location.protocol == 'https:', 'Strict')
  }
}
