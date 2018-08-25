import { ActivatedRoute, Router, NavigationEnd } from '@angular/router'
import { Injectable } from '@angular/core'
import { filter } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BreadcrumbService {
  path = "/"

  constructor(private activatedRoute: ActivatedRoute, private router:Router) {
    router.events.subscribe(e => {
      if (e instanceof NavigationEnd) {
        this.path = e.url
      }
    })
  }
}
