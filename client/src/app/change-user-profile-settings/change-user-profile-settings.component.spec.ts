import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangeUserProfileSettingsComponent } from './change-user-profile-settings.component';

describe('ChangeUserProfileSettingsComponent', () => {
  let component: ChangeUserProfileSettingsComponent;
  let fixture: ComponentFixture<ChangeUserProfileSettingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChangeUserProfileSettingsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ChangeUserProfileSettingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
