import { Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'
import { MyAccount } from './api';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: `root`
})
export class UserService {
  data$ = new BehaviorSubject<MyAccount>(null)

  constructor(private http: HttpClient, private cookieService : CookieService) { 
    var sessionID = cookieService.get(`SessionID`)
    if (sessionID) {
      this.Subscribe()
    }
  }

  Subscribe() {
    this.Request().subscribe(account => {
      if (account) {
        console.debug('user.service', account);
        this.data$.next(account);
      }
      else this.cookieService.delete(`SessionID`);
    });
  }

  Request() {
    return this.http.get<MyAccount>(`/api/myaccount.json`)
  }
}
