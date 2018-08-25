import { Component, OnInit, Input, ViewChildren, ViewChild, ContentChildren } from '@angular/core';
import { QueryList } from '@angular/core/src/render3';

@Component({
  selector: 'article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.css']
})
export class ArticleComponent implements OnInit {
  @Input() title = ''

  constructor() { }

  ngOnInit() {  }
}
