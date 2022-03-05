package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// {
// 	"business_status": "OPERATIONAL",
// 	"geometry": {
// 		"location": {
// 			"lat": 29.6254076,
// 			"lng": -82.37590070000002
// 		},
// 		"viewport": {
// 			"northeast": {
// 				"lat": 29.62690757989272,
// 				"lng": -82.37484922010728
// 			},
// 			"southwest": {
// 				"lat": 29.62420792010728,
// 				"lng": -82.37754887989271
// 			}
// 		}
// 	},
// 	"icon": "https://maps.gstatic.com/mapfiles/place_api/icons/v1/png_71/shopping-71.png",
// 	"icon_background_color": "#4B96F3",
// 	"icon_mask_base_uri": "https://maps.gstatic.com/mapfiles/place_api/icons/v2/shoppingcart_pinlet",
// 	"name": "Whole Foods Market",
// 	"opening_hours": {
// 		"open_now": true
// 	},
// 	"photos": [
// 		{
// 			"height": 4096,
// 			"html_attributions": [
// 				"<a href=\"https://maps.google.com/maps/contrib/103293559241366476507\">A Google User</a>"
// 			],
// 			"photo_reference": "Aap_uECJQ1rkr7kPy2hd_Hh3EmoN8kRWt-uD36_gv6y94a-GkRZb1xd1fw5MxfF6s4S5Zmm4FsSlNoVIXn4IuMHnWkZ3Hr-uhZY9WLvwTn9yYivvM1nkGpAXIRX1Ys5aT4XO_LEhQR8UUS08CC5jyoE71ZjhjGIkkYFcnoWBPKareVYihUU9",
// 			"width": 3072
// 		}
// 	],
// 	"place_id": "ChIJbxdDECWj6IgRNRNDwpIAZjI",
// 	"plus_code": {
// 		"compound_code": "JJGF+5J Gainesville, Florida",
// 		"global_code": "76XVJJGF+5J"
// 	},
// 	"price_level": 3,
// 	"rating": 4.3,
// 	"reference": "ChIJbxdDECWj6IgRNRNDwpIAZjI",
// 	"scope": "GOOGLE",
// 	"types": [
// 		"grocery_or_supermarket",
// 		"supermarket",
// 		"food",
// 		"health",
// 		"point_of_interest",
// 		"store",
// 		"establishment"
// 	],
// 	"user_ratings_total": 805,
// 	"vicinity": "3490 SW Archer Rd, Gainesville"
// }

type Stores struct {
	html_attributions []string                   `json:"html_attributions"`
	results           []struct{ nearest_stores } `json:"results"`
	status            string                     `json:"status"`
}

type nearest_stores struct {
	business_status       string                     `json:"business_status"`
	geometry              struct{ precise_location } `json:"geometry"`
	icon                  string                     `json:"icon"`
	icon_background_color string                     `json:"icon_background_color"`
	icon_mask_base_uri    string                     `json:"icon_mask_base_uri"`
	name                  string                     `json:"name"`
	opening_hours         struct{ open_status }      `json:"name"`
	photos                []struct{ photo_detail }   `json:"photos"`
	place_id              string                     `json:"place_id"`
	plus_code             struct{ other_codes }      `json:"plus_code"`
	price_level           int                        `json:"price_level"`
	rating                float32                    `json:"rating"`
	reference             string                     `json:"reference"`
	scope                 string                     `json:"scope"`
	types                 []string                   `json:"types"`
	user_ratings_total    int                        `json:"user_ratings_total"`
	vicinity              string                     `json:"vicinity"`
}

type precise_location struct {
	location struct{ latlong }    `json:"location"`
	viewport struct{ directions } `json:"viewport"`
}

type latlong struct {
	lat float32 `json:"lat"`
	lng float32 `json:"lng"`
}

type directions struct {
	northeast struct{ latlong } `json:"northeast"`
	southwest struct{ latlong } `json:"southwest"`
}

type open_status struct {
	open_now bool `json:"open_now"`
}

type photo_detail struct {
	height            int      `json:"height"`
	html_attributions []string `json:"html_attributions"`
	photo_reference   string   `json:"string"`
	width             int      `json:"width"`
}

type other_codes struct {
	compound_code string `json:"compound_code"`
	global_code   string `json:"global_code"`
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

	// stores = append(stores, Stores{StoresID: "1", StoreName: "IndiMart", Catogiry: "Glosseries"})
	//Rought handeler or endpoints
	r.HandleFunc("/user/{id}/searchResults/{searchItem}", getSearchResults).Methods("POST")
	r.HandleFunc("/user/{id}/stores", getNearestStores).Methods("GET")
	r.HandleFunc("/user/{id}/stores/{storeId}", getStoresDetails).Methods("GET")
	r.HandleFunc("/user/{id}/stores/{storeId}/{productId}", getProductDetails).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
