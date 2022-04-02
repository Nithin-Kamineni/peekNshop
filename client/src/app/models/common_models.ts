import { results } from '../models/results'

export interface LoginModel {
    
    Msg:           string;
    UserDetails:   string;
    
}
export interface SignupModel {
    
    Msg:           string;
}
export interface Products {  
    Msg:           string;
}
export interface Stores {  
    results:        results[];
    icon:           string;
    name:           string;
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