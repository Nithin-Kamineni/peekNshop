import { Component, OnInit } from '@angular/core';
import { CatogoriesService} from './catogories.service';
import { environment } from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"; 
import { Stores } from '../models/common_models'
import { results } from '../models/results'

@Component({
  selector: 'app-catogories',
  templateUrl: './catogories.component.html',
  styleUrls: ['./catogories.component.scss']
})
export class CatogoriesComponent implements OnInit {
  city = "Gainesville"
  offers: any;
  stores:any;

  constructor(public service: CatogoriesService, private http: HttpClient) { }

  

  ngOnInit(): void {
    this.service.getOffers().subscribe(data => {
      this.offers = data;
  })

  setTimeout(() => {
  this.service.getStores().subscribe(data => {
    this.stores = data;
    console.log(this.stores.results[0].rating)
    for (var x of this.stores.results){
      console.log(x.name)
    }
    console.log(this.stores.results.icon)
  }); }, 5000);


  // setTimeout(() => {  this.http.get<Stores>('http://localhost:10000/address/?'+'search=store'+'&lat='+ environment.lat+'&long='+environment.lon, {}).subscribe( (data: Stores) => {
  //     this.stores = data
  //     console.log(this.stores.results[0].rating)
  //     console.log(this.stores.results.icon)
  //   }); }, 5000);
  }



}
