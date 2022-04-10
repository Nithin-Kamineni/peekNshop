import { Component, OnInit } from '@angular/core';
import { ApiService } from '../services/api.service';
import { userdetails } from '../environments/User_Details'
@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.scss']
})
export class CartComponent implements OnInit {

  cartdetails:any
  constructor(private api: ApiService) { }

  ngOnInit(): void {
    this.api.cartdisplay(userdetails.id).subscribe((data: any) => {
      this.cartdetails= data
      console.log(data)
    })
  }

}
