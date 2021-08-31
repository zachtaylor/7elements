import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlayChoiceViewerComponent } from './play-choice-viewer.component';

describe('PlayChoiceViewerComponent', () => {
  let component: PlayChoiceViewerComponent;
  let fixture: ComponentFixture<PlayChoiceViewerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlayChoiceViewerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlayChoiceViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
