import { Component, Input, OnInit} from '@angular/core';
import {MapsService} from '../services/maps.service';
import { environment } from '../environments/environments'

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit{
  title = 'PeekNshop';
  subtitle = 'choose where to shop?';
  
  
  

  constructor(public service: MapsService) { }

  ngOnInit(): void {
    if (environment.isLocation){
      this.service.getLocation();
    }
    
  }

  

}
