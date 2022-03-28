import {Injectable} from "@angular/core"
import { environment } from '../environments/environments'
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {Location} from '../models/common_models'

@Injectable({
    providedIn: 'root'
})
export class MapsService{
  public lat!: number;
  public lon!: number;

    constructor(private http: HttpClient){}

    getLocation() {
        if (navigator.geolocation) {
          navigator.geolocation.getCurrentPosition((position: GeolocationPosition) => {
            if (position) {
              console.log("Latitude: " + position.coords.latitude +
                "Longitude: " + position.coords.longitude);
              this.lat = position.coords.latitude;
              this.lon = position.coords.longitude;
              environment.lat = this.lat.toString()
              environment.lon = this.lon.toString()
              
            }
          },
            (error: GeolocationPositionError) => console.log(error));
        } else {
          alert("Geolocation is not supported by this browser.");
        }
      }
      getCity(){
        this.http.post<Location>('http://localhost:10000/user', { lat: environment.lat, lon: environment.lon }).subscribe(data => {
          console.log(data.city)
          environment.city=data.city
        })
      }

}