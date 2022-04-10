import { Component, OnInit } from '@angular/core';
import { Product } from '../core/product';
import { ProductService } from '../services/product.service';
import { SearchBarComponent } from '../search-bar/search-bar.component'
import { environment} from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"
import { ApiService } from '../services/api.service';
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
  addtocart(item:string){
    
  }
}
