import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IssueEffectivenessComponent } from './issue-effectiveness.component';

describe('IssueEffectivenessComponent', () => {
  let component: IssueEffectivenessComponent;
  let fixture: ComponentFixture<IssueEffectivenessComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [IssueEffectivenessComponent]
    });
    fixture = TestBed.createComponent(IssueEffectivenessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
