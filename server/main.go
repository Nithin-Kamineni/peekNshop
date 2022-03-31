package main

import (
	"fmt"
	"log"
	"net/http"
	"src/controllers"
	"src/utils"

	//"os/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	//"strconv"
)

type App struct {
	db *gorm.DB
	r  *mux.Router
}

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

	router.HandleFunc("/address", controllers.returnLat) //returning lat
	router.HandleFunc("/address/", controllers.returnNearBy)
	router.HandleFunc("/offers", controllers.returnOffers)
	router.HandleFunc("/user", controllers.userLogin).Methods("GET")
	router.HandleFunc("/user/a", controllers.userSignUp).Methods("POST")
	router.HandleFunc("/address", controllers.returnLat) //returning lat
	router.HandleFunc("/stores/", controllers.returnNearBy)
	router.HandleFunc("/city", controllers.homePageReload)
	router.HandleFunc("/stores/add/{storeID}", controllers.addInventory).Methods("POST")
	router.HandleFunc("/stores/edit/{storeID}", controllers.editInventory).Methods("POST")
	router.HandleFunc("/stores/delete/", controllers.deleteInventory).Methods("POST")
	router.HandleFunc("/stores/items", controllers.returnStoreInv).Methods("POST")
	router.HandleFunc("/stores/items/{product_id}", controllers.returnProductPage)
	router.HandleFunc("/user", controllers.userLogin).Methods("POST")
	router.HandleFunc("/user/a", controllers.userSignUp).Methods("POST")
	router.HandleFunc("/user/forgotpassword", controllers.ForgotUserDetails).Methods("POST")
	router.HandleFunc("/userStatus", controllers.userStatus).Methods("POST")     //this
	router.HandleFunc("/userCheck", controllers.userStatusCheck).Methods("POST") //this
	router.HandleFunc("/cart", controllers.cartDisplay).Methods("POST")          //this
	router.HandleFunc("/cart/additem", controllers.cartAddition).Methods("POST") //this
	router.HandleFunc("/contact", controllers.contact).Methods("POST")           //this
	router.HandleFunc("/user", controllers.changeUserDetails).Methods("PUT")
	router.HandleFunc("/user/orders", controllers.sendUserOrders).Methods("POST")
	router.HandleFunc("/students/", controllers.getAllStudents).Methods("GET")
	router.HandleFunc("/students/", controllers.addStudent).Methods("POST")
	router.HandleFunc("/students/{id}", controllers.updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", controllers.deleteStudent).Methods("DELETE")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
