import { Component, OnInit } from '@angular/core';
import { Product } from '../core/product';
import { ProductService } from '../services/product.service';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.scss']
})
export class ProductsComponent implements OnInit {
  products: Product[] = [];

  constructor(private productService: ProductService) {}

  getProducts(): void {
    this.products = this.productService.getProducts();
    // getProducts(){
      //     const httpHeaders = new HttpHeaders();
      //     httpHeaders.append('content-type','application/json')
      //     return this.http.get('http://', {headers: httpHeaders});
      //   }
      
  }

  ngOnInit() {
    this.getProducts();
  }
}
