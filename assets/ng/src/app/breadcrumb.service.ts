import { ActivatedRoute, Router, NavigationEnd } from '@angular/router'
import { Injectable } from '@angular/core'

@Injectable({
  providedIn: 'root'
})
export class BreadcrumbService {
  path = "/"

  constructor(private activatedRoute: ActivatedRoute, private router:Router) {
    router.events.subscribe(e => {
      if (e instanceof NavigationEnd) {
        this.path = window.location.pathname
      }
    })
  }
}
