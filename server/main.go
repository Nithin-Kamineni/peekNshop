package main

import (
	"log"
	"net/http"
	"src/controllers"
	"src/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// cookie, err := r.Cookie("token")
		// if err != nil {
		// 	if err == http.ErrNoCookie {
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// tokenStr := cookie.Value
		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func main() {

	router := mux.NewRouter()

	utils.ConnectDatabase()

	router.HandleFunc("/user", controllers.UserLogin).Methods("GET")                            //jwt proc
	router.HandleFunc("/user", controllers.UserSignUp).Methods("POST")                          //jwt proc
	router.HandleFunc("/user", controllers.ChangeUserDetails).Methods("PATCH")                  //*changing user details
	router.HandleFunc("/user/forgotpassword", controllers.ForgotUserDetails).Methods("POST")    //progress
	router.HandleFunc("/user/address", controllers.ChangeUserAddress).Methods("POST")           //address change user proc
	router.HandleFunc("/user/orders", controllers.SendUserOrders).Methods("POST")               //*sending user orders *list
	router.HandleFunc("/user/favorate-stores", controllers.ShowFavorateStores).Methods("POST")  //* favorateStores
	router.HandleFunc("/user/favorate-stores", controllers.AddingFavorateStores).Methods("PUT") //* add favorateStores
	router.HandleFunc("/user/favorate-stores", controllers.DeleFavorateStores).Methods("PATCH") //*  del favorateStores

	router.HandleFunc("/userStatus", controllers.UserStatus).Methods("POST")     //this
	router.HandleFunc("/userCheck", controllers.UserStatusCheck).Methods("POST") //this

	router.HandleFunc("/address/city", controllers.HomePageReload).Methods("POST") //static to google api
	router.HandleFunc("/contact", controllers.Contact).Methods("POST")             //this
	router.HandleFunc("/offers", controllers.ReturnOffers)                         //static
	router.HandleFunc("/address", controllers.ReturnLat).Methods("POST")           //*returning lat
	router.HandleFunc("/stores/", controllers.ReturnNearBy)                        //filter data from interface

	router.HandleFunc("/stores/add", controllers.AddStore).Methods("POST")                         //*add store information
	router.HandleFunc("/stores/add/{storeID}", controllers.AddInventory).Methods("POST")           //*add store inventory
	router.HandleFunc("/stores/edit/{storeID}", controllers.EditInventory).Methods("POST")         //*edit store inventory
	router.HandleFunc("/stores/delete/{storeID}", controllers.DeleteInventory).Methods("POST")     //*delete store inventory
	router.HandleFunc("/stores/items", controllers.ReturnStoreInv).Methods("GET")                  //*return store inventory
	router.HandleFunc("/stores/items/{product_id}", controllers.ReturnProductPage)                 //*display the product page
	router.HandleFunc("/stores/items/{product_id}", controllers.SendProductReview).Methods("POST") //*display the product page
	router.HandleFunc("/stores/items/{product_id}", controllers.SendProductRating).Methods("POST") //*display the product page

	router.HandleFunc("/cart", controllers.CartDisplay).Methods("POST")          //*this
	router.HandleFunc("/cart/clear-cart", controllers.ClearCart).Methods("POST") //*this
	router.HandleFunc("/cart", controllers.CartManipulation).Methods("PATCH")    //*this
	router.HandleFunc("/cart/additem", controllers.CartAddition).Methods("POST") //*this
	router.HandleFunc("/cart/", controllers.CartDeletion).Methods("PUT")         //*this

	router.HandleFunc("/order/payment", controllers.OrderPayment).Methods("POST")                  //secure
	router.HandleFunc("/order/{order_id}", controllers.DisplayOrders).Methods("POST")              //secure
	router.HandleFunc("/order/{order_id}/reviw", controllers.OrderReview).Methods("POST")          //secure
	router.HandleFunc("/order/{order_id}/reviw/edit", controllers.OrderReviewEdit).Methods("POST") //secure
	router.HandleFunc("/order/{order_id}/delivary", controllers.OrderDelivary).Methods("POST")     //secure

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
