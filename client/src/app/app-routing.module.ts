import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomepageComponent } from './homepage/homepage.component';
import { ProductPageComponent } from './product-page/product-page.component';

import { UserComponent } from './user/user.component';
import { UserHomepageComponent } from './user-homepage/user-homepage.component';
import { ProductsComponent } from './products/products.component';
import { AboutComponent } from './about/about.component';
import { SidenavComponent } from './sidenav/sidenav.component';
import { ContactComponent } from './contact/contact.component';
import { DeliveryComponent } from './delivery/delivery.component';
import { PaymentComponent } from './payment/payment.component';
import { ShopingCartComponent } from './shoping-cart/shoping-cart.component';
import { ChangeAddresspageComponent } from './change-addresspage/change-addresspage.component';
import { OrderDetailsComponent } from './order-details/order-details.component';
import { ChangeUserProfileSettingsComponent } from './change-user-profile-settings/change-user-profile-settings.component';


const routes: Routes = [
  {path: "product/:id", component: ProductPageComponent},
  {path: "", component: HomepageComponent},

  {path: "user", component: UserComponent},
  {path: "user/changeaddress", component: ChangeAddresspageComponent},
  {path: "user/orders", component: OrderDetailsComponent},
  {path: "user-homepage", component: UserHomepageComponent},
  {path: "products", component: ProductsComponent}, 
  {path: "about", component: AboutComponent},
  {path: "delivery", component: DeliveryComponent},
  // {path: "", component: SidenavComponent},
  {path: "contact", component: ContactComponent},
  {path: "payment", component: PaymentComponent},
  {path: "product/id", component: ProductPageComponent},
  {path: "cart", component: ShopingCartComponent},
  {path: "user/changeusersettings", component: ChangeUserProfileSettingsComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

