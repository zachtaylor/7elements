import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlayHandComponent } from './play-hand.component';

describe('PlayHandComponent', () => {
  let component: PlayHandComponent;
  let fixture: ComponentFixture<PlayHandComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlayHandComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlayHandComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
