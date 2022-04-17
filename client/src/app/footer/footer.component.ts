import { Component, OnInit } from '@angular/core';
import { userdetails } from '../environments/User_Details'
@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.scss']
})
export class FooterComponent implements OnInit {
isLogin=userdetails.loggedIn
  constructor() { }

  ngOnInit(): void {
  }

}
