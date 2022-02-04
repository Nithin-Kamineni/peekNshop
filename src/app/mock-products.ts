import { Product } from './core/product';
import { Size } from './core/size';

export const PRODUCTS: Product[] = [
  {
    id: 1,
    name: 'Amazon',
    imageUrls: ['https://air-marketing-assets.s3.amazonaws.com/blog/logo-db/amazon-logo-png/amazon-logo-png-svg-4.svg'
    , '../assets/ice-cream-cherry.svg', '../assets/ice-cream-squash.svg'],
    price: 10,
    flavors: [
      { name: 'prune', color: '#5A188E' },
      { name: 'squash', color: '#F88532' },
      { name: 'cherry', color: '#E91E63' },
    ],
    sizes: [Size.SMALL, Size.MEDIUM, Size.LARGE],
  },
  {
    id: 2,
    name: 'Walmart',
    imageUrls: ['https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.freepnglogos.com%2Fpics%2Fwalmart-logo&psig=AOvVaw3q03RxKrYuHa29iWsBlkuk&ust=1644040354881000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCPilkd2t5fUCFQAAAAAdAAAAABAD', 
    '../assets/popsicle-lettuce.svg', '../assets/popsicle-cherry.svg'],
    price: 8,
    flavors: [
      { name: 'lime', color: '#00CACA' },
      { name: 'lettuce', color: '#80DC0B' },
      { name: 'cherry', color: '#E91E63' },
    ],
    sizes: [Size.SMALL, Size.LARGE],
  },
];
