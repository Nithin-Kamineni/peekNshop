import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http";


@Injectable({
    providedIn: 'root'
})
export class CatogoriesService{

    constructor(private httpclient: HttpClient){}

    getOffers(){

        //Headers
        const httpHeaders = new HttpHeaders();
        httpHeaders.append('content-type','application/json')

        //Get the HTTP Method working for you
        return  this.httpclient.get('http://localhost:8080/api/projects', {headers: httpHeaders});

    }

    
}