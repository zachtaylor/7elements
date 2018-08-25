import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DeckSummaryComponent } from './deck-summary.component';

describe('DeckSummaryComponent', () => {
  let component: DeckSummaryComponent;
  let fixture: ComponentFixture<DeckSummaryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeckSummaryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeckSummaryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
