import {Injectable} from "@angular/core"
import { environment } from '../environments/environments'

@Injectable({
    providedIn: 'root'
})
export class MapsService{
  public lat!: number;
  public lon!: number;

    constructor(){}

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

}