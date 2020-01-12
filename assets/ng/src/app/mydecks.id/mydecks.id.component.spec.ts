import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyDecksIdComponent } from './mydecks.id.component';

describe('Mydecks.IdComponent', () => {
  let component: MyDecksIdComponent;
  let fixture: ComponentFixture<MyDecksIdComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [MyDecksIdComponent]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyDecksIdComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
