import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-stores',
  templateUrl: './stores.component.html',
  styleUrls: ['./stores.component.scss']
})
export class StoresComponent implements OnInit {
isFavorite = false
  constructor() { }

  ngOnInit(): void {
  }
  favorite(){
    if(this.isFavorite==false){
      this.isFavorite=true
      document.getElementsByTagName("a")[6].style.backgroundColor = "pink";
    }else{
      this.isFavorite=false
      document.getElementsByTagName("a")[6].style.backgroundColor = "gray";
    }
    
  }
}
