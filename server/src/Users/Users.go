package Users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Stores struct {
	StoresID  string `json:"StoresId"`
	StoreName string `json:"StoreName"`
	Catogiry  string `json:"catogiry"`
}

var stores []Stores

//sending search results
func getSearchResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stores)
}

//getting nearest stores avilable based on the location of the user
func getNearestStores(w http.ResponseWriter, r *http.Request) {

}

//sending inventory avilable in the store
func getStoresDetails(w http.ResponseWriter, r *http.Request) {

}

//sending the details of the product
func getProductDetails(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init Router
	r := mux.NewRouter()

	stores = append(stores, Stores{StoresID: "1", StoreName: "IndiMart", Catogiry: "Glosseries"})
	//Rought handeler or endpoints
	r.HandleFunc("/user/{id}/searchResults/{searchItem}", getSearchResults).Methods("POST")
	r.HandleFunc("/user/{id}/stores", getNearestStores).Methods("GET")
	r.HandleFunc("/user/{id}/stores/{storeId}", getStoresDetails).Methods("GET")
	r.HandleFunc("/user/{id}/stores/{storeId}/{productId}", getProductDetails).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", r))
}
