package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"src/src/Carts"
	"src/src/Offers"
	"src/src/Stores"
	"src/src/Users"

	"github.com/google/uuid"
)

func (a *App) addInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store Stores.Store_inventory
	reply := Users.SignInReply{Msg: "sucessfull"}
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	store.StoreID = uuid.New().String()
	err = a.db.Save(&store).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) editInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storeUser Stores.Store_inventory
	//var storeDB Stores.Store_inventory
	err := json.NewDecoder(r.Body).Decode(&storeUser)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	// err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE StoreID = ? and ProductID = ?", s1.ID).Scan(&s2).Error
	// if err != nil {
	// 	sendErr(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// fmt.Println(s1.Acesskey)
	// fmt.Println(s2.Acesskey)
	// fmt.Println()
	//if s2.Acesskey == s1.Acesskey {
	err = a.db.Exec("UPDATE store_inventory SET ProductPrice = ?, ProductName = ?, Quantity = ?, ModifiedAt = ? WHERE StoreID = ? and ProductID = ?", storeUser.ProductPrice, storeUser.ProductName, storeUser.Quantity, storeUser.ModifiedAt, storeUser.StoreID, storeUser.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	reply := Users.SignInReply{Msg: "sucessfully changed your details"}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) deleteInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storeUser Stores.Store_inventory
	//var storeDB Stores.Store_inventory
	err := json.NewDecoder(r.Body).Decode(&storeUser)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	// err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE StoreID = ? and ProductID = ?", s1.ID).Scan(&s2).Error
	// if err != nil {
	// 	sendErr(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// fmt.Println(s1.Acesskey)
	// fmt.Println(s2.Acesskey)
	// fmt.Println()
	//if s2.Acesskey == s1.Acesskey {
	err = a.db.Exec("DELETE from store_inventory WHERE StoreID = ? and ProductID = ?", storeUser.StoreID, storeUser.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	reply := Users.SignInReply{Msg: "sucessfully changed your details"}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

	//a.db.Raw("SELECT  FROM user3 WHERE acesskey = ?", username)
	// err = a.db.Save(&s).Error
	// if err != nil {
	// 	sendErr(w, http.StatusInternalServerError, err.Error())
	// }
	//}
}

func (a *App) returnOffers(w http.ResponseWriter, r *http.Request) {
	a.db.Model(&Offers.Offer{}).Create([]map[string]interface{}{
		{"name": "jinzhu_1", "description": "10% off on all items"},
		{"name": "jinzhu_2", "description": "20% off on all items"},
	})
	w.Header().Set("Content-Type", "application/json")
	var all []Offers.Offer
	err := a.db.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) returnNearBy(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")
	w.Header().Set("Content-Type", "application/json")

	keyword := search
	radius := "1500"
	field := "formatted_address,name,rating,opening_hours,geometry"
	location := lat + "," + long
	// fmt.Println(location)
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "keyword=" + url.QueryEscape(keyword) + "&" +
		"radius=" + url.QueryEscape(radius) + "&" +
		"field=" + url.QueryEscape(field) + "&" +
		"location=" + url.QueryEscape(location) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/place/nearbysearch/json?", params)
	// fmt.Println(path)
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

func (a *App) filterInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inv Carts.Cart_items
	var userID Carts.UserIDtab
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT * FROM storesInventory WHERE storeID = ?", userID).Scan(&inv).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(inv)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) returnStoreInv(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	keyword := "foods"
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

func (a *App) returnProductPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	keyword := "foods"
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

func (a *App) ConvAddressToCord(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	address := "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "address=" + url.QueryEscape(address) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/geocode/json?", params)
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

func (a *App) returnLat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	address := "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "address=" + url.QueryEscape(address) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/geocode/json?", params)
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
	var f interface {
		getHtml() string
	}

	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}
