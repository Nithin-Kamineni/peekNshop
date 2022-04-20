import { Component, OnInit } from '@angular/core';
import { Product } from '../core/product';
import { ProductService } from '../services/product.service';
import { environment} from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"
import { ApiService } from '../services/api.service';
import { userdetails } from '../environments/User_Details'
import { FormControl, FormGroup, Validators } from '@angular/forms';
@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.scss']
})
export class ProductsComponent implements OnInit {
  products: Product[] = [];
  photo=""
  productName=""
  price=""
  productsdetails:any
  cartItems=environment.numberOfItemsInCart
  addtocarttext="Add to cart"
  
  quantity=0;
  constructor(private productService: ProductService, private http: HttpClient, private api: ApiService) {}

  

  ngOnInit() {
    this.productsdetails=this.api.getProducts().subscribe((data: any) => {
      this.productsdetails = data
      console.log(this.productsdetails)
    })
    
  }
  addtocart(i:number){
    var k=0
    for(let products of this.productsdetails){ 
      if (k==i){
        console.log(products)
        if (products.quantity>=1){
          this.addtocarttext="Added to cart ✓"
          
        }
        if (products.quantity==0){
          this.addtocarttext="Add to cart"
          
        }

        var quantity = products.quantity
        quantity=quantity.toString()
        console.log(userdetails.id, products.productID, products.quantity, products.created, products.modified)
        this.api.addtocart(userdetails.id, products.productID, quantity, products.created, products.modified).subscribe((data: any) => {
        })
      }
      k = k+1
    }
    this.getNumberOfItemsInCart()
  }
  updateitem(i:number){

  }
  getNumberOfItemsInCart(){
    this.api.cartdisplay(userdetails.id).subscribe((data: any) => {
      var cartdetails= data
      var i=0
      for (let products of cartdetails){
        i = i+1
        environment.numberOfItemsInCart=i
        this.cartItems=i
      }
      console.log(environment.numberOfItemsInCart)
      
    })
    
  }
  increase(i:number){
    var k=0
    for (let products of this.productsdetails){
      if (k==i){
        if(this.addtocarttext=="Added to cart ✓"){
          this.addtocarttext="Update quantity"
        }
        products.quantity++
      }
      k=k+1
    }
  }
  decrease(i:number){
    var k=0
    for (let products of this.productsdetails){
      if (k==i){
        if(this.addtocarttext=="Added to cart ✓"){
          this.addtocarttext="Update quantity"
        }
        products.quantity--
      }
      k=k+1
    }
  }
  
}
