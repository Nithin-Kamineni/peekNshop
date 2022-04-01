import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from "@angular/router";
import {MapsService} from '../services/maps.service';
import {LoginModel} from '../models/common_models'
import { SignupModel } from '../models/common_models'
import { environment } from '../environments/environments'
import * as shajs from 'sha.js';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss']
})
export class SidenavComponent implements OnInit {
  loginForm!: FormGroup;
  signupForm!: FormGroup;
  locationForm!: FormGroup
  city = environment.city
  name = environment.fullname
  IsmodelShow!: boolean;
  loginmsg!: string;
  signupmsg!: string;
  isLogin = environment.isLogin
  isLocation=environment.isLocation
  storesSearchForm!:FormGroup
  storesSearchText!:string

  

  constructor(private http: HttpClient, private router: Router,public service: MapsService) { }
  ngOnInit(): void {

    this.isLogin=true
     

    this.loginForm = new FormGroup({
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })
    this.locationForm = new FormGroup({
      street: new FormControl('', [Validators.required]),
      city: new FormControl('', [Validators.required]),
      state: new FormControl('', [Validators.required]),
      zipcode: new FormControl('', [Validators.required])
    })

    this.signupForm = new FormGroup({
      first_name: new FormControl('', [Validators.required]),
      last_name: new FormControl('', [Validators.required]),
      signup_email: new FormControl('', [Validators.required, Validators.email]),
      signup_password: new FormControl('', [Validators.required]),
      signup_confirm_password: new FormControl('', [Validators.required])
    })
    this.storesSearchForm = new FormGroup({
      storesSearchText: new FormControl('', [Validators.required])
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
  get streetField(): any {
    return this.locationForm.get('street');
  }
  get cityField(): any {
    return this.locationForm.get('city');
  }
  get stateField(): any {
    return this.locationForm.get('state');
  }
  get zipcodeField(): any {
    return this.locationForm.get('zipcode');
  }
  storesSearch(){
    this.storesSearchText=this.storesSearchForm.getRawValue().storesSearchText
    console.log(this.storesSearchText)
    environment.storesSearchText=this.storesSearchText
  }
  userProfile(){
    this.router.navigate(['/user-homepage/user'])
  }
  logout(){
    this.updateisLogin()
    console.log(environment.isLogin)
    console.log("Logout")
    this.router.navigate(['/'])
  }
  updateisLogin(){
    this.isLogin=!this.isLogin
    environment.isLogin=!environment.isLogin
  }
  delivery(){
    if (environment.isLogin=false){
      this.router.navigate(['user-homepage/delivery'])
    }else{
      alert("Please login")
    }
    
  }

  locationFormSubmit(){

  }
  

  loginFormSubmit(): void {

    if (this.loginForm.valid) {
      var email = this.loginForm.getRawValue().email;
      var password = this.loginForm.getRawValue().password;
      console.log(email,password)
      password = shajs('sha256').update(password).digest('hex')
      // this.http.post<any>('http://localhost:10000/students/', { Email: email, Password: password }).subscribe(data => { })
      var user = "email=" + email + "&passkey=" + password
      this.http.get<LoginModel>('http://localhost:10000/user?'+"email=" + email + "&passkey=" + password, {}).subscribe( (data: LoginModel) => {
          this.loginmsg = data.Msg;
          console.log(data);
          var details = Object.values(data.UserDetails)
          environment.id=details[0]
          environment.firstname=details[1]
          environment.lastname=details[2]
          environment.email=details[3]
          environment.password=details[4]
          environment.accesskey=details[5]
          environment.refreshkey=details[6]
          environment.address1=details[7]
          environment.address2=details[8]
          environment.address3=details[9]
          environment.fullname=environment.firstname.concat(" ", environment.lastname)
          this.name=environment.firstname.concat(" ", environment.lastname)
          
          if (this.loginmsg == "Login Sucessfull"){
            alert(this.loginmsg) 
            let element: HTMLElement = document.getElementsByClassName('btn-close')[1] as HTMLElement;
            element.click();
            this.updateisLogin()
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
      password = shajs('sha256').update(password).digest('hex')
      console.log(password)
      var confirm_password = this.signupForm.getRawValue().signup_confirm_password;
      confirm_password = shajs('sha256').update(confirm_password).digest('hex')
      console.log(confirm_password)
      if (password==confirm_password){
        this.http.post<SignupModel>('http://localhost:10000/user', { First_name: first_name, Last_name: last_name, Email: email, Password: password }).subscribe(data => {
            console.log(data.Msg)
            this.signupmsg = data.Msg
            if (this.signupmsg == "Sucessfull"){
              alert("Signup Successful")
              let element: HTMLElement = document.getElementsByClassName('btn-close')[2] as HTMLElement;
              element.click();
              this.updateisLogin()
              this.router.navigate(['/user-homepage'])
            }else{
              console.log("Wrong User")
              alert("User already registered")
              let element: HTMLElement = document.getElementsByClassName('btn-close')[2] as HTMLElement;
              element.click();
              this.router.navigate([''])
            }
        })
      }else{
        alert("Your passwords doesn't match")
      }
      
      

  } else {
      console.log('There is a problem with the signup form');
  }  
  
}

}

