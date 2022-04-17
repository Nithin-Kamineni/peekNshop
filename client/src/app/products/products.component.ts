import { Component, OnInit } from '@angular/core';
import { Product } from '../core/product';
import { ProductService } from '../services/product.service';
import { SearchBarComponent } from '../search-bar/search-bar.component'
import { environment} from '../environments/environments'
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http"
import { productsDisplay} from '../models/common_models'
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
  
  constructor(private productService: ProductService, private http: HttpClient) {}

  getProducts(): void {

    

    // this.products = this.productService.getProducts();
    // getProducts(){
      //     const httpHeaders = new HttpHeaders();
      //     httpHeaders.append('content-type','application/json')
      //     return this.http.get('http://', {headers: httpHeaders});
      //   }
      
  }

  ngOnInit() {
    this.http.get<any>('http://localhost:10000/stores/items?store_id='+environment.storeId, {}).subscribe( (data: any) => {
      console.log(data)
      this.photo = data[0].photo
      this.productName = data[0].productName
      this.price = data[0].price
    })
    
  }
}
