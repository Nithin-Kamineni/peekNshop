import { Component, OnInit } from '@angular/core';
import { ApiService } from '../services/api.service';
import { userdetails } from '../environments/User_Details'
@Component({
  selector: 'app-favorate-stores',
  templateUrl: './favorate-stores.component.html',
  styleUrls: ['./favorate-stores.component.scss']
})
export class FavorateStoresComponent implements OnInit {

  constructor(private api: ApiService) { }

  ngOnInit(): void {
    this.api.displayFavoriteStores(userdetails.id).subscribe((data: any) => {
      console.log(data)
      console.log(data.JWToken)
    })
  }

}
