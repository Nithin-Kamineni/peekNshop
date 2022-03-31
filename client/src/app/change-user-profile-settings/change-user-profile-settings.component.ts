import { Component, OnInit } from '@angular/core';
import {environment} from '../environments/environments'
@Component({
  selector: 'app-change-user-profile-settings',
  templateUrl: './change-user-profile-settings.component.html',
  styleUrls: ['./change-user-profile-settings.component.scss']
})
export class ChangeUserProfileSettingsComponent implements OnInit {
  name=environment.fullname
  firstname = environment.firstname
  lastname = environment.lastname
  email = environment.email
  city=environment.city
  constructor() { }

  ngOnInit(): void {
    
  }

}
