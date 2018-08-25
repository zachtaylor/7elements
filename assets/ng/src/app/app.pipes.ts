import { Pipe, PipeTransform } from '@angular/core'
import { DomSanitizer } from '@angular/platform-browser'

@Pipe({ name: 'keepHtml', pure: false })
export class EscapeHtmlPipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) { }

  transform(content) {
    return this.sanitizer.bypassSecurityTrustHtml(content)
  }
}

@Pipe({ name: 'mapKeys', pure: false })
export class MapKeysPipe implements PipeTransform {
  constructor() { }

  transform(content : Map<any, any>) {
    if (!content) return [];
    return Object.keys(content)
  }
}

@Pipe({ name: 'mapValues', pure: false })
export class MapValuesPipe implements PipeTransform {
  constructor() { }

  transform(content : Map<any, any>) {
    if (!content) return [];
    return Object.values(content)
  }
}

@Pipe({ name: 'count', pure: false })
export class CountPipe implements PipeTransform {
  constructor() { }

  transform(content : number) {
    return Array(content)
  }
}
