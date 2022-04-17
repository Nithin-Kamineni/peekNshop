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

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contactMsg models.ContactMsgObj
	reply_on_succ := models.SignInReply{Msg: "Message Recived"}
	reply_on_fail := models.SignInReply{Msg: "Message Failed."}
	err := json.NewDecoder(r.Body).Decode(&contactMsg)
	if err != nil {
		// sendErr(w, http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_fail)
		return
	}
	err = utils.DB.Save(&contactMsg).Error
	if err != nil {
		// sendErr(w, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_fail)
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_succ)
	}
}

func HomePageReload(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(statusCode: 200)
	//w.WriteHeader(statusCode: 200)
	w.Header().Set("Content-Type", "application/json")

	//params := mux.Vars(r)
	//username := params["username"]
	//fmt.Println(username)
	var reply models.HomePageCity
	var cord models.Coardinates

	err := json.NewDecoder(r.Body).Decode(&cord)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	//credentials := a.db.First(&s, "email = ?", username)
	fmt.Println("Gainsville")
	reply = models.HomePageCity{City: "Gainesville"}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ConvAddressToCord(w http.ResponseWriter, r *http.Request) {

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

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
