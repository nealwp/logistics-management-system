import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddAllowanceFormComponent } from './add-allowance-form.component';

describe('AddAllowanceFormComponent', () => {
  let component: AddAllowanceFormComponent;
  let fixture: ComponentFixture<AddAllowanceFormComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AddAllowanceFormComponent]
    });
    fixture = TestBed.createComponent(AddAllowanceFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
