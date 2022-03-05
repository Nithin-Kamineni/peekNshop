import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"; 
import { Router } from "@angular/router";
import {MapsService} from '../services/maps.service';
import { FormControl, FormGroup, Validators } from '@angular/forms';


@Component({
  selector: 'app-search-bar',
  templateUrl: './search-bar.component.html',
  styleUrls: ['./search-bar.component.scss']
})
export class SearchBarComponent implements OnInit{

  searchForm!: FormGroup;


  ngOnInit(): void {
    this.searchForm = new FormGroup({
      searchText: new FormControl('', [Validators.required])
    })
  }

  constructor(private http: HttpClient, private router: Router) { }

  searchRequest(){
    var searchText = this.searchForm.getRawValue().searchText
    console.log(this.searchForm.getRawValue())
    // this.http.post<any>('http://localhost:10000/students/', { Searchdata: searchText }).subscribe(data => { })
      this.router.navigate(['/products'])
  }
// We should give this in products page oninit
//  
  
}
