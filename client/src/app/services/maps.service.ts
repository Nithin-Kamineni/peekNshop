import {Injectable} from "@angular/core"

@Injectable({
    providedIn: 'root'
})
export class MapsService{
  public lat!: number;
  public lng!: number;

    constructor(){}

    getLocation() {
        if (navigator.geolocation) {
          navigator.geolocation.getCurrentPosition((position: GeolocationPosition) => {
            if (position) {
              console.log("Latitude: " + position.coords.latitude +
                "Longitude: " + position.coords.longitude);
              this.lat = position.coords.latitude;
              this.lng = position.coords.longitude;
              console.log(this.lat);
              console.log(this.lat);
            }
          },
            (error: GeolocationPositionError) => console.log(error));
        } else {
          alert("Geolocation is not supported by this browser.");
        }
      }

}