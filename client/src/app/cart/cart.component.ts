import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { ApiService } from '../services/api.service';
import { userdetails } from '../environments/User_Details'
import { environment } from '../environments/environments';
@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.scss']
})
export class CartComponent implements OnInit {

  cartdetails:any
  productID:any
  quantity:any
  created:any
  modified:any
  constructor(private api: ApiService, private router: Router) { }

  ngOnInit(): void {
    this.api.cartdisplay(userdetails.id).subscribe((data: any) => {
      this.cartdetails= data
      console.log(data)
    })
  }
  checkout(){
    this.router.navigate(['/user-homepage/user/cart/payment'])
  }
  deleteFromCart(){
    this.api.removeProductFromCart(userdetails.id, this.productID, this.quantity, this.created, this.modified)
  }
}
