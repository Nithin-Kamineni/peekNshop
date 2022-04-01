package main

import (
	"fmt"
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

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func main() {

	router := mux.NewRouter()

	utils.ConnectDatabase()

	router.HandleFunc("/address", controllers.ReturnLat) //returning lat
	router.HandleFunc("/address/", controllers.ReturnNearBy)
	router.HandleFunc("/offers", controllers.ReturnOffers)
	router.HandleFunc("/user", controllers.UserLogin).Methods("GET")
	router.HandleFunc("/user/a", controllers.UserSignUp).Methods("POST")
	router.HandleFunc("/address", controllers.ReturnLat) //returning lat
	router.HandleFunc("/stores/", controllers.ReturnNearBy)
	router.HandleFunc("/city", controllers.HomePageReload)
	router.HandleFunc("/stores/add/{storeID}", controllers.AddInventory).Methods("POST")
	router.HandleFunc("/stores/edit/{storeID}", controllers.EditInventory).Methods("POST")
	router.HandleFunc("/stores/delete/", controllers.DeleteInventory).Methods("POST")
	router.HandleFunc("/stores/items", controllers.ReturnStoreInv).Methods("POST")
	router.HandleFunc("/stores/items/{product_id}", controllers.ReturnProductPage)
	router.HandleFunc("/user", controllers.UserLogin).Methods("POST")
	router.HandleFunc("/user/a", controllers.UserSignUp).Methods("POST")
	router.HandleFunc("/user/forgotpassword", controllers.ForgotUserDetails).Methods("POST")
	router.HandleFunc("/userStatus", controllers.UserStatus).Methods("POST")     //this
	router.HandleFunc("/userCheck", controllers.UserStatusCheck).Methods("POST") //this
	router.HandleFunc("/cart", controllers.CartDisplay).Methods("POST")          //this
	router.HandleFunc("/cart/additem", controllers.CartAddition).Methods("POST") //this
	router.HandleFunc("/contact", controllers.Contact).Methods("POST")           //this
	router.HandleFunc("/user", controllers.ChangeUserDetails).Methods("PUT")
	router.HandleFunc("/user/orders", controllers.SendUserOrders).Methods("POST")
	router.HandleFunc("/students/", controllers.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/", controllers.AddStudent).Methods("POST")
	router.HandleFunc("/students/{id}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
