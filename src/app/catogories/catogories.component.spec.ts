import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CatogoriesComponent } from './catogories.component';

describe('CatogoriesComponent', () => {
  let component: CatogoriesComponent;
  let fixture: ComponentFixture<CatogoriesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CatogoriesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CatogoriesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
