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
  deleteFromCart(i:number){
    var k=0
    console.log(i)
    for (let cart of this.cartdetails){
      if (k==i){
        this.productID=cart.productID
        this.quantity=cart.quantity
        this.created=cart.created
        this.modified=cart.modified
        console.log(userdetails.id, this.productID, this.quantity, this.created, this.modified)
      }
      k=k+1
    }
    this.api.removeProductFromCart(userdetails.id, this.productID, this.quantity, this.created, this.modified)
  }
}
