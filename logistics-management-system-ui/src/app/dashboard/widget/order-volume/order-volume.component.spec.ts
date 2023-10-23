import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrderVolumeComponent } from './order-volume.component';

describe('OrderVolumeComponent', () => {
  let component: OrderVolumeComponent;
  let fixture: ComponentFixture<OrderVolumeComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [OrderVolumeComponent]
    });
    fixture = TestBed.createComponent(OrderVolumeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
