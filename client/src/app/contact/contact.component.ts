import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.scss']
})
export class ContactComponent implements OnInit {
  focus: any;
  focus1: any;

  constructor(private router: Router) { }

  ngOnInit(): void {
  }
  sendMessage(){
    alert("message sent")
    this.router.navigate(['/'])
  }

}
