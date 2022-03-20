import { results } from '../models/results'
export interface LoginModel {
    
    Msg:           string;
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