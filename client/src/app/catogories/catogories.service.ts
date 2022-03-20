import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http";
import { Stores } from '../models/common_models'
import { environment } from '../environments/environments'
@Injectable({
    providedIn: 'root'
})
export class CatogoriesService{

    constructor(private http: HttpClient){}

    getOffers(){

        //Headers
        const httpHeaders = new HttpHeaders();
        httpHeaders.append('content-type','application/json')

        //Get the HTTP Method working for you
        return  this.http.get('http://localhost:8080/api/projects', {headers: httpHeaders});

    }
    getStores(){
        
        return  this.http.get<Stores>('http://localhost:10000/address/?'+'search=store'+'&lat='+ environment.lat+'&long='+environment.lon, {});

    }

    
}