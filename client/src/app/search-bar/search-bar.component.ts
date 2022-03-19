import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"; 
import { Router } from "@angular/router";
import {MapsService} from '../services/maps.service';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { environment } from '../environments/environments'
import { Products } from '../models/common_models'


@Component({
  selector: 'app-search-bar',
  templateUrl: './search-bar.component.html',
  styleUrls: ['./search-bar.component.scss']
})
export class SearchBarComponent implements OnInit{
  searchText1! : string
  searchForm!: FormGroup;
  


  ngOnInit(): void {
    this.searchForm = new FormGroup({
      searchText: new FormControl('', [Validators.required])
    })
  }

  constructor(private http: HttpClient, private router: Router) { }

  searchRequest(){
    this.searchText1 = this.searchForm.getRawValue().searchText
    console.log(this.searchText1)
    environment.searchText = this.searchText1
    // this.http.get<Products>('http://localhost:10000/user?'+"search_text=" +environment.searchText + "&lat=" + environment.lat + "&lon=" + environment.lon, {}).subscribe( (data: Products) => {
    //   console.log(data.Msg)
    // })
    this.router.navigate(['/products'])
    
      
  }
// We should give this in products page oninit
//  
  
}
