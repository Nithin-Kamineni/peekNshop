import {Injectable} from "@angular/core"
import { environment } from '../environments/environments'
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {Location} from '../models/common_models'
import { NavigationEnd, Router } from "@angular/router";
@Injectable({
    providedIn: 'root'
})
export class MapsService{
  public lat!: number;
  public lon!: number;

    constructor(private http: HttpClient, private router: Router){}

    getLocation() {
      if(environment.lat == ""){
        if (navigator.geolocation) {
          navigator.geolocation.getCurrentPosition((position: GeolocationPosition) => {
            if (position) {
              console.log("Latitude: " + position.coords.latitude +
                "Longitude: " + position.coords.longitude);
              this.lat = position.coords.latitude;
              this.lon = position.coords.longitude;
              environment.lat = this.lat.toString()
              environment.lon = this.lon.toString()
              environment.isLocation=true
              
            }
          },
            (error: GeolocationPositionError) => console.log(error));
        } else {
          environment.isLocation=false
          alert("Geolocation is not supported by this browser.");
        }
      }
        
      }
      getCity(){
        this.http.post<Location>('http://localhost:10000/address/city', { Lat: environment.lat, Lng: environment.lon }).subscribe(data => {
          console.log(data.city)
          if(environment.lat == ""){
            environment.isLocation=false
          }
          console.log(environment.isLocation)
          environment.city=data.city
        })
      }

}