import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Router } from "@angular/router";
import {MapsService} from '../services/maps.service';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss']
})
export class SidenavComponent implements OnInit {
  loginForm!: FormGroup;
  signupForm!: FormGroup;
  city = "Gainesville"
  IsmodelShow!: boolean;

  constructor(private http: HttpClient, private router: Router,public service: MapsService) { }
  ngOnInit(): void {

     

    

    this.loginForm = new FormGroup({
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })

    this.signupForm = new FormGroup({
      first_name: new FormControl('', [Validators.required]),
      last_name: new FormControl('', [Validators.required]),
      signup_email: new FormControl('', [Validators.required, Validators.email]),
      signup_password: new FormControl('', [Validators.required]),
      signup_confirm_password: new FormControl('', [Validators.required])
    })
  }
  get emailField(): any {
    return this.loginForm.get('email');
  }
  get passwordField(): any {
    return this.loginForm.get('password');
  }
  get firstNameField(): any {
    return this.signupForm.get('first_name');
  }
  get lastNameField(): any {
    return this.signupForm.get('last_name');
  }
  get signupEmailField(): any {
    return this.signupForm.get('signup_email');
  }
  get signupPasswordField(): any {
    return this.signupForm.get('signup_password');
  }
  get signupConfirmPasswordField(): any {
    return this.signupForm.get('signup_confirm_password');
  }

  close() {
    //Can I close modal window manually?
    this.IsmodelShow=false;
  }

  loginFormSubmit(): void {

    if (this.loginForm.valid) {
      var email = this.loginForm.getRawValue().email;
      var password = this.loginForm.getRawValue().password;
      console.log(email,password)
      let params = new HttpParams().set('email',email).set('passkey',password);
      // this.http.post<any>('http://localhost:10000/students/', { Email: email, Password: password }).subscribe(data => { })
      console.log(params)
      this.http.get('http://localhost:10000/user', {params: params}).subscribe(data => { })
      // this.router.navigate(['/user-homepage'])
      
  } else {
      console.log('There is a problem with the login form');
  }
}
  signupFormSubmit(): void {

    console.log(this.signupForm.getRawValue())
    console.log(this.signupForm.getRawValue().signup_email)

    if (this.signupForm.valid) {
      var first_name = this.signupForm.getRawValue().first_name;
      var last_name = this.signupForm.getRawValue().last_name;
      var email = this.signupForm.getRawValue().signup_email;
      var password = this.signupForm.getRawValue().signup_password;
      var confirm_password = this.signupForm.getRawValue().signup_confirm_password;
      console.log(first_name, last_name, email, password, confirm_password)
      
      this.http.post<any>('https://localhost:10000/user', { Email: email, Password: password}).subscribe(data => {})
      this.router.navigate(['/user-homepage'])
    } else {
      console.log('There is a problem with the signup form');
  }  
  
}

}

