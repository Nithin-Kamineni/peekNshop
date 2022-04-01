package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"src/models"
	"src/utils"
)

func AddInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	var accessID string
	reply := models.SignInReply{Msg: "sucessfully added the new item/items in the inventory"}
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	err = utils.DB.Raw("SELECT accessKey FROM user3 WHERE store_id = ?", store.StoreID).Scan(&accessID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if accessID == store.AccessKey {
		err = utils.DB.Save(&store).Error
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
}

func EditInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	var accessID string
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT accessKey FROM user3 WHERE store_id = ?", store.StoreID).Scan(&accessID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if accessID == store.AccessKey {
		err = utils.DB.Exec("UPDATE store_inventory SET ProductPrice = ?, ProductName = ?, Quantity = ?, ModifiedAt = ? WHERE StoreID = ? and ProductID = ?", store.ProductPrice, store.ProductName, store.Quantity, store.ModifiedAt, store.StoreID, store.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully changed your details"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	var accessID string
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT accessKey FROM user3 WHERE store_id = ?", store.StoreID).Scan(&accessID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if accessID == store.AccessKey {
		err = utils.DB.Exec("DELETE from store_inventory WHERE StoreID = ? and ProductID = ?", store.StoreID, store.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully changed your details"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func ReturnOffers(w http.ResponseWriter, r *http.Request) {
	utils.DB.Model(&models.Offer{}).Create([]map[string]interface{}{
		{"name": "jinzhu_1", "description": "10% off on all items"},
		{"name": "jinzhu_2", "description": "20% off on all items"},
		{"name": "ROSS", "description": "BOGO offer 50% off on all items"},
		{"name": "Whole Foods", "description": "Friday foods 50% of on all ready to eat meals"},
	})
	w.Header().Set("Content-Type", "application/json")
	var all []models.Offer
	err := utils.DB.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnNearBy(w http.ResponseWriter, r *http.Request) {

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

func FilterInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inv models.Cart_items
	var userID models.UserIDtab
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM storesInventory WHERE storeID = ?", userID).Scan(&inv).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(inv)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnStoreInv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store []models.Store_inventory
	storeID := r.URL.Query().Get("store_id")
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	err = utils.DB.Raw("SELECT * FROM user3 WHERE store_id = ?", storeID).Scan(&store).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnProductPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	product_id := r.URL.Query().Get("product_id")

	err := utils.DB.Raw("SELECT * FROM store_inventories WHERE product_id = ?", product_id).Scan(&store).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnLat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var address models.Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	//address := "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "address=" + url.QueryEscape(address.Address) + "&" +
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
