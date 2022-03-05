package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

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

	search := r.URL.Query().Get("search")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stores)

	w.Header().Set("Content-Type", "application/json")

	keyword := search
	radius := "1500"
	field := "formatted_address,name,rating,opening_hours,geometry"
	location := "29.61872,-82.37299"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "keyword=" + url.QueryEscape(keyword) + "&" +
		"radius=" + url.QueryEscape(radius) + "&" +
		"field=" + url.QueryEscape(field) + "&" +
		"location=" + url.QueryEscape(location) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/place/nearbysearch/json?", params)
	fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface{}
	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
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

// func (a *App) userLogin(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	//params := mux.Vars(r)
// 	//username := params["username"]
// 	//fmt.Println(username)
// 	var s Stores
// 	var reply Stores.status

// 	// var s Users.User3
// 	// var reply Users.LogInReply

// 	search := r.URL.Query().Get("search")
// 	// username := r.URL.Query().Get("email")
// 	// passkey := r.URL.Query().Get("passkey")
// 	//credentials := a.db.First(&s, "email = ?", username)
// 	err := a.db.Raw("SELECT id FROM user3 WHERE email = ?", username).Scan(&s).Error
// 	if err != nil {
// 		sendErr(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	// a.db.where("username = ?",username)
// 	//fmt.Println(&s)
// 	data, err := json.Marshal(&s)

// 	if s.ID == "" {
// 		fmt.Println("User does not exist/registered")
// 		reply = Users.LogInReply{AccessKey: "", RefreshKey: "", Msg: "User does not exist/registered", UserDetails: s}
// 		err = json.NewEncoder(w).Encode(reply)
// 		if err != nil {
// 			sendErr(w, http.StatusInternalServerError, err.Error())
// 		}

// 	} else {
// 		fmt.Println(s.ID)
// 		err = a.db.Raw("SELECT * FROM user3 WHERE id = ? AND password = ?", s.ID, passkey).Scan(&s).Error
// 		if err != nil {
// 			sendErr(w, http.StatusInternalServerError, err.Error())
// 			return
// 		}

// 		if s.Email == "" {
// 			fmt.Println("Password is incorrect")
// 			reply = Users.LogInReply{AccessKey: "", RefreshKey: "", Msg: "Password is incorrect", UserDetails: s}
// 			err = json.NewEncoder(w).Encode(reply)
// 			if err != nil {
// 				sendErr(w, http.StatusInternalServerError, err.Error())
// 			}
// 		} else {
// 			fmt.Println("Login Sucessfull")
// 			reply = Users.LogInReply{AccessKey: "", RefreshKey: "", Msg: "Login Sucessfull", UserDetails: s}
// 			err = json.NewEncoder(w).Encode(reply)
// 			if err != nil {
// 				sendErr(w, http.StatusInternalServerError, err.Error())
// 			}
// 		}
// 	}
// 	fmt.Println(string(data))
// 	fmt.Println()
// }
