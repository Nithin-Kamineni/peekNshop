import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {LoginModel, 
       SignupModel,
       offers,
       Stores} from '../models/common_models'
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
  favstores(user_id: string, accesskey:string, favoriteStoreId: string): Observable<any>{
    return this.http.post<any>('http://localhost:10000/user/favorate-stores', { UserId: user_id, Acesskey: accesskey, FavorateStoreId: favoriteStoreId})
  }
  getoffers(): Observable<any>{
    return  this.http.get<offers>('http://localhost:10000/offers')
  }
  getstores(lat:string, lon:string): Observable<Stores>{
    return  this.http.get<Stores>('http://localhost:10000/stores/?'+'search=store'+'&lat='+ lat+'&long='+lon, {})
  }
} 
