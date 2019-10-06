import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GameSeatComponent } from './game-seat.component';

describe('GameSeatComponent', () => {
  let component: GameSeatComponent;
  let fixture: ComponentFixture<GameSeatComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GameSeatComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GameSeatComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
