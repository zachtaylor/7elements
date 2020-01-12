import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { Notification } from './api'

@Injectable({
  providedIn: 'root'
})
export class NotificationService {
  data$ = new BehaviorSubject<Array<Notification>>([])

  constructor() { }

  add(source: string, message: string, autodeleteT: number) {
    let nots = this.data$.value
    let not = new Notification()
    not.source = source
    not.message = message
    nots.push(not)
    this.data$.next(nots)
    let me = this.data$
    window.setTimeout(() => {
      nots = me.value
      let i = nots.indexOf(not, 0)
      if (i >= 0) {
        nots.splice(i, 1)
      }
      me.next(nots)
    }, autodeleteT)
  }

  clear() {
    this.data$.next([])
  }
}
