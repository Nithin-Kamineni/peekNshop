import { photos } from '../models/photos'

export interface results { 
    icon:           string;
    name:           string;
    photos:        photos[];
}