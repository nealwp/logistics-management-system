import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllowanceListComponent } from './allowance-list.component';

describe('AllowanceListComponent', () => {
  let component: AllowanceListComponent;
  let fixture: ComponentFixture<AllowanceListComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AllowanceListComponent]
    });
    fixture = TestBed.createComponent(AllowanceListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
