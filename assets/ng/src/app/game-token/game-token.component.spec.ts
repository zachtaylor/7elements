import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GameTokenComponent } from './game-token.component';

describe('GameTokenComponent', () => {
  let component: GameTokenComponent;
  let fixture: ComponentFixture<GameTokenComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GameTokenComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GameTokenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
