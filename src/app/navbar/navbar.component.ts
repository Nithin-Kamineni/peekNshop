import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  
  loginForm!: FormGroup;

  constructor(private http: HttpClient) { }
  ngOnInit(): void {

    
    
    this.loginForm = new FormGroup({
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })
  }
  get emailField(): any {
    return this.loginForm.get('email');
  }
  get passwordField(): any {
    return this.loginForm.get('password');
  }

  loginFormSubmit(): void {

    if (this.loginForm.valid) {
      var email = this.loginForm.getRawValue().email;
      var password = this.loginForm.getRawValue().password;
      console.log(email,password)
      this.http.post<any>('https://reqres.in/api/posts', { email: email, password: password }).subscribe(data => {
            
        })
  } else {
      console.log('There is a problem with the form');
  }
    // Call Api
  }
}
