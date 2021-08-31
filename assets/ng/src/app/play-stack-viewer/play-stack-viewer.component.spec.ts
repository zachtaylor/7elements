import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlayStackViewerComponent } from './play-stack-viewer.component';

describe('PlayStackViewerComponent', () => {
  let component: PlayStackViewerComponent;
  let fixture: ComponentFixture<PlayStackViewerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlayStackViewerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlayStackViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
