import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {LoginModel, 
       SignupModel} from '../models/common_models'
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }


  login(email: string, password:string): Observable<LoginModel>{
    return this.http.get<LoginModel>('http://localhost:10000/user?'+"email=" + email + "&passkey=" + password)
  }
  signup(first_name: string, last_name:string, email: string, password:string): Observable<SignupModel>{
    return this.http.post<SignupModel>('http://localhost:10000/user', { firstname: first_name, lastname: last_name, email: email, password: password })
  }

} 
