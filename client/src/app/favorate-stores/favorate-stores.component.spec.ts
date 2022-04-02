import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FavorateStoresComponent } from './favorate-stores.component';

describe('FavorateStoresComponent', () => {
  let component: FavorateStoresComponent;
  let fixture: ComponentFixture<FavorateStoresComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FavorateStoresComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FavorateStoresComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
