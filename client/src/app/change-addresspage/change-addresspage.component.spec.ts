import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangeAddresspageComponent } from './change-addresspage.component';

describe('ChangeAddresspageComponent', () => {
  let component: ChangeAddresspageComponent;
  let fixture: ComponentFixture<ChangeAddresspageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChangeAddresspageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ChangeAddresspageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
