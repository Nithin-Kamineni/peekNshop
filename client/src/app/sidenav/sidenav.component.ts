import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from "@angular/router";
import {MapsService} from '../services/maps.service';
import {LoginModel} from '../models/common_models'
import { SignupModel } from '../models/common_models'
import { data } from 'cypress/types/jquery';
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
  loginmsg!: string;
  signupmsg!: string

  

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

  

  loginFormSubmit(): void {

    if (this.loginForm.valid) {
      var email = this.loginForm.getRawValue().email;
      var password = this.loginForm.getRawValue().password;
      console.log(email,password)
      // this.http.post<any>('http://localhost:10000/students/', { Email: email, Password: password }).subscribe(data => { })
      var user = "email=" + email + "&passkey=" + password
      this.http.get<LoginModel>('http://localhost:10000/user?'+"email=" + email + "&passkey=" + password, {}).subscribe( (data: LoginModel) => {
          this.loginmsg = data.Msg;
          console.log(data);
          console.log(this.loginmsg)
          if (this.loginmsg == "Login Sucessfull"){
            alert(this.loginmsg) 
            let element: HTMLElement = document.getElementsByClassName('btn-close')[0] as HTMLElement;
            element.click();
           
            this.router.navigate(['/user-homepage'])
          }else{
            alert(this.loginmsg)
            this.router.navigate([''])
          }
        })
        
      
      
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
      
      this.http.post<SignupModel>('http://localhost:10000/user', { First_name: first_name, Last_name: last_name, Email: email, Password: password }).subscribe(data => {
            console.log(data.Msg)
            this.signupmsg = data.Msg
            if (this.signupmsg == "sucessfull"){
              alert("Signup Successful")
              let element: HTMLElement = document.getElementsByClassName('btn-close')[1] as HTMLElement;
                element.click();
              this.router.navigate(['/user-homepage'])
            }else{
              alert("User already registered")
            }
        })
        

  } else {
      console.log('There is a problem with the signup form');
  }  
  
}

}

