import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DelivaryPageComponent } from './delivary-page.component';

describe('DelivaryPageComponent', () => {
  let component: DelivaryPageComponent;
  let fixture: ComponentFixture<DelivaryPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DelivaryPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DelivaryPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
