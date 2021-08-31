import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'

@Injectable({
  providedIn: 'root'
})
export class AdaptiveService { 
  /**
   * Scale indicates the magnitude of the display size
   * 
   * 0 none
   * 
   * 1 mobile
   * 
   * 2 medium
   * 
   * 3 large
   */
  scale$ = new BehaviorSubject<number>(0)

  constructor() {
    this.scale$.subscribe(scale => this.onScale(scale))
    this.onResize(window.innerWidth)
  }

  isSmall() : boolean {
    return this.scale$.value == 1
  }
  isMedium() : boolean {
    return this.scale$.value == 2
  }
  isLarge() : boolean {
    return this.scale$.value == 3
  }

  onResize(width : number) {
    if (width < 600) {
      this.setScale(1)
    } else if (width < 900) {
      this.setScale(2)
    } else {
      this.setScale(3)
    }
  }

  setScale(scale : number) {
    if (this.scale$.value != scale) {
      console.debug('adapt: update layout scale:', scale)
      this.scale$.next(scale)
    }
  }

  onScale(scale : number) {
    if (scale <= 1) {
      this.setProperties(32, 2, 18)
    } else if (scale <= 2) {
      this.setProperties(36, 4, 21)
    } else {
      this.setProperties(48, 8, 24)
    }
  }

  setProperties(uiSize : number, paddingSize : number, fontSize : number) {
    document.body.style.setProperty('--ui', uiSize+'px')
    document.body.style.setProperty('--pad', paddingSize+'px')
    document.body.style.setProperty('--font', fontSize+'px')
    document.body.style.setProperty('font-size', fontSize+'px')
  }
}
