import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InventoryValueComponent } from './inventory-value.component';

describe('InventoryValueComponent', () => {
  let component: InventoryValueComponent;
  let fixture: ComponentFixture<InventoryValueComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [InventoryValueComponent]
    });
    fixture = TestBed.createComponent(InventoryValueComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
