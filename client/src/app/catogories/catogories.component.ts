import { Component, OnInit } from '@angular/core';
import { CatogoriesService} from './catogories.service';
import { environment } from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"; 
import { Stores } from '../models/common_models'
import { results } from '../models/results'
import { Router } from "@angular/router";
@Component({
  selector: 'app-catogories',
  templateUrl: './catogories.component.html',
  styleUrls: ['./catogories.component.scss']
})
export class CatogoriesComponent implements OnInit {
  city = "Gainesville"
  offers: any;
  stores:any; 
  isFavorite = false;
  toggle = true;
  status = 'Enable';
  arr = new Array(19).fill(false);
  storesarr = new Array(19)

  constructor(public service: CatogoriesService, private http: HttpClient, private router: Router) { }

  

  ngOnInit(): void {
    this.service.getOffers().subscribe(data => {
      this.offers = data;

  })

  

  setTimeout(() => {
  this.service.getStores().subscribe(data => {
    this.stores = data;
    console.log(this.stores.results[0].rating)
    var i = 0 
    for (var x of this.stores.results){
      this.storesarr[i] = this.stores.results[i].place_id
      i = i+1
    }
    console.log(this.storesarr)
    console.log(this.stores.results.icon)
  }); }, 4000);


  // setTimeout(() => {  this.http.get<Stores>('http://localhost:10000/address/?'+'search=store'+'&lat='+ environment.lat+'&long='+environment.lon, {}).subscribe( (data: Stores) => {
  //     this.stores = data
  //     console.log(this.stores.results[0].rating)
  //     console.log(this.stores.results.icon)
  //   }); }, 5000);
  }
  enableDisableRule(job: any) {
    this.toggle = !this.toggle;
    this.status = this.toggle ? 'Enable' : 'Disable';
}
  visitStore(i:number){
    
    environment.storeId = this.storesarr[i]
    console.log(environment.storeId)
    this.router.navigate(['/products'])
  }

  favorite(i:number){
    if (environment.isLogin=true){
      var k = 8+(2*i)
      var favoriteStoreId = this.stores.results[i].place_id
      var user_id=environment.id
      console.log(favoriteStoreId)
      console.log(user_id)
      var accesskey = environment.accesskey
      this.http.post<any>('http://localhost:10000/user/favorate-stores', { UserId: user_id, Acesskey: accesskey, FavorateStoreId: favoriteStoreId}).subscribe(data => {
      console.log(data)
      })
      if(this.arr[i]==false){
        this.arr[i]=true
        document.getElementsByTagName("a")[k].style.backgroundColor = "pink";
      }else{
        this.arr[i]=false
        document.getElementsByTagName("a")[k].style.backgroundColor = "gray";
      }
    }else{
      alert('please login')
    }
    
    
  }



}
