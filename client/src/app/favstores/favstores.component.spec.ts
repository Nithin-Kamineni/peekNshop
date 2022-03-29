import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FavstoresComponent } from './favstores.component';

describe('FavstoresComponent', () => {
  let component: FavstoresComponent;
  let fixture: ComponentFixture<FavstoresComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FavstoresComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FavstoresComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
