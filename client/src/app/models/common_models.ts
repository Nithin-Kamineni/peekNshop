import { results } from '../models/results'
import { Products1 } from './Products';

export interface LoginModel {
    
    Msg:           string;
    UserDetails:   string;
    
}
export interface SignupModel {
    
    Msg:           string;
    UserDetails:   string;
}
export interface Products {  
    Msg:           string;
}
export interface Stores {  
    results:        results[];
    icon:           string;
    name:           string;
}
export interface Products {  
    products1:        Products1;
    
}

export interface offers {  
    description:    string;
    name:           string;
}
export interface Location {  
    city:    string;
    
}
export interface productsDisplay {  
    photo:    string;
    price:           string;
    product_name:           string;
}