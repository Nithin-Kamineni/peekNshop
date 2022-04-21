# PeekNshop

### Team Members
#### BackEnd
- Venkata Nithin Kamineni: vkamineni@ufl.edu github - Nithin-Kamineni

- Saieswar Reddy Vaka: vaka.s@ufl.edu github - sai6221
#### FrontEnd
- Vamsi Pachamatla: vamsi.pachamatla@ufl.edu github - vamsi3379

- Mattaparthi Nitish Veera: veeramattaparthi@ufl.edu github - VeeraNitish7

Github repository link: https://github.com/Nithin-Kamineni/peekNshop.git

## Outline

By using peekNshop users can look at the store's inventory and know whether the items they are looking for are available or not. They can also find the best deal in terms of cost and quality for the items they are looking for with our application.

> Technical stack, their pre-requisites and how to setup and run both frontend and backend can be found at this [wiki](https://github.com/Nithin-Kamineni/peekNshop/wiki/Installation)

### DEMO
[Demo Video of PeekNshop](https://www.youtube.com/watch?v=bLQmdahIzR0)

<strong>Cypress test video - </strong>
[Cypress test video](https://github.com/Nithin-Kamineni/peekNshop/tree/main/client/cypress/videos)

<strong>Backend unit test video - </strong>
[Backend unit test video](https://youtu.be/fgYtMIBUJJQ)



#  Frontend File Structure
###  Services: 
1.	All the HTTP requests like get, post, put and patch requests are stored and can be accessed from the api.service.ts file.
2.	maps.service.ts file consists of the geolocation function which can retrieve the user’s location and save it locally.

### Environments:
1.	Here we store the details of the user details, lat, and lon of the user. So that we can use that variable in the different components.

### Sidenav Component: This component has the login signup, location, cart, and logout features. 
1.	In the login form if users enter the correct credentials then it allows users to navigate to the user homepage.
2.	In the Signup form, It allows users to signup only once. If users successfully signup with the new email then the data is sent to the back end and registered for a new user.
3.	If the user denies sharing the location then the user can able to enter their address manually in the location form. 
4.	If users successfully log in then the navigation bar displays the username, logout button, and the cart. If the user has some items in the cart then the count of the product is also displayed in the navigation bar.
5.	Here we store the JWT token in local storage to log in the user if the page is reloaded. After successful login.
6.	Users can able to search for types of stores nearby them. For instance, if a user search for “food” then it redirects to the stores component and displays the food stores.

### Homepage Component: It contains the categories component and the footer.
1.	We call the geolocation function from this component to retrieve the user’s location. If the user allows then it is stored in environment variables. After retrieving the location it calls for the HTTP request to for getting offers and local stores dynamically.
2.	After getting the stores user can able to visit the store to see products available and add that particular store to their favorites.

### Search-Bar Component:
1.	Users can able to search for products in the search bar and it redirects to the products component and displays the products.

### User-profile Component:
1.	Here users can able to see their previous orders, change account settings, and change shipping addresses, and also users can able to add their pictures to this component.
2.	We have also included freelance delivery buttons for our future work.

### Change-Account-setting Component:
1.	Right after getting into this component with the help of a post HTTP request, it displays the user details where users can edit them and make changes to them and also it updates the user table in the backend.

### Change-Shipping-Address:
1.	 Here users can able to add the address for the delivery.

### Logout:
1.	  When the user logout then the user details are completely erased from the environment and also user’s JWT tokens are removed from the local storage.

### Shopping-Cart Component:
1.	If the cart is empty then it asks uses to shop.
2.	If the user has some products in the cart then the user can able to see them in this component.
3.	Users can able to delete the product from the cart. Where it sends the post request and removes the particular product in the database.
4.	If users click the empty cart button then all the products from the cart are removed.
5.	It displays the quantities of the different products and displays the grand total of the cart.
6.	After clicking on the checkout button it goes to the checkout component and asks for card details.

### About-us Component:
1.	It displays complete details of our project.

### Product-Page-Component:
1.	It displays the complete details of the particular product.

### Favorite-Stores Component:
1.	If users add some stores as a favorite then this component displays the favorite stores of the user.
  
  
  
# BacKend File structure
  ###  main.go
  main.go is the entry point to start the server and is present in package main. All the http configuration needed for accepting incoming request and forwarding it to the handler methods is done in main.go. Database connection is also initialized when main.go is run.
  
  ### controllers
In controllers package, the handler methods that serve the request are present.
store_controller.go file has the handler methods for adding inventories and updating items of stores.
User_Controller.go has the handler methods to update bio, user details, address and cart of users.
cart_controller.go has the methods to add and delete  items in the users cart respectively.
orders_controller.go contains the methods to create, update, retrieve and re-order to the store for  items by the users.

All the controller files have a corresponding *-test.go files and contain the test cases for all create and update methods.
  
 
### models
In models package, model structures for Users, Cart_items, offers, store_inventory are currently defined. These structs relate to the database table when GORM configuration is included in them.

### utils
utils package deals with initializing database connection, creating required tables, and auto migrating the models.
