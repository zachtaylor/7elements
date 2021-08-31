import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminSinkComponent } from './admin-sink.component';

describe('AdminSinkComponent', () => {
  let component: AdminSinkComponent;
  let fixture: ComponentFixture<AdminSinkComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminSinkComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminSinkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
