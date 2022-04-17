import { Component, OnInit } from '@angular/core';
import { Product } from '../core/product';
import { ProductService } from '../services/product.service';
import { environment} from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"
import { ApiService } from '../services/api.service';
import { userdetails } from '../environments/User_Details'
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
  
  constructor(private productService: ProductService, private http: HttpClient, private api: ApiService) {}

  

  ngOnInit() {
    this.productsdetails=this.api.getProducts().subscribe((data: any) => {
      this.productsdetails = data
    })
    
  }
  addtocart(i:number){
    console.log(this.productsdetails)
      let k=0
    for(let products of this.productsdetails){ 
      if (k==i){
        this.api.addtocart(userdetails.id,products.id,products.quantity, products.created, products.modified).subscribe((data: any) => {
        })
        
      }
    }
    
    
  }
}
