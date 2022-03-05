import { Component, OnInit } from '@angular/core';
import { CatogoriesService} from './catogories.service';

@Component({
  selector: 'app-catogories',
  templateUrl: './catogories.component.html',
  styleUrls: ['./catogories.component.scss']
})
export class CatogoriesComponent implements OnInit {
  city = "Gainesville"

  constructor(public service: CatogoriesService) { }

  offers: any;

  ngOnInit(): void {
    this.service.getOffers().subscribe(data => {
      this.offers = data;
  })
  }

}
