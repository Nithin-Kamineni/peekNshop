import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from "@angular/router";
import {MapsService} from '../services/maps.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  loginForm!: FormGroup;
  signupForm!: FormGroup;
  city = "Gainesville"

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
  get signupEmailField(): any {
    return this.signupForm.get('signup_email');
  }
  get signupPasswordField(): any {
    return this.signupForm.get('signup_password');
  }

  loginFormSubmit(): void {

    if (this.loginForm.valid) {
      var email = this.loginForm.getRawValue().email;
      var password = this.loginForm.getRawValue().password;
      console.log(email,password)
      this.http.post<any>('http://localhost:10000/students/', { Email: email, Password: password }).subscribe(data => { })
      this.router.navigate(['/user-homepage'])
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
      
      this.http.post<any>('https://reqres.in/api/posts', { First_name: first_name, Last_name: last_name, Email: email, Password: password }).subscribe(data => {
            
        })
  } else {
      console.log('There is a problem with the signup form');
  }  
  
}

}
