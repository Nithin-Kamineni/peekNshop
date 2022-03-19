import { Component, Input, OnInit} from '@angular/core';
import {MapsService} from '../services/maps.service';
import { environment } from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"; 
import { Stores } from '../models/common_models'

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit{
  title = 'PeekNshop';
  subtitle = 'choose where to shop?';
  

  constructor(public service: MapsService, private http: HttpClient) { }

  ngOnInit(): void {
    this.service.getLocation();
    setTimeout(() => {  this.http.get<Stores>('http://localhost:10000/address/?'+'search=store'+'&lat='+ environment.lat+'&long='+environment.lon, {}).subscribe( (data: Stores) => {
      console.log(JSON.stringify(data))
      console.log(data.icon)
    }); }, 5000);
    

  }

  

}
