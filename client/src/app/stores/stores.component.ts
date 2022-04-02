import { Component, OnInit, Input } from '@angular/core';
import { Stores } from '../models/common_models'
import { environment} from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"
@Component({
  selector: 'app-stores',
  templateUrl: './stores.component.html',
  styleUrls: ['./stores.component.scss']
})
export class StoresComponent implements OnInit {
isFavorite = false
stores:any; 
storesSearchText=environment.storesSearchText
city = environment.city
  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.http.get<Stores>('http://localhost:10000/stores/?'+'search='+environment.storesSearchText+'&lat='+ environment.lat+'&long='+environment.lon, {}).subscribe( (data: Stores) => {
      console.log(data)
      this.stores = data;
      console.log(this.stores.results[0].rating)

    })
  }
  favorite(){
    if(this.isFavorite==false){
      this.isFavorite=true
      document.getElementsByTagName("a")[6].style.backgroundColor = "pink";
    }else{
      this.isFavorite=false
      document.getElementsByTagName("a")[6].style.backgroundColor = "gray";
    }
    
  }
  
}
