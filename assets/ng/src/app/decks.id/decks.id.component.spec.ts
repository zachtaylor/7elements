import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DecksIdComponent } from './decks.id.component';

describe('DecksIdComponent', () => {
  let component: DecksIdComponent;
  let fixture: ComponentFixture<DecksIdComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DecksIdComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DecksIdComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
