import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import {MapsService} from './services/maps.service';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {MatToolbarModule} from '@angular/material/toolbar';
import { HomepageComponent } from './homepage/homepage.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatCardModule} from '@angular/material/card';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatDividerModule} from '@angular/material/divider';
import {MatChipsModule} from '@angular/material/chips';
import {MatIconModule} from '@angular/material/icon';
import {MatSelectModule} from '@angular/material/select';
import {MatButtonModule} from '@angular/material/button';
import {MatListModule} from '@angular/material/list';
import { ProductsComponent } from './products/products.component';
import { ProductComponent } from './product/product.component';
import { ProductPageComponent } from './product-page/product-page.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { NavbarComponent } from './navbar/navbar.component';
import { CatogoriesComponent } from './catogories/catogories.component';
import { MainSliderComponent } from './main-slider/main-slider.component';
import { FooterComponent } from './footer/footer.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SidenavComponent } from './sidenav/sidenav.component';
import { HttpClientModule } from '@angular/common/http';
import { SearchBarComponent } from './search-bar/search-bar.component';
import { UserComponent } from './user/user.component';
import { UserHomepageComponent } from './user-homepage/user-homepage.component';
import { AboutComponent } from './about/about.component';
import { ContactComponent } from './contact/contact.component';
import { DeliveryComponent } from './delivery/delivery.component';
import { PaymentComponent } from './payment/payment.component';
import { ShopingCartComponent } from './shoping-cart/shoping-cart.component';
import { OrderDetailsComponent } from './order-details/order-details.component';
import { ChangeAddresspageComponent } from './change-addresspage/change-addresspage.component';
import { ChangeUserProfileSettingsComponent } from './change-user-profile-settings/change-user-profile-settings.component';
import { FavstoresComponent } from './favstores/favstores.component';
import { TrackingorderComponent } from './trackingorder/trackingorder.component';
import { StoresComponent } from './stores/stores.component';
import { DelivaryPageComponent } from './delivary-page/delivary-page.component';
import { FavorateStoresComponent } from './favorate-stores/favorate-stores.component';
import { CartComponent } from './cart/cart.component';

@NgModule({
  declarations: [
    AppComponent,
    HomepageComponent,
    ProductsComponent,
    ProductComponent,
    ProductPageComponent,
    NavbarComponent,
    
    CatogoriesComponent,
    MainSliderComponent,
    FooterComponent,
    SidenavComponent,
    SearchBarComponent,
    UserComponent,
    UserHomepageComponent,
    AboutComponent,
    ContactComponent,
    DeliveryComponent,
    PaymentComponent,
    ShopingCartComponent,
    OrderDetailsComponent,
    ChangeAddresspageComponent,
    ChangeUserProfileSettingsComponent,
    FavstoresComponent,
    TrackingorderComponent,
    StoresComponent,
    DelivaryPageComponent,
    FavorateStoresComponent,
    CartComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FlexLayoutModule,
    MatButtonModule,
    MatCardModule,
    MatChipsModule,
    MatIconModule,
    MatSelectModule,
    MatToolbarModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatSidenavModule,
    MatButtonModule,
    MatIconModule,
    MatDividerModule,
    ReactiveFormsModule,
    FormsModule,
    HttpClientModule,
    MatListModule,
    
    
    
  ],
  providers: [MapsService],
  bootstrap: [AppComponent]
})
export class AppModule { }
