import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router, ActivatedRoute } from "@angular/router";
import {MapsService} from '../services/maps.service';
import {LoginModel} from '../models/common_models'
import { SignupModel } from '../models/common_models'
import { environment } from '../environments/environments'
import * as shajs from 'sha.js';
import { ApiService } from '../services/api.service'
import { userdetails } from '../environments/User_Details'
import { observable } from 'rxjs/internal/symbol/observable';
import { JwtHelperService } from '@auth0/angular-jwt';
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
  name = userdetails.fullname
  IsmodelShow!: boolean;
  loginmsg!: string;
  signupmsg!: string;
  isLogin = userdetails.loggedIn
  isLocation=environment.isLocation
  storesSearchForm!:FormGroup
  storesSearchText!:string;
  cartItems = environment.numberOfItemsInCart;
  returnUrl!: string;

  

  constructor(private http: HttpClient, private router: Router,public service: MapsService, private api: ApiService, private route: ActivatedRoute,) { }
  ngOnInit(): void {
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
    let token = localStorage.getItem('token');
    if (token){
      this.updateUserDetails()
      userdetails.loggedIn = true
      this.isLogin=true
      this.router.navigate(['/user-homepage'])
    }
    this.isLogin=userdetails.loggedIn
    this.isLocation=environment.isLocation
    this.cartItems=environment.numberOfItemsInCart
    console.log(this.cartItems)
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
    this.router.navigate(['/stores'])
  }
  userProfile(){
    this.router.navigate(['/user-homepage/user'])
  }
  removeFormDetails(){
    this.loginForm.controls['email'].setValue('');
    this.loginForm.controls['password'].setValue('');
    this.signupForm.controls['first_name'].setValue('');
    this.signupForm.controls['last_name'].setValue('');
    this.signupForm.controls['signup_email'].setValue('');
    this.signupForm.controls['signup_password'].setValue('');
    this.signupForm.controls['signup_confirm_password'].setValue('');
  }
  logout(){
    localStorage.removeItem('token');
    this.removeFormDetails()
    this.isLogin=false
    userdetails.loggedIn=false
    this.router.navigate(['/'])
  }
  
  updateUserDetails(){
      userdetails.id=this.CurrentUser.data.id
      userdetails.firstname=this.CurrentUser.data.firstname
      userdetails.lastname=this.CurrentUser.data.lastname
      userdetails.email=this.CurrentUser.data.email
      userdetails.password=this.CurrentUser.data.password
      userdetails.accesskey=this.CurrentUser.data.Accesskey
      userdetails.refreshkey=this.CurrentUser.data.RefreshKey
      userdetails.address1=this.CurrentUser.data.Address1
      userdetails.address2=this.CurrentUser.data.Address2
      userdetails.address3=this.CurrentUser.data.Address3
      userdetails.fullname=this.CurrentUser.data.firstname.concat(" ", this.CurrentUser.data.lastname)
      console.log(userdetails.fullname)
      this.name=userdetails.fullname
  }
  delivery(){
    console.log(this.isLogin)
    if (this.isLogin==true){
      this.router.navigate(['user-homepage/delivery'])
    }else{
      alert("Please login")
    }
    
  }
  cart(){
    this.router.navigate(['/user-homepage/user/cart'])
  }
  getNumberOfitemsInCart(){
    this.api.cartdisplay(userdetails.id).subscribe((data: any) => {
      var cartdetails= data
      var i=0
      for (let products of cartdetails){
        i = i+1
        environment.numberOfItemsInCart=i
        this.cartItems=i
      }
      console.log(this.cartItems)
    })
  }

  locationFormSubmit(){

  }
  

  loginFormSubmit(): void {
    let returnUrl = this.route.snapshot.queryParamMap.get('returnUrl') || '/';
    localStorage.setItem('returnUrl',returnUrl);
    if (this.loginForm.valid) {
      var email = this.loginForm.getRawValue().email;
      var password = this.loginForm.getRawValue().password;

      password = shajs('sha256').update(password).digest('hex')

      this.api.login(email, password).subscribe((data: any) => {
        console.log(data)
        console.log(data.JWToken)
        if(data && data.JWToken){
            localStorage.setItem('token',data.JWToken);
            this.getNumberOfitemsInCart()
            userdetails.loggedIn = true
            this.isLogin=true
            alert("Login Successful");
            let element: HTMLElement = document.getElementsByClassName('btn-close')[1] as HTMLElement;
            element.click();
            this.router.navigate(['/user-homepage'])
            this.updateUserDetails()
          }else{
            alert("Login Unsuccessful");
            this.router.navigate(['/']);
          }
        })
}}
  signupFormSubmit(): void {

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
        this.api.signup(first_name, last_name, email, password).subscribe((data: any) => {
          console.log(data)
          console.log(data.JWToken)
        if(data && data.JWToken){
            localStorage.setItem('token',data.JWToken);
            this.getNumberOfitemsInCart()
            userdetails.loggedIn = true
            this.isLogin=true
            alert("Login Successful");
            let element: HTMLElement = document.getElementsByClassName('btn-close')[2] as HTMLElement;
            element.click();
            this.router.navigate(['/user-homepage'])
            this.updateUserDetails()
          }else{
            alert("Login Unsuccessful");
            this.router.navigate(['/']);
          }
        })
      }else{
        alert("Your passwords doesn't match")
      }
      
      

    } else {
        console.log('There is a problem with the signup form');
    }  
  
  }
  get CurrentUser(){
    let token = localStorage.getItem('token');
    if(!token) return null;

    return new JwtHelperService().decodeToken(token);

  }

}

