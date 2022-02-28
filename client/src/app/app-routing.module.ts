import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomepageComponent } from './homepage/homepage.component';
import { ProductPageComponent } from './product-page/product-page.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { UserComponent } from './user/user.component';
import { UserHomepageComponent } from './user-homepage/user-homepage.component';
import { ProductsComponent } from './products/products.component';
import { AboutComponent } from './about/about.component';


const routes: Routes = [
  {path: "product/:id", component: ProductPageComponent},
  {path: "", component: HomepageComponent},
  {path: "login", component: LoginComponent},
  {path: "signup", component: SignupComponent},
  {path: "user", component: UserComponent},
  {path: "user-homepage", component: UserHomepageComponent},
  {path: "products", component: ProductsComponent},
  {path: "about", component: AboutComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
export const routingComponents = [LoginComponent]
