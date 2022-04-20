import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {LoginModel, 
       SignupModel,
       offers,
       Stores} from '../models/common_models'
import { Observable } from 'rxjs';
import { environment } from '../environments/environments';
import { userdetails } from '../environments/User_Details';

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
  getProducts(){
    return this.http.get<any>('http://localhost:10000/stores/items?store_id='+environment.storeId+"&user_id="+userdetails.id, {})
  }
  addtocart(user_id:string, productID:string, quantity:string, created:string, modified:string){
    return this.http.post<any>('http://localhost:10000/cart/additem', {"user_id":user_id,"productID":productID, "quantity":quantity,"created":created, "modified":modified})
  }
  cartdisplay(user_id:string){
    return this.http.post<any>('http://localhost:10000/cart', {"user_id":user_id})
  } 
  removeProductFromCart(user_id:string, productID:string){
    return this.http.put<any>('http://localhost:10000/cart/', {"user_id":user_id,"productID":productID})
  }
  emptycart(user_id:String){
    return this.http.post<any>('http://localhost:10000/cart/clear-cart', {"user_id":user_id})
  }
  location(street:string, city:string, state:string, zipcode:string){
    return this.http.post<any>('http://localhost:10000/address', {"Street":street, "City":city, "State":state, "Zipcode":zipcode})
  }
  displayFavoriteStores(UserID:string){
    return this.http.post<any>('http://localhost:10000/user/favorate-stores', { "UserID": UserID})
  }
} 
