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
import { StoresComponent } from './stores/stores.component';


const routes: Routes = [
  {path: "product/:id", component: ProductPageComponent},
  {path: "", component: HomepageComponent},

  {path: "user-homepage/user", component: UserComponent},
  {path: "user-homepage/userchangeaddress", component: ChangeAddresspageComponent},
  {path: "user-homepage/user/orders", component: OrderDetailsComponent},
  {path: "user-homepage", component: UserHomepageComponent},
  {path: "products", component: ProductsComponent}, 
  {path: "about", component: AboutComponent},
  {path: "user-homepage/delivery", component: DeliveryComponent},
  {path: "contact", component: ContactComponent},
  {path: "payment", component: PaymentComponent},
  {path: "product/id", component: ProductPageComponent},
  {path: "cart", component: ShopingCartComponent},
  {path: "user-homepage/user/changeusersettings", component: ChangeUserProfileSettingsComponent},
  {path: "stores", component: StoresComponent},
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

