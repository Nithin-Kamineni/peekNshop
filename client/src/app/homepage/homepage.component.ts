import { Component, Input, OnInit} from '@angular/core';
import {MapsService} from '../services/maps.service';


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
    this.service.getLocation();
    this.service.getCity();
  }

  

}
