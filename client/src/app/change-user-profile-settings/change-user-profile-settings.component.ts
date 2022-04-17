import { Component, OnInit } from '@angular/core';
import {environment} from '../environments/environments'
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { userdetails } from '../environments/User_Details'
@Component({
  selector: 'app-change-user-profile-settings',
  templateUrl: './change-user-profile-settings.component.html',
  styleUrls: ['./change-user-profile-settings.component.scss']
})
export class ChangeUserProfileSettingsComponent implements OnInit {
  name=userdetails.fullname
  firstname = userdetails.firstname
  lastname = userdetails.lastname
  email = userdetails.email
  city=environment.city
  phone=userdetails.phone
  userDetailsForm!: FormGroup;

  constructor() { }

  ngOnInit(): void {

    this.userDetailsForm= new FormGroup({
      firstname: new FormControl(this.firstname, [Validators.required]),
      lastname: new FormControl(this.lastname, [Validators.required]),
      phone: new FormControl(this.phone, [Validators.required]),
      email: new FormControl(this.email, [Validators.required]),
      street: new FormControl('Street', [Validators.required]),
      city: new FormControl(this.city, [Validators.required]),
      state: new FormControl('State', [Validators.required]),
      zipcode: new FormControl('Zipcode', [Validators.required])
    })
    
  }
  userDetailsUpdateRequest(){

  }

}
