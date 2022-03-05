import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";

@Component({
  selector: 'app-user-homepage',
  templateUrl: './user-homepage.component.html',
  styleUrls: ['./user-homepage.component.scss']
})
export class UserHomepageComponent implements OnInit {

  title = 'PeekNshop';
  subtitle = 'choose where to shop?';
  city = "Gainesville";
  user = "Nithin Kamineni"
  

  constructor(private router: Router) { }

  ngOnInit(): void {
    
  }
  userform(): void{
    this.router.navigate(['/user'])
  }

}
